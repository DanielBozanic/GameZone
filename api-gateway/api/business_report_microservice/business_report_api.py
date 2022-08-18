import requests
from flask import jsonify
from utils import token_utils, role
from utils.routes.business_report_microservice import business_report_api_routes


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_products_with_biggest_profit_last_thirty_days():
    try:
        r = requests.get(business_report_api_routes.BASE + business_report_api_routes.API +
                         business_report_api_routes.GET_PRODUCTS_WITH_BIGGEST_PROFIT_LAST_THIRTY_DAYS)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_sold_video_games_by_platform():
    try:
        r = requests.get(business_report_api_routes.BASE + business_report_api_routes.API +
                         business_report_api_routes.GET_SOLD_VIDEO_GAMES_BY_PLATFORM)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_sold_video_games_by_form():
    try:
        r = requests.get(business_report_api_routes.BASE + business_report_api_routes.API +
                         business_report_api_routes.GET_SOLD_VIDEO_GAMES_BY_FORM)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
