from flask import Blueprint
from api.product_microservice.monitor_api import *
from utils.routes.product_microservice import monitor_api_routes

monitor_api_bp = Blueprint('monitor_api_bp', __name__)
monitor_api_bp.route(monitor_api_routes.GET_ALL, methods=['GET'])(get_all)
monitor_api_bp.route(monitor_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
monitor_api_bp.route(monitor_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
monitor_api_bp.route(monitor_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
monitor_api_bp.route(monitor_api_routes.GET_NUMBER_OF_RECORDS_SEARCH, methods=['GET'])(get_number_of_records_search)
monitor_api_bp.route(monitor_api_routes.FILTER, methods=['POST'])(filter)
monitor_api_bp.route(monitor_api_routes.GET_NUMBER_OF_RECORDS_FILTER, methods=['POST'])(get_number_of_records_filter)
monitor_api_bp.route(monitor_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
monitor_api_bp.route(monitor_api_routes.GET_ASPECT_RATIOS, methods=['GET'])(get_aspect_ratios)
monitor_api_bp.route(monitor_api_routes.GET_RESOLUTIONS, methods=['GET'])(get_resolutions)
monitor_api_bp.route(monitor_api_routes.GET_REFRESH_RATES, methods=['GET'])(get_refresh_rates)
monitor_api_bp.route(monitor_api_routes.CREATE, methods=['POST'])(create)
monitor_api_bp.route(monitor_api_routes.UPDATE, methods=['PUT'])(update)
monitor_api_bp.route(monitor_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
