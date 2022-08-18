import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.contact_and_report_microservice import contact_api_routes


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def get_contact_messages():
    headers = request.headers
    try:
        r = requests.get(contact_api_routes.BASE + contact_api_routes.API +
                         contact_api_routes.GET_CONTACT_MESSAGES, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_ADMIN])
def get_contact_messages_by_user_id(userId):
    headers = request.headers
    try:
        r = requests.get(contact_api_routes.BASE + contact_api_routes.API +
                         contact_api_routes.GET_CONTACT_MESSAGES_BY_USER_ID + "/{}".format(userId),
                         headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_number_of_unanswered_contact_messages_by_user_id(userId):
    headers = request.headers
    try:
        r = requests.get(contact_api_routes.BASE + contact_api_routes.API +
                         contact_api_routes.GET_NUMBER_OF_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID + "/{}".format(userId),
                         headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE, role.ROLE_ADMIN])
def answer_contact_message():
    headers = request.headers
    data = request.json
    try:
        r = requests.put(contact_api_routes.BASE + contact_api_routes.API +
                         contact_api_routes.ANSWER_CONTACT_MESSAGE, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def send_contact_message():
    headers = request.headers
    data = request.json
    try:
        r = requests.post(contact_api_routes.BASE + contact_api_routes.API +
                          contact_api_routes.SEND_CONTACT_MESSAGE, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
