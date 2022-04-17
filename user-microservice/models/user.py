from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.orm import validates
import re

db = SQLAlchemy()
email_regex = r'\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b'


class User(db.Model):
    __tablename__ = 'users'

    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    user_name = db.Column(db.String(120), unique=True, nullable=False)
    email = db.Column(db.String(120), unique=True, nullable=False)
    name = db.Column(db.String(120), nullable=False)
    surname = db.Column(db.String(120), nullable=False)

    def serialize(self):
        return {
            'id': self.id,
            'user_name': self.user_name,
            'email': self.email,
            'name': self.name,
            'surname': self.surname
        }

    @validates('email')
    def validate_email(self, key, email):
        if not re.fullmatch(email_regex, email):
            raise ValueError("Failed email validation")
        return email
