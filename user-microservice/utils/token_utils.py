from functools import wraps
from flask import request, jsonify
import jwt
import config
import datetime


def encode_auth_token(user):
    try:
        payload = {
            'exp': datetime.datetime.utcnow() + datetime.timedelta(days=0, hours=12),
            'iat': datetime.datetime.utcnow(),
            'sub': {
                "user_id": user.id,
                "role": user.role
            }
        }
        return jwt.encode(
            payload,
            config.SECRET_KEY,
            algorithm='HS256'
        )
    except Exception as e:
        return e


def authentification_required(f):
    @wraps(f)
    def decorated():
        token = request.headers.get('Authorization')
        token = token.split(" ")[1]
        if not token:
            return jsonify({'message': "Token is missing!"}), 403

        try:
            token = jwt.decode(token, config.SECRET_KEY)
        except:
            return jsonify({"message": "Token is invalid!"}), 403

        return f(token)
    return decorated


def admin_required(f):
    @wraps(f)
    def decorated(token):
        print(token)
        if token['sub']['role'] != "ROLE_ADMIN":
            return jsonify({"message": "Unauthorized!"}), 401
        else:
            return f()

    return decorated
