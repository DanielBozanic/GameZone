from flask import Blueprint
from api.product_microservice.headphones_api import *
from utils.routes.product_microservice import headphones_api_routes

headphones_api_bp = Blueprint('headphones_api_bp', __name__)
headphones_api_bp.route(headphones_api_routes.GET_ALL, methods=['GET'])(get_all)
headphones_api_bp.route(headphones_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
headphones_api_bp.route(headphones_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
headphones_api_bp.route(headphones_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
headphones_api_bp.route(headphones_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                        methods=['GET'])(get_number_of_records_search)
headphones_api_bp.route(headphones_api_routes.FILTER, methods=['POST'])(filter)
headphones_api_bp.route(headphones_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                        methods=['POST'])(get_number_of_records_filter)
headphones_api_bp.route(headphones_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
headphones_api_bp.route(headphones_api_routes.GET_CONNECTION_TYPES, methods=['GET'])(get_connection_types)
headphones_api_bp.route(headphones_api_routes.CREATE, methods=['POST'])(create)
headphones_api_bp.route(headphones_api_routes.UPDATE, methods=['PUT'])(update)
headphones_api_bp.route(headphones_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
