from flask import Blueprint
from api.news_microservice.news_comment_api import *
from utils.routes.news_microservice import news_comment_api_routes

news_comment_api_bp = Blueprint('news_comment_api_bp', __name__)
news_comment_api_bp.route(news_comment_api_routes.GET_BY_NEWS_ARTICLE + "/<newsArticleId>",
                          methods=['GET'])(get_by_news_article)
news_comment_api_bp.route(news_comment_api_routes.ADD_NEWS_COMMENT, methods=['POST'])(add_news_comment)
news_comment_api_bp.route(news_comment_api_routes.EDIT_NEWS_COMMENT, methods=['PUT'])(edit_news_comment)
news_comment_api_bp.route(news_comment_api_routes.DELETE_NEWS_COMMENT + "/<id>",
                          methods=['DELETE'])(delete_news_comment)
news_comment_api_bp.route(news_comment_api_routes.DELETE_NEWS_COMMENTS_BY_NEWS_ARTICLE_ID + "/<newsArticleId>",
                          methods=['DELETE'])(delete_news_comments_by_news_article_id)
news_comment_api_bp.route(news_comment_api_routes.GET_BY_USER_ID + "/<userId>",
                          methods=['GET'])(get_by_user_id)
