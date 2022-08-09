import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.contact_and_report_microservice import report_api_routes


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_EMPLOYEE])
def add_report():
    headers = request.headers
    data = request.json
    r = requests.post(report_api_routes.BASE + report_api_routes.API +
                         report_api_routes.ADD_REPORT, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_ADMIN])
def get_reports_by_user_id(userId):
    headers = request.headers
    r = requests.get(report_api_routes.BASE + report_api_routes.API +
                        report_api_routes.GET_REPORTS_BY_USER_ID + "/{}".format(userId), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
