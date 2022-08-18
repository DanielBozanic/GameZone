import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.comment_and_rating_microservice import product_comment_api_routes


def get_all():
    try:
        r = requests.get(product_comment_api_routes.BASE + product_comment_api_routes.API +
                         product_comment_api_routes.GET_ALL)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_by_id(id):
    try:
        r = requests.get(product_comment_api_routes.BASE + product_comment_api_routes.API +
                         product_comment_api_routes.GET_BY_ID + "/{}".format(id))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_by_product_id(productId):
    try:
        r = requests.get(product_comment_api_routes.BASE + product_comment_api_routes.API +
                         product_comment_api_routes.GET_BY_PRODUCT_ID + "/{}".format(productId))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def add_comment():
    headers = request.headers
    data = request.json
    try:
        r = requests.post(product_comment_api_routes.BASE + product_comment_api_routes.API +
                          product_comment_api_routes.ADD_COMMENT, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def edit_comment():
    headers = request.headers
    data = request.json
    try:
        r = requests.put(product_comment_api_routes.BASE + product_comment_api_routes.API +
                         product_comment_api_routes.EDIT_COMMENT, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_ADMIN])
def delete_comment(id):
    headers = request.headers
    try:
        r = requests.delete(product_comment_api_routes.BASE + product_comment_api_routes.API +
                            product_comment_api_routes.DELETE_COMMENT + "/{}".format(id), headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete_comments_by_product_id(productId):
    headers = request.headers
    try:
        r = requests.delete(product_comment_api_routes.BASE + product_comment_api_routes.API +
                            product_comment_api_routes.DELETE_COMMENTS_BY_PRODUCT_ID + "/{}".format(productId),
                            headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_ADMIN])
def get_by_user_id(userId):
    headers = request.headers
    try:
        r = requests.get(product_comment_api_routes.BASE + product_comment_api_routes.API +
                         product_comment_api_routes.GET_BY_USER_ID + "/{}".format(userId), headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
