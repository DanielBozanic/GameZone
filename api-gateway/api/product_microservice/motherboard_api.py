import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.product_microservice import motherboard_api_routes


def get_all():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API + motherboard_api_routes.GET_ALL +
                         "?page={}&pageSize={}".format(page, page_size))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_number_of_records():
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_NUMBER_OF_RECORDS)
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
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_BY_ID + "/{}".format(id))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def search_by_name():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    name = args.get("name")
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.SEARCH_BY_NAME +
                         "?page={}&pageSize={}&name={}".format(page, page_size, name))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_number_of_records_search():
    args = request.args.to_dict()
    name = args.get("name")
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_NUMBER_OF_RECORDS_SEARCH + "?name={}".format(name))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def filter():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    data = request.json
    try:
        r = requests.post(motherboard_api_routes.BASE + motherboard_api_routes.API +
                          motherboard_api_routes.FILTER +
                          "?page={}&pageSize={}".format(page, page_size), json=data)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_number_of_records_filter():
    data = request.json
    try:
        r = requests.post(motherboard_api_routes.BASE + motherboard_api_routes.API +
                          motherboard_api_routes.GET_NUMBER_OF_RECORDS_FILTER, json=data)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_manufacturers():
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_MANUFACTURERS)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_processor_types():
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_PROCESSOR_TYPES)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_sockets():
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_SOCKETS)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


def get_form_factors():
    try:
        r = requests.get(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.GET_FORM_FACTORS)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def create():
    data = request.json
    headers = request.headers
    try:
        r = requests.post(motherboard_api_routes.BASE + motherboard_api_routes.API +
                          motherboard_api_routes.CREATE, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def update():
    data = request.json
    headers = request.headers
    try:
        r = requests.put(motherboard_api_routes.BASE + motherboard_api_routes.API +
                         motherboard_api_routes.UPDATE, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete(id):
    headers = request.headers
    try:
        r = requests.delete(motherboard_api_routes.BASE + motherboard_api_routes.API +
                            motherboard_api_routes.DELETE + "/{}".format(id), headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
