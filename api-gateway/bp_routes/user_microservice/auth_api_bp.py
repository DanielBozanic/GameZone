from flask import Blueprint
from api.user_microservice.auth_api import *
from utils.routes.user_microservice import auth_api_routes

auth_api_bp = Blueprint('auth_api_bp', __name__)
auth_api_bp.route(auth_api_routes.LOGIN, methods=['POST'])(login)
