from flask import Blueprint
from api.news_microservice.news_article_api import *
from utils.routes.news_microservice import news_article_api_routes

news_article_api_bp = Blueprint('news_article_api_bp', __name__)
news_article_api_bp.route(news_article_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
news_article_api_bp.route(news_article_api_routes.GET_PUBLISHED_ARTICLES, methods=['GET'])(get_published_news_articles)
news_article_api_bp.route(news_article_api_routes.GET_NUMBER_OF_RECORDS_PUBLISHED_ARTICLES,
                          methods=['GET'])(get_number_of_records_published_news_articles)
news_article_api_bp.route(news_article_api_routes.GET_ALL, methods=['GET'])(get_all)
news_article_api_bp.route(news_article_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
news_article_api_bp.route(news_article_api_routes.ADD_NEWS_ARTICLE, methods=['POST'])(add_news_article)
news_article_api_bp.route(news_article_api_routes.EDIT_NEWS_ARTICLE, methods=['PUT'])(edit_news_article)
news_article_api_bp.route(news_article_api_routes.DELETE_NEWS_ARTICLE + "/<id>",
                          methods=['DELETE'])(delete_news_article)
news_article_api_bp.route(news_article_api_routes.PUBLISH_NEWS_ARTICLE, methods=['PUT'])(publish_news_article)
