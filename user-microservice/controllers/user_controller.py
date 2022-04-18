from flask import request, jsonify
import services.user_service
import utils.token_utils


def register():
    msg = services.user_service.register(request.json)
    if msg == "":
        resp = jsonify(message="User successfully registered.")
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp


@utils.token_utils.authentification_required
@utils.token_utils.roles_required(roles=["ROLE_ADMIN"])
def add_employee_and_admin():
    msg = services.user_service.add_employee_and_admin(request.json)
    if msg == "":
        resp = jsonify(message="Employee successfully added.")
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp
