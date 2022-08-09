from flask import Blueprint
from api.contact_and_report_microservice.contact_api import *
from utils.routes.contact_and_report_microservice import contact_api_routes

contact_api_bp = Blueprint('contact_api_bp', __name__)
contact_api_bp.route(contact_api_routes.GET_UNANSWERED_CONTACT_MESSAGES_BY_USER_ID + "/<userId>",
                     methods=['GET'])(get_unanswered_contact_messages_by_user_id)
contact_api_bp.route(contact_api_routes.GET_CONTACT_MESSAGES_BY_USER_ID + "/<userId>",
                     methods=['GET'])(get_contact_messages_by_user_id)
contact_api_bp.route(contact_api_routes.GET_UNANSWERED_CONTACT_MESSAGES,
                     methods=['GET'])(get_unanswered_contact_messages)
contact_api_bp.route(contact_api_routes.SEND_CONTACT_MESSAGE,
                     methods=['POST'])(send_contact_message)
contact_api_bp.route(contact_api_routes.ANSWER_CONTACT_MESSAGE,
                     methods=['PUT'])(answer_contact_message)
