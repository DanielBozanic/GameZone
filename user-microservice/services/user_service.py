from db.database import db
from models.user import User


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
                active=True
            )
            db.session.add(new_user)
            db.session.commit()
            return ""
        except ValueError as err:
            return str(err)
    else:
        return "User with this email/username already exists!"

    # TODO Email service


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
                active=True
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


def delete(user_id):
    user_db = get_by_id(user_id)
    if isinstance(user_db, str):
        return "User with id " + str(user_id) + " does not exist!"

    db.session.query(User). \
        filter(User.id == user_id).delete()
    db.session.commit()
    return ""
