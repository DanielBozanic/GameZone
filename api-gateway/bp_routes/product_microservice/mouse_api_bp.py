from flask import Blueprint
from api.product_microservice.mouse_api import *
from utils.routes.product_microservice import mouse_api_routes

mouse_api_bp = Blueprint('mouse_api_bp', __name__)
mouse_api_bp.route(mouse_api_routes.GET_ALL, methods=['GET'])(get_all)
mouse_api_bp.route(mouse_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
mouse_api_bp.route(mouse_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
mouse_api_bp.route(mouse_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
mouse_api_bp.route(mouse_api_routes.GET_NUMBER_OF_RECORDS_SEARCH, methods=['GET'])(get_number_of_records_search)
mouse_api_bp.route(mouse_api_routes.FILTER, methods=['POST'])(filter)
mouse_api_bp.route(mouse_api_routes.GET_NUMBER_OF_RECORDS_FILTER, methods=['POST'])(get_number_of_records_filter)
mouse_api_bp.route(mouse_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
mouse_api_bp.route(mouse_api_routes.GET_DPIS, methods=['GET'])(get_dpis)
mouse_api_bp.route(mouse_api_routes.GET_CONNECTIONS, methods=['GET'])(get_connections)
mouse_api_bp.route(mouse_api_routes.CREATE, methods=['POST'])(create)
mouse_api_bp.route(mouse_api_routes.UPDATE, methods=['PUT'])(update)
mouse_api_bp.route(mouse_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
