from flask import Blueprint
from api.product_microservice.power_supply_unit_api import *
from utils.routes.product_microservice import psu_api_routes

power_supply_unit_api_bp = Blueprint('power_supply_unit_api_bp', __name__)
power_supply_unit_api_bp.route(psu_api_routes.GET_ALL, methods=['GET'])(get_all)
power_supply_unit_api_bp.route(psu_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
power_supply_unit_api_bp.route(psu_api_routes.GET_NUMBER_OF_RECORDS, methods=['GET'])(get_number_of_records)
power_supply_unit_api_bp.route(psu_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
power_supply_unit_api_bp.route(psu_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                               methods=['GET'])(get_number_of_records_search)
power_supply_unit_api_bp.route(psu_api_routes.FILTER, methods=['POST'])(filter)
power_supply_unit_api_bp.route(psu_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                               methods=['POST'])(get_number_of_records_filter)
power_supply_unit_api_bp.route(psu_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
power_supply_unit_api_bp.route(psu_api_routes.GET_POWERS, methods=['GET'])(get_powers)
power_supply_unit_api_bp.route(psu_api_routes.GET_TYPES, methods=['GET'])(get_types)
power_supply_unit_api_bp.route(psu_api_routes.GET_FORM_FACTORS, methods=['GET'])(get_form_factors)
power_supply_unit_api_bp.route(psu_api_routes.CREATE, methods=['POST'])(create)
power_supply_unit_api_bp.route(psu_api_routes.UPDATE, methods=['PUT'])(update)
power_supply_unit_api_bp.route(psu_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
