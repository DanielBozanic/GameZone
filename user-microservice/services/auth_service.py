from flask_sqlalchemy import SQLAlchemy
from models.user import User

db = SQLAlchemy()


def login(data):
    user = User.query\
        .filter(data["user_name"] == User.user_name)\
        .filter(data["password"] == User.password).first()
    if not user:
        return "Incorrect username/password"

    return user
