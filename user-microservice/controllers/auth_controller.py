from flask import jsonify, request
import services.auth_service
import utils.token_utils


def login():
    user = services.auth_service.login(request.json)
    if not isinstance(user, str):
        token = utils.token_utils.encode_auth_token(user)
        resp = jsonify({
            "token": token.decode(),
            "user": user.serialize()
        })
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=user)
        resp.status_code = 400
        return resp
