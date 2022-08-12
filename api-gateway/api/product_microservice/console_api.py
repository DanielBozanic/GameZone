import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.product_microservice import console_api_routes


def get_all():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    r = requests.get(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.GET_ALL
                        + "?page={}&pageSize={}".format(page, page_size))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records():
    r = requests.get(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.GET_NUMBER_OF_RECORDS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_by_id(id):
    r = requests.get(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.GET_BY_ID + "/{}".format(id))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def search_by_name():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    name = args.get("name")
    r = requests.get(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.SEARCH_BY_NAME
                        + "?page={}&pageSize={}&name={}".format(page, page_size, name))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records_search():
    args = request.args.to_dict()
    name = args.get("name")
    r = requests.get(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.GET_NUMBER_OF_RECORDS_SEARCH + "?name={}".format(name))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def filter():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    data = request.json
    r = requests.post(console_api_routes.BASE + console_api_routes.API +
                         console_api_routes.FILTER +
                         "?page={}&pageSize={}".format(page, page_size), json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records_filter():
    data = request.json
    r = requests.post(console_api_routes.BASE + console_api_routes.API +
                         console_api_routes.GET_NUMBER_OF_RECORDS_FILTER, json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_platforms():
    r = requests.get(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.GET_PLATFORMS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def create():
    data = request.json
    headers = request.headers
    r = requests.post(console_api_routes.BASE + console_api_routes.API +
                         console_api_routes.CREATE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def update():
    data = request.json
    headers = request.headers
    r = requests.put(console_api_routes.BASE + console_api_routes.API +
                        console_api_routes.UPDATE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete(id):
    headers = request.headers
    r = requests.delete(console_api_routes.BASE + console_api_routes.API +
                           console_api_routes.DELETE + "/{}".format(id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
