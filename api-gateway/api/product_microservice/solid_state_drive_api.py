import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.product_microservice import solid_state_drive_api_routes


def get_all():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_ALL
                        + "?page={}&pageSize={}".format(page, page_size))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_NUMBER_OF_RECORDS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_by_id(id):
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_BY_ID + "/{}".format(id))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def search_by_name():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    name = args.get("name")
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.SEARCH_BY_NAME
                        + "?page={}&pageSize={}&name={}".format(page, page_size, name))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records_search():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_NUMBER_OF_RECORDS_SEARCH)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def filter():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    data = request.json
    r = requests.post(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                         solid_state_drive_api_routes.FILTER +
                         "?page={}&pageSize={}".format(page, page_size), json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records_filter():
    data = request.json
    r = requests.post(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                         solid_state_drive_api_routes.GET_NUMBER_OF_RECORDS_FILTER, json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_manufacturers():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_MANUFACTURERS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_forms():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_FORMS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_capacities():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_CAPACITIES)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_max_sequential_reads():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_MAX_SEQUENTIAL_READS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_max_sequential_writes():
    r = requests.get(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.GET_MAX_SEQUENTIAL_WRITES)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def create():
    data = request.json
    headers = request.headers
    r = requests.post(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                         solid_state_drive_api_routes.CREATE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def update():
    data = request.json
    headers = request.headers
    r = requests.put(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                        solid_state_drive_api_routes.UPDATE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete(id):
    headers = request.headers
    r = requests.delete(solid_state_drive_api_routes.BASE + solid_state_drive_api_routes.API +
                           solid_state_drive_api_routes.DELETE + "/{}".format(id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
