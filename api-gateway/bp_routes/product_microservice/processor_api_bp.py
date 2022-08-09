from flask import Blueprint
from api.product_microservice.processor_api import *
from utils.routes.product_microservice import processor_api_routes

processor_api_bp = Blueprint('processor_api_bp', __name__)
processor_api_bp.route(processor_api_routes.GET_ALL, methods=['GET'])(get_all)
processor_api_bp.route(processor_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
processor_api_bp.route(processor_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
processor_api_bp.route(processor_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
processor_api_bp.route(processor_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                       methods=['GET'])(get_number_of_records_search)
processor_api_bp.route(processor_api_routes.FILTER, methods=['POST'])(filter)
processor_api_bp.route(processor_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                       methods=['POST'])(get_number_of_records_filter)
processor_api_bp.route(processor_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
processor_api_bp.route(processor_api_routes.GET_TYPES, methods=['GET'])(get_types)
processor_api_bp.route(processor_api_routes.GET_SOCKETS, methods=['GET'])(get_sockets)
processor_api_bp.route(processor_api_routes.GET_NUMBER_OF_CORES, methods=['GET'])(get_number_of_cores)
processor_api_bp.route(processor_api_routes.GET_THREADS, methods=['GET'])(get_threads)
processor_api_bp.route(processor_api_routes.CREATE, methods=['POST'])(create)
processor_api_bp.route(processor_api_routes.UPDATE, methods=['PUT'])(update)
processor_api_bp.route(processor_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
