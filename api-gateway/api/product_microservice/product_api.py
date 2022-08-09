import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.product_microservice import product_api_routes


def get_product_by_id(id):
    r = requests.get(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.GET_PRODUCT_BY_ID + "/{}".format(id))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def search_by_name():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    name = args.get("name")
    r = requests.get(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.SEARCH_BY_NAME
                        + "?page={}&pageSize={}&name={}".format(page, page_size, name))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records_search():
    r = requests.get(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.GET_NUMBER_OF_RECORDS_SEARCH)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_main_page_products():
    r = requests.get(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.GET_MAIN_PAGE_PRODUCTS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_popular_products():
    r = requests.get(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.GET_POPULAR_PRODUCTS)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def add_product_to_main_page(productId):
    headers = request.headers
    r = requests.put(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.ADD_PRODUCT_TO_MAIN_PAGE + "/{}".format(productId), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def remove_product_from_main_page(productId):
    headers = request.headers
    r = requests.delete(product_api_routes.BASE + product_api_routes.API +
                           product_api_routes.REMOVE_PRODUCT_FROM_MAIN_PAGE + "/{}".format(productId), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete_product(id):
    headers = request.headers
    r = requests.delete(product_api_routes.BASE + product_api_routes.API +
                           product_api_routes.DELETE_PRODUCT + "/{}".format(id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def get_recommended_products():
    headers = request.headers
    r = requests.get(product_api_routes.BASE + product_api_routes.API +
                        product_api_routes.GET_RECOMMENDED_PRODUCTS, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
