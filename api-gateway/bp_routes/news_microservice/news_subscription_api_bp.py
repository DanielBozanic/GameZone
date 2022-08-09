from flask import Blueprint
from api.news_microservice.news_subscription_api import *
from utils.routes.news_microservice import news_subscription_api_routes

news_subscription_api_bp = Blueprint('news_subscription_api_bp', __name__)
news_subscription_api_bp.route(news_subscription_api_routes.SUBSCRIBE, methods=['POST'])(subscribe)
news_subscription_api_bp.route(news_subscription_api_routes.UNSUBSCRIBE, methods=['DELETE'])(unsubscribe)
news_subscription_api_bp.route(news_subscription_api_routes.IS_USER_SUBSCRIBED, methods=['GET'])(is_user_subscribed)
