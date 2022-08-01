from db.database import db
from models.user import User
import requests


def login(data):
    admin_exists = User.query\
        .filter("ROLE_ADMIN" == User.role).first()

    if not admin_exists:
        admin = User(
            user_name="admin",
            password="123",
            email="admin@gmail.com",
            name="admin",
            surname="admin",
            role="ROLE_ADMIN",
            active=True
        )
        db.session.add(admin)
        db.session.commit()

    user = User.query\
        .filter(data["user_name"] == User.user_name)\
        .filter(data["password"] == User.password).first()
    if not user:
        return "Incorrect username/password"

    if not user.verified:
        return "Your account is not verified."

    try:
        resp = requests.get(
            'http://localhost:7003/api/contactAndReport/bans/isUserBanned/' + str(user.id),
            headers={'Content-Type': 'application/json', 'Accept': 'application/json'}
        )
        if resp.status_code == 200 and resp.json():
            return "Your account is banned"
    except requests.exceptions.RequestException:
        pass

    return user
