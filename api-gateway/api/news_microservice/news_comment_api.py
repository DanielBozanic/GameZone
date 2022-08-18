import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.news_microservice import news_comment_api_routes


def get_by_news_article(newsArticleId):
    try:
        r = requests.get(news_comment_api_routes.BASE + news_comment_api_routes.API +
                         news_comment_api_routes.GET_BY_NEWS_ARTICLE + "/{}".format(newsArticleId))
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def add_news_comment():
    data = request.json
    headers = request.headers
    try:
        r = requests.post(news_comment_api_routes.BASE + news_comment_api_routes.API +
                          news_comment_api_routes.ADD_NEWS_COMMENT, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER])
def edit_news_comment():
    data = request.json
    headers = request.headers
    try:
        r = requests.put(news_comment_api_routes.BASE + news_comment_api_routes.API +
                         news_comment_api_routes.EDIT_NEWS_COMMENT, json=data, headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_USER, role.ROLE_ADMIN])
def delete_news_comment(id):
    headers = request.headers
    try:
        r = requests.delete(news_comment_api_routes.BASE + news_comment_api_routes.API +
                            news_comment_api_routes.DELETE_NEWS_COMMENT + "/{}".format(id), headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete_news_comments_by_news_article_id(newsArticleId):
    headers = request.headers
    try:
        r = requests.delete(news_comment_api_routes.BASE + news_comment_api_routes.API +
                            news_comment_api_routes.DELETE_NEWS_COMMENTS_BY_NEWS_ARTICLE_ID +
                            "/{}".format(newsArticleId), headers=headers)
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
        r = requests.get(news_comment_api_routes.BASE + news_comment_api_routes.API +
                         news_comment_api_routes.GET_BY_USER_ID + "/{}".format(userId), headers=headers)
        resp = jsonify(r.json())
        resp.status_code = r.status_code
        return resp
    except requests.exceptions.RequestException as err:
        print(err)
        resp = jsonify(str(err))
        resp.status_code = 404
        return resp
