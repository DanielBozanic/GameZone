from flask import Blueprint
from api.product_microservice.ram_api import *
from utils.routes.product_microservice import ram_api_routes

ram_api_bp = Blueprint('ram_api_bp', __name__)
ram_api_bp.route(ram_api_routes.GET_ALL, methods=['GET'])(get_all)
ram_api_bp.route(ram_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
ram_api_bp.route(ram_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
ram_api_bp.route(ram_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
ram_api_bp.route(ram_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                 methods=['GET'])(get_number_of_records_search)
ram_api_bp.route(ram_api_routes.FILTER, methods=['POST'])(filter)
ram_api_bp.route(ram_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                 methods=['POST'])(get_number_of_records_filter)
ram_api_bp.route(ram_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
ram_api_bp.route(ram_api_routes.GET_CAPACITIES, methods=['GET'])(get_capacities)
ram_api_bp.route(ram_api_routes.GET_MEMORY_TYPES, methods=['GET'])(get_memory_types)
ram_api_bp.route(ram_api_routes.GET_SPEEDS, methods=['GET'])(get_speeds)
ram_api_bp.route(ram_api_routes.CREATE, methods=['POST'])(create)
ram_api_bp.route(ram_api_routes.UPDATE, methods=['PUT'])(update)
ram_api_bp.route(ram_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
