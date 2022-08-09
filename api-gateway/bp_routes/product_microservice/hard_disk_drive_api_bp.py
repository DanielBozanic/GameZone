from flask import Blueprint
from api.product_microservice.hard_disk_drive_api import *
from utils.routes.product_microservice import hard_disk_drive_api_routes

hard_disk_drive_api_bp = Blueprint('hard_disk_drive_api_bp', __name__)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_ALL, methods=['GET'])(get_all)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_NUMBER_OF_RECORDS,
                             methods=['GET'])(get_number_of_records)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_NUMBER_OF_RECORDS_SEARCH,
                             methods=['GET'])(get_number_of_records_search)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.FILTER, methods=['POST'])(filter)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_NUMBER_OF_RECORDS_FILTER,
                             methods=['POST'])(get_number_of_records_filter)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_MANUFACTURERS, methods=['GET'])(get_manufacturers)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_CAPACITIES, methods=['GET'])(get_capacities)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_FORMS, methods=['GET'])(get_forms)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.GET_DISK_SPEEDS, methods=['GET'])(get_disk_speeds)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.CREATE, methods=['POST'])(create)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.UPDATE, methods=['PUT'])(update)
hard_disk_drive_api_bp.route(hard_disk_drive_api_routes.DELETE + "/<id>", methods=['DELETE'])(delete)
