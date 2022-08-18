import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.contact_and_report_microservice import ban_api_routes


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_user_ban_history(userId):
    headers = request.headers
    try:
        r = requests.get(ban_api_routes.BASE + ban_api_routes.API +
                         ban_api_routes.GET_USER_BAN_HISTORY + "/{}".format(userId), headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def add_ban():
    headers = request.headers
    data = request.json
    try:
        r = requests.post(ban_api_routes.BASE + ban_api_routes.API +
                          ban_api_routes.ADD_BAN, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def send_email_to_banned_user():
    headers = request.headers
    data = request.json
    try:
        r = requests.post(ban_api_routes.BASE + ban_api_routes.API +
                          ban_api_routes.SEND_EMAIL_TO_BANNED_USER, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
