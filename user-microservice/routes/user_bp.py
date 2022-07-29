from flask import Blueprint
from controllers.user_controller import *

user_bp = Blueprint('user_bp', __name__)
user_bp.route('/register', methods=['POST'])(register)
user_bp.route('/getVerificationCode', methods=['GET'])(get_verification_code)
user_bp.route('/verifyAccount', methods=['PUT'])(verify_account)
user_bp.route('/addEmployeeAndAdmin', methods=['POST'])(add_employee_and_admin)
user_bp.route('/getById', methods=['GET'])(get_by_id)
user_bp.route('/getAllRegisteredUsers', methods=['GET'])(get_all_registered_users)
user_bp.route('/getNumberOfRecordsRegisteredUsers', methods=['GET'])(get_number_of_records_registered_users)
