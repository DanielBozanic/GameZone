import requests
from flask import request, jsonify
from utils.routes.user_microservice import auth_api_routes


def login():
    data = request.json
    try:
        r = requests.post(auth_api_routes.BASE + auth_api_routes.API + auth_api_routes.LOGIN, json=data)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
