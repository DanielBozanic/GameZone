from functools import wraps
from flask import request, jsonify
from config import SECRET_KEY
import jwt
import datetime


def encode_auth_token(user):
    try:
        payload = {
            'exp': datetime.datetime.utcnow() + datetime.timedelta(days=0, hours=12),
            'iat': datetime.datetime.utcnow(),
            'sub': {
                "user_id": user.id,
                "email": user.email,
                "role": user.role
            }
        }
        return jwt.encode(
            payload,
            SECRET_KEY,
            algorithm='HS256'
        )
    except Exception as e:
        return e


def authentification_required(f):
    @wraps(f)
    def decorated():
        token = request.headers.get('Authorization')
        if not token:
            return jsonify({'message': "Token is missing!"}), 403

        token = token.split(" ")[1]
        try:
            token = jwt.decode(token, SECRET_KEY)
        except:
            return jsonify({"message": "Token is invalid!"}), 403

        return f(token)
    return decorated


def roles_required(roles):
    def decorate(f):
        @wraps(f)
        def wrapper(token):
            if token['sub']['role'] not in roles:
                return jsonify({"message": "Unauthorized!"}), 401
            else:
                return f()
        return wrapper
    return decorate
