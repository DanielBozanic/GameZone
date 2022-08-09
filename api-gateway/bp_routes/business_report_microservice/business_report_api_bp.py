from flask import Blueprint
from api.business_report_microservice.business_report_api import *
from utils.routes.business_report_microservice import business_report_api_routes

business_report_api_bp = Blueprint('business_report_api_bp', __name__)
business_report_api_bp.route(business_report_api_routes.GET_PRODUCTS_WITH_BIGGEST_PROFIT_LAST_THIRTY_DAYS,
                             methods=['GET'])(get_products_with_biggest_profit_last_thirty_days)
business_report_api_bp.route(business_report_api_routes.GET_SOLD_VIDEO_GAMES_BY_PLATFORM,
                             methods=['GET'])(get_sold_video_games_by_platform)
business_report_api_bp.route(business_report_api_routes.GET_SOLD_VIDEO_GAMES_BY_FORM,
                             methods=['GET'])(get_sold_video_games_by_form)
