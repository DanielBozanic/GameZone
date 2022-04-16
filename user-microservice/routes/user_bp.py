from flask import Blueprint
from controllers.user_controller import *

user_bp = Blueprint('user_bp', __name__)
user_bp.route('/', methods=['GET'])(get_all)
user_bp.route('/<int:user_id>', methods=['GET'])(get_by_id)
user_bp.route('/', methods=['POST'])(create)
user_bp.route('/<int:user_id>', methods=['PUT'])(update)
user_bp.route('/<int:user_id>', methods=['DELETE'])(delete)