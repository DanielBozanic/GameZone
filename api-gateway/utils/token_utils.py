from functools import wraps
from flask import request, jsonify
import jwt

SECRET_KEY = "7kRh7cgjun9S83Hu06HhqhB8sGYkZKHrZ7CRkpQJHfOzXTllQPcIWCSn3IcUccq"


def authentification_required(f):
    @wraps(f)
    def decorated(*args, **kwargs):
        token = request.headers.get('Authorization')
        if not token:
            return jsonify({'message': "Token is missing!"}), 403

        token = token.split(" ")[1]
        try:
             jwt.decode(token, SECRET_KEY)
        except:
            return jsonify({"message": "Token is invalid!"}), 403

        return f(*args, **kwargs)
    return decorated


def authorization_required(roles):
    def decorate(f):
        @wraps(f)
        def wrapper(*args, **kwargs):
            token = request.headers.get('Authorization')
            if not token:
                return jsonify({'message': "Token is missing!"}), 403

            token = token.split(" ")[1]
            try:
                token = jwt.decode(token, SECRET_KEY)
            except:
                return jsonify({"message": "Token is invalid!"}), 403

            if token['sub']['role'] not in roles:
                return jsonify({"message": "Unauthorized!"}), 401
            else:
                return f(*args, **kwargs)
        return wrapper
    return decorate
