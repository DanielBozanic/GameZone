from flask import Blueprint
from api.contact_and_report_microservice.ban_api import *
from utils.routes.contact_and_report_microservice import ban_api_routes

ban_api_bp = Blueprint('ban_api_bp', __name__)
ban_api_bp.route(ban_api_routes.GET_USER_BAN_HISTORY + "/<userId>", methods=['GET'])(get_user_ban_history)
ban_api_bp.route(ban_api_routes.ADD_BAN, methods=['POST'])(add_ban)
ban_api_bp.route(ban_api_routes.SEND_EMAIL_TO_BANNED_USER, methods=['POST'])(send_email_to_banned_user)
