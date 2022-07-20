from db.database import db


class UserVerification(db.Model):
    __tablename__ = 'user_verification'

    id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    email = db.Column(db.String(120), unique=True, nullable=False)
    code = db.Column(db.String(30), nullable=False)
    expiration_date = db.Column(db.DateTime, nullable=False)

    def serialize(self):
        return {
            'id': self.id,
            'email': self.email,
            'code': self.code,
            'expiration_date': self.expiration_date
        }
