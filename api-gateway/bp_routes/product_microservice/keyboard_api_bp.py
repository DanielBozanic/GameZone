from flask import Blueprint
from api.product_microservice.keyboard_api import *
from utils.routes.product_microservice import keyboard_api_routes

keyboard_api_bp = Blueprint('keyboard_api_bp', __name__)
keyboard_api_bp.route(keyboard_api_routes.GET_ALL, methods=['GET'])(get_all)
keyboard_api_bp.route(keyboard_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
keyboard_api_bp.route(keyboard_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
keyboard_api_bp.route(keyboard_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
keyboard_api_bp.route(keyboard_api_routes.GET_NUMBER_OF_RECORDS_SEARCH, methods=['GET'])(get_number_of_records_search)
keyboard_api_bp.route(keyboard_api_routes.FILTER, methods=['POST'])(filter)
keyboard_api_bp.route(keyboard_api_routes.GET_NUMBER_OF_RECORDS_FILTER, methods=['POST'])(get_number_of_records_filter)
keyboard_api_bp.route(keyboard_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
keyboard_api_bp.route(keyboard_api_routes.GET_KEYBOARD_CONNECTORS, methods=['GET'])(get_keyboard_connectors)
keyboard_api_bp.route(keyboard_api_routes.GET_KEY_TYPES, methods=['GET'])(get_key_types)
keyboard_api_bp.route(keyboard_api_routes.CREATE, methods=['POST'])(create)
keyboard_api_bp.route(keyboard_api_routes.UPDATE, methods=['PUT'])(update)
keyboard_api_bp.route(keyboard_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
