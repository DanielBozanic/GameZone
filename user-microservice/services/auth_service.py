from db.database import db
from models.user import User


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
        return "Your account is not verified, please check your email"

    return user
