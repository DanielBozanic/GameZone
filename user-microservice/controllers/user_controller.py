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
    pass


def get_by_id(user_id):
    pass


def update(user_id):
    pass


def delete(user_id):
    pass