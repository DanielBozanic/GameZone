import requests
from flask import request, jsonify
from utils import token_utils, role
from utils.routes.news_microservice import news_article_api_routes


def get_published_news_articles():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    r = requests.get(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.GET_PUBLISHED_ARTICLES +
                        "?page={}&pageSize={}".format(page, page_size))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_number_of_records_published_news_articles():
    r = requests.get(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.GET_NUMBER_OF_RECORDS_PUBLISHED_ARTICLES)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


def get_by_id(id):
    r = requests.get(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.GET_BY_ID + "/{}".format(id))
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def get_all():
    args = request.args.to_dict()
    page = args.get("page")
    page_size = args.get("pageSize")
    headers = request.headers
    r = requests.get(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.GET_ALL
                        + "?page={}&pageSize={}".format(page, page_size), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def get_number_of_records():
    headers = request.headers
    r = requests.get(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.GET_NUMBER_OF_RECORDS, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def add_news_article():
    data = request.json
    headers = request.headers
    r = requests.post(news_article_api_routes.BASE + news_article_api_routes.API +
                         news_article_api_routes.ADD_NEWS_ARTICLE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def edit_news_article():
    data = request.json
    headers = request.headers
    r = requests.put(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.EDIT_NEWS_ARTICLE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def delete_news_article(id):
    headers = request.headers
    r = requests.delete(news_article_api_routes.BASE + news_article_api_routes.API +
                           news_article_api_routes.DELETE_NEWS_ARTICLE + "/{}".format(id), headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp


@token_utils.authorization_required(roles=[role.ROLE_EMPLOYEE])
def publish_news_article():
    data = request.json
    headers = request.headers
    r = requests.put(news_article_api_routes.BASE + news_article_api_routes.API +
                        news_article_api_routes.PUBLISH_NEWS_ARTICLE, json=data, headers=headers)
    resp = jsonify(r.json())
    resp.status_code = r.status_code
    return resp
