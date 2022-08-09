from flask import Blueprint
from api.product_microservice.console_api import *
from utils.routes.product_microservice import console_api_routes

console_api_bp = Blueprint('console_api_bp', __name__)
console_api_bp.route(console_api_routes.GET_ALL, methods=['GET'])(get_all)
console_api_bp.route(console_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
console_api_bp.route(console_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
console_api_bp.route(console_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
console_api_bp.route(console_api_routes.GET_NUMBER_OF_RECORDS_SEARCH, methods=['GET'])(get_number_of_records_search)
console_api_bp.route(console_api_routes.FILTER, methods=['POST'])(filter)
console_api_bp.route(console_api_routes.GET_NUMBER_OF_RECORDS_FILTER, methods=['POST'])(get_number_of_records_filter)
console_api_bp.route(console_api_routes.GET_PLATFORMS, methods=['GET'])(get_platforms)
console_api_bp.route(console_api_routes.CREATE, methods=['POST'])(create)
console_api_bp.route(console_api_routes.UPDATE, methods=['PUT'])(update)
console_api_bp.route(console_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
