from flask import request, jsonify
import services.user_service


def create():
    msg = services.user_service.create(request.json)
    if msg == "":
        resp = jsonify(message="User successfully created.")
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp


def get_all():
    users = services.user_service.get_all()
    resp = jsonify(users=[user.serialize() for user in users])
    resp.status_code = 200
    return resp


def get_by_id(user_id):
    user = services.user_service.get_by_id(user_id)
    if isinstance(user, str):
        resp = jsonify(user=user)
        resp.status_code = 400
        return resp
    else:
        resp = jsonify(user=user.serialize())
        resp.status_code = 200
        return resp


def update():
    msg = services.user_service.update(request.json)
    if msg == "":
        resp = jsonify(message="User successfully updated.")
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp


def delete(user_id):
    msg = services.user_service.delete(user_id)
    if msg == "":
        resp = jsonify(message="User successfully deleted.")
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp
