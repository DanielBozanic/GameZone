from flask import Blueprint
from api.product_microservice.motherboard_api import *
from utils.routes.product_microservice import motherboard_api_routes

motherboard_api_bp = Blueprint('motherboard_api_bp', __name__)
motherboard_api_bp.route(motherboard_api_routes.GET_ALL, methods=['GET'])(get_all)
motherboard_api_bp.route(motherboard_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
motherboard_api_bp.route(motherboard_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
motherboard_api_bp.route(motherboard_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
motherboard_api_bp.route(motherboard_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                         methods=['GET'])(get_number_of_records_search)
motherboard_api_bp.route(motherboard_api_routes.FILTER, methods=['POST'])(filter)
motherboard_api_bp.route(motherboard_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                         methods=['POST'])(get_number_of_records_filter)
motherboard_api_bp.route(motherboard_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
motherboard_api_bp.route(motherboard_api_routes.GET_PROCESSOR_TYPES, methods=['GET'])(get_processor_types)
motherboard_api_bp.route(motherboard_api_routes.GET_SOCKETS, methods=['GET'])(get_sockets)
motherboard_api_bp.route(motherboard_api_routes.GET_FORM_FACTORS, methods=['GET'])(get_form_factors)
motherboard_api_bp.route(motherboard_api_routes.CREATE, methods=['POST'])(create)
motherboard_api_bp.route(motherboard_api_routes.UPDATE, methods=['PUT'])(update)
motherboard_api_bp.route(motherboard_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
