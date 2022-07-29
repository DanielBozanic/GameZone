from flask import request, jsonify
import services.user_service
import utils.token_utils


def register():
    msg = services.user_service.register(request.json)
    if msg == "":
        resp = jsonify(message="Check email for your next step")
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp


def get_verification_code():
    args = request.args.to_dict()
    email = args.get("email")
    msg, status_code = services.user_service.get_verification_code(email)
    resp = jsonify(message=msg)
    resp.status_code = status_code
    return resp


def verify_account():
    msg = services.user_service.verify_account(request.json)
    if msg == "":
        resp = jsonify(message="Account verified")
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
        if request.json['role'] == "ROLE_EMPLOYEE":
            msg = "Employee successfully added."
        elif request.json['role'] == "ROLE_ADMIN":
            msg = "Admin successfully added."
        resp = jsonify(message=msg)
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=msg)
        resp.status_code = 400
        return resp


def get_by_id():
    args = request.args.to_dict()
    user_id = args.get("userId")
    user = services.user_service.get_by_id(user_id)
    if not isinstance(user, str):
        resp = jsonify(user=user.serialize())
        resp.status_code = 200
        return resp
    else:
        resp = jsonify(message=user)
        resp.status_code = 400
        return resp


@utils.token_utils.authentification_required
@utils.token_utils.roles_required(roles=["ROLE_ADMIN"])
def get_all_registered_users():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    users = services.user_service.get_all_registered_users(int(page), int(page_size))
    resp = jsonify(users=[user.serialize() for user in users.items])
    resp.status_code = 200
    return resp


@utils.token_utils.authentification_required
@utils.token_utils.roles_required(roles=["ROLE_ADMIN"])
def get_number_of_records_registered_users():
    count = services.user_service.get_number_of_records_registered_users()
    resp = jsonify(count=count)
    resp.status_code = 200
    return resp
