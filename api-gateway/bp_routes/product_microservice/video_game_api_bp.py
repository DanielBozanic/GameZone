from flask import Blueprint
from api.product_microservice.video_game_api import *
from utils.routes.product_microservice import video_game_api_routes

video_game_api_bp = Blueprint('video_game_api_bp', __name__)
video_game_api_bp.route(video_game_api_routes.GET_ALL, methods=['GET'])(get_all)
video_game_api_bp.route(video_game_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
video_game_api_bp.route(video_game_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
video_game_api_bp.route(video_game_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
video_game_api_bp.route(video_game_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                        methods=['GET'])(get_number_of_records_search)
video_game_api_bp.route(video_game_api_routes.FILTER, methods=['POST'])(filter)
video_game_api_bp.route(video_game_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                        methods=['POST'])(get_number_of_records_filter)
video_game_api_bp.route(video_game_api_routes.GET_PLATFORMS, methods=['GET'])(get_platforms)
video_game_api_bp.route(video_game_api_routes.GET_GENRES, methods=['GET'])(get_genres)
video_game_api_bp.route(video_game_api_routes.CREATE, methods=['POST'])(create)
video_game_api_bp.route(video_game_api_routes.UPDATE, methods=['PUT'])(update)
video_game_api_bp.route(video_game_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
