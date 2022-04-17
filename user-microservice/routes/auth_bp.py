from flask import Blueprint
from controllers.auth_controller import *

auth_bp = Blueprint('auth_bp', __name__)
auth_bp.route('/login', methods=['POST'])(login)
