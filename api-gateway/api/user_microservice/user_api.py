import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.user_microservice import user_api_routes


def register():
    data = request.json
    r = requests.post(user_api_routes.BASE + user_api_routes.API + user_api_routes.REGISTER, json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_verification_code():
    args = request.args.to_dict()
    email = args.get("email")
    r = requests.get(user_api_routes.BASE + user_api_routes.API + user_api_routes.GET_VERIFICATION_CODE +
                        "?email={}".format(email))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def verify_account():
    data = request.json
    r = requests.put(user_api_routes.BASE + user_api_routes.API + user_api_routes.VERIFY_ACCOUNT, json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def add_employee_and_admin():
    data = request.json
    headers = request.headers
    r = requests.post(user_api_routes.BASE + user_api_routes.API + user_api_routes.ADD_EMPLOYEE_AND_ADMIN,
                         json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_by_id():
    args = request.args.to_dict()
    user_id = args.get("userId")
    r = requests.get(user_api_routes.BASE + user_api_routes.API + user_api_routes.GET_USER_BY_ID +
                        "?userId={}".format(user_id))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_all_registered_users():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    headers = request.headers
    r = requests.get(user_api_routes.BASE + user_api_routes.API +
                        user_api_routes.GET_ALL_REGISTERED_USERS
                        + "?page={}&pageSize={}".format(page, page_size), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_number_of_records_registered_users():
    headers = request.headers
    r = requests.get(user_api_routes.BASE + user_api_routes.API +
                        user_api_routes.GET_NUMBER_OF_RECORDS_REGISTERED_USERS, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authentification_required
def update():
    data = request.json
    headers = request.headers
    r = requests.put(user_api_routes.BASE + user_api_routes.API + user_api_routes.UPDATE,
                        json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authentification_required
def change_password():
    data = request.json
    headers = request.headers
    r = requests.put(user_api_routes.BASE + user_api_routes.API + user_api_routes.CHANGE_PASSWORD,
                        json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
