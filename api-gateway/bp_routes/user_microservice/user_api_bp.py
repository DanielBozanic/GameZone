from flask import Blueprint
from api.user_microservice.user_api import *
from utils.routes.user_microservice import user_api_routes

user_api_bp = Blueprint('user_api_bp', __name__)
user_api_bp.route(user_api_routes.REGISTER, methods=['POST'])(register)
user_api_bp.route(user_api_routes.GET_VERIFICATION_CODE, methods=['GET'])(get_verification_code)
user_api_bp.route(user_api_routes.VERIFY_ACCOUNT, methods=['PUT'])(verify_account)
user_api_bp.route(user_api_routes.ADD_EMPLOYEE_AND_ADMIN, methods=['POST'])(add_employee_and_admin)
user_api_bp.route(user_api_routes.GET_USER_BY_ID, methods=['GET'])(get_by_id)
user_api_bp.route(user_api_routes.GET_ALL_REGISTERED_USERS, methods=['GET'])(get_all_registered_users)
user_api_bp.route(user_api_routes.GET_NUMBER_OF_RECORDS_REGISTERED_USERS,
                  methods=['GET'])(get_number_of_records_registered_users)
user_api_bp.route(user_api_routes.SEARCH_REGISTERED_USERS, methods=['GET'])(search_registered_users)
user_api_bp.route(user_api_routes.GET_NUMBER_OF_RECORDS_REGISTERED_USERS_SEARCH,
                  methods=['GET'])(get_number_of_records_registered_users_search)
user_api_bp.route(user_api_routes.UPDATE, methods=['PUT'])(update)
user_api_bp.route(user_api_routes.CHANGE_PASSWORD, methods=['PUT'])(change_password)
