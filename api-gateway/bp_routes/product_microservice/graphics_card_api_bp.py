from flask import Blueprint
from api.product_microservice.graphics_card_api import *
from utils.routes.product_microservice import graphics_card_api_routes

graphics_card_api_bp = Blueprint('graphics_card_api_bp', __name__)
graphics_card_api_bp.route(graphics_card_api_routes.GET_ALL, methods=['GET'])(get_all)
graphics_card_api_bp.route(graphics_card_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
graphics_card_api_bp.route(graphics_card_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
graphics_card_api_bp.route(graphics_card_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
graphics_card_api_bp.route(graphics_card_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                           methods=['GET'])(get_number_of_records_search)
graphics_card_api_bp.route(graphics_card_api_routes.FILTER, methods=['POST'])(filter)
graphics_card_api_bp.route(graphics_card_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                           methods=['POST'])(get_number_of_records_filter)
graphics_card_api_bp.route(graphics_card_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
graphics_card_api_bp.route(graphics_card_api_routes.GET_CHIP_MANUFACTURERS, methods=['GET'])(get_chip_manufacturers)
graphics_card_api_bp.route(graphics_card_api_routes.GET_MEMORY_SIZES, methods=['GET'])(get_memory_sizes)
graphics_card_api_bp.route(graphics_card_api_routes.GET_MEMORY_TYPES, methods=['GET'])(get_memory_types)
graphics_card_api_bp.route(graphics_card_api_routes.GET_MODEL_NAMES, methods=['GET'])(get_model_names)
graphics_card_api_bp.route(graphics_card_api_routes.CREATE, methods=['POST'])(create)
graphics_card_api_bp.route(graphics_card_api_routes.UPDATE, methods=['PUT'])(update)
graphics_card_api_bp.route(graphics_card_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
