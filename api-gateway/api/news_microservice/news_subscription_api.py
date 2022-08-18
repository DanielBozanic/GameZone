import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.news_microservice import news_subscription_api_routes


@token_utils.authorization_required(roles=[role.ROLE_USER])
def subscribe():
    headers = request.headers
    try:
        r = requests.post(news_subscription_api_routes.BASE + news_subscription_api_routes.API +
                          news_subscription_api_routes.SUBSCRIBE, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def unsubscribe():
    headers = request.headers
    try:
        r = requests.delete(news_subscription_api_routes.BASE + news_subscription_api_routes.API +
                            news_subscription_api_routes.UNSUBSCRIBE, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def is_user_subscribed():
    headers = request.headers
    try:
        r = requests.get(news_subscription_api_routes.BASE + news_subscription_api_routes.API +
                         news_subscription_api_routes.IS_USER_SUBSCRIBED, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
