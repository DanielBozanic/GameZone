import requests
from flask import request, jsonify
from utils.routes.user_microservice import auth_api_routes


def login():
    data = request.json
    r = requests.post(auth_api_routes.BASE + auth_api_routes.API + auth_api_routes.LOGIN, json=data)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp

