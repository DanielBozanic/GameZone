from flask_sqlalchemy import SQLAlchemy
from models.user import User

db = SQLAlchemy()


def create(data):
    user = User.query.filter((data['email'] == User.email) |
                             (data['user_name'] == User.user_name)).first()
    if not user:
        new_user = User(
            user_name=data['user_name'],
            email=data['email'],
            name=data['name'],
            surname=data['surname']
        )
        db.session.add(new_user)
        db.session.commit()
        return ""
    else:
        return "User with this email/username already exists!"
