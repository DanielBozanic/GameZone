from flask import Blueprint
from controllers.user_controller import *

user_bp = Blueprint('user_bp', __name__)
user_bp.route('/register', methods=['POST'])(register)
user_bp.route('/addEmployeeAndAdmin', methods=['POST'])(add_employee_and_admin)
