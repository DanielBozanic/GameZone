from flask import Blueprint
from api.product_microservice.solid_state_drive_api import *
from utils.routes.product_microservice import solid_state_drive_api_routes

solid_state_drive_api_bp = Blueprint('solid_state_drive_api_bp', __name__)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_ALL, methods=['GET'])(get_all)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_NUMBER_OF_RECORDS,
                               methods=['GET'])(get_number_of_records)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                               methods=['GET'])(get_number_of_records_search)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.FILTER, methods=['POST'])(filter)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                               methods=['POST'])(get_number_of_records_filter)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_CAPACITIES, methods=['GET'])(get_capacities)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_FORMS, methods=['GET'])(get_forms)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_MAX_SEQUENTIAL_READS,
                               methods=['GET'])(get_max_sequential_reads)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.GET_MAX_SEQUENTIAL_WRITES,
                               methods=['GET'])(get_max_sequential_writes)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.CREATE, methods=['POST'])(create)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.UPDATE, methods=['PUT'])(update)
solid_state_drive_api_bp.route(solid_state_drive_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
