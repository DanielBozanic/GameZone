from db.database import db
from models.user import User
from models.user_verification import UserVerification
import requests
import random
import string
import datetime


def register(data):
    user = User.query.filter((data['email'] == User.email) |
                             (data['user_name'] == User.user_name)).first()
    if not user:
        try:
            new_user = User(
                user_name=data['user_name'],
                password=data["password"],
                email=data['email'],
                name=data['name'],
                surname=data['surname'],
                role="ROLE_USER",
                verified=False
            )
            db.session.add(new_user)
            db.session.commit()
            return ""
        except ValueError as err:
            return str(err)
    else:
        return "User with this email/username already exists!"


def get_verification_code(email):
    user = User.query.filter(User.email.like(email)).first()
    if not user:
        return "User with email " + str(email) + " does not exist!", 400

    exists = UserVerification.query \
        .filter(UserVerification.email.like(email)).first()
    if exists:
        db.session.query(UserVerification). \
            filter(UserVerification.id == exists.id).delete()
        db.session.commit()

    user_verification = UserVerification()
    user_verification.email = user.email
    user_verification.code = ''.join(random.choices(string.ascii_uppercase + string.digits, k=7))
    user_verification.expiration_date = datetime.datetime.now() + datetime.timedelta(minutes=20)
    db.session.add(user_verification)
    db.session.commit()
    content = {
        "subject": "Account verification",
        "recipients": [user.email],
        "content": {
            "template": "verification",
            "params": {
                "name": user.name,
                "surname": user.surname,
                "email": user.email,
                "code": user_verification.code
            }
        }
    }
    resp = requests.post(
        'http://localhost:5001/api/email/sendEmail',
        json=content,
        headers={'Content-Type': 'application/json', 'Accept': 'application/json'}
    )
    return resp.text, resp.status_code


def verify_account(data):
    user = User.query \
        .filter(User.email.like(data["email"])).first()
    if not user:
        return "Email not found"

    user_verification = UserVerification.query\
        .filter(UserVerification.email.like(data["email"])) \
        .filter(UserVerification.code.like(data["code"])).first()
    if not user_verification:
        return "Invalid code"

    if datetime.datetime.now() > user_verification.expiration_date:
        return "Code has expired"

    db.session.query(UserVerification). \
        filter(UserVerification.id == user_verification.id).delete()

    User.query.filter_by(email=user.email).update(dict(verified=True))
    db.session.commit()

    return ""


def add_employee_and_admin(data):
    user = User.query.filter((data['email'] == User.email) |
                             (data['user_name'] == User.user_name)).first()
    if not user:
        try:
            new_user = User(
                user_name=data['user_name'],
                password=data["password"],
                email=data['email'],
                name=data['name'],
                surname=data['surname'],
                role=data['role'],
                verified=True
            )
            db.session.add(new_user)
            db.session.commit()
            return ""
        except ValueError as err:
            return str(err)
    else:
        return "User with this email/username already exists!"


def get_all():
    return User.query.all()


def get_all_registered_users():
    return User.query \
        .filter(("ROLE_USER" == User.role)).all()


def get_by_id(user_id):
    user = User.query.filter((user_id == User.id)).first()
    if user:
        return user
    return "User with id " + str(user_id) + " does not exist!"


def update(data):
    user_db = get_by_id(data["id"])
    if isinstance(user_db, str):
        return "User with id " + str(data["id"]) + " does not exist!"

    user_db = User.query\
        .filter(data["id"] != User.id)\
        .filter((data['email'] == User.email) |
                (data['user_name'] == User.user_name)).first()

    if not user_db:
        db.session.query(User). \
            filter(User.id == data["id"]). \
            update({'email': data["email"], "user_name": data["user_name"],
                    "password": data["password"], "name": data["name"],
                    "surname": data["surname"]})
        db.session.commit()
        return ""
    else:
        return "User with this email/username already exists!"
