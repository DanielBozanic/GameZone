import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.product_microservice import product_purchase_api_routes


@token_utils.authorization_required(roles=[role.ROLE_USER])
def check_if_product_is_paid_for():
    args = request.args.to_dict()
    product_id = args.get("productId")
    headers = request.headers
    r = requests.get(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                        product_purchase_api_routes.CHECK_IF_PRODUCT_IS_PAID_FOR +
                        "?productId={}".format(product_id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def confirm_purchase():
    data = request.json
    headers = request.headers
    r = requests.post(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                         product_purchase_api_routes.CONFIRM_PURCHASE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def send_purchase_confirmation_mail():
    data = request.json
    headers = request.headers
    r = requests.post(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                         product_purchase_api_routes.SEND_PURCHASE_CONFIRMATION_MAIL, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def get_product_alert_by_product_id_and_user_id():
    args = request.args.to_dict()
    product_id = args.get("productId")
    headers = request.headers
    r = requests.get(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                        product_purchase_api_routes.GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_USER_ID +
                        "?productId={}".format(product_id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def add_product_alert():
    args = request.args.to_dict()
    product_id = args.get("productId")
    headers = request.headers
    r = requests.post(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                         product_purchase_api_routes.ADD_PRODUCT_ALERT +
                         "?productId={}".format(product_id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def notify_product_availability():
    args = request.args.to_dict()
    product_id = args.get("productId")
    headers = request.headers
    r = requests.get(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                        product_purchase_api_routes.NOTIFY_PRODUCT_AVAILABILITY +
                        "?productId={}".format(product_id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def confirm_payment():
    data = request.json
    headers = request.headers
    r = requests.put(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                        product_purchase_api_routes.CONFIRM_PAYMENT, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def send_purchased_digital_video_games():
    data = request.json
    headers = request.headers
    r = requests.post(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                         product_purchase_api_routes.SEND_PURCHASED_DIGITAL_VIDEO_GAMES, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_ADMIN])
def get_purchase_history():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    user_id = args.get("userId")
    headers = request.headers
    r = requests.get(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                        product_purchase_api_routes.GET_PURCHASE_HISTORY
                        + "?userId={}&page={}&pageSize={}".format(user_id, page, page_size), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_ADMIN])
def get_number_of_records_purchase_history():
    args = request.args.to_dict()
    user_id = args.get("userId")
    headers = request.headers
    r = requests.get(product_purchase_api_routes.BASE + product_purchase_api_routes.API +
                        product_purchase_api_routes.GET_NUMBER_OF_RECORDS_PURCHASE_HISTORY +
                        "?userId={}".format(user_id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
