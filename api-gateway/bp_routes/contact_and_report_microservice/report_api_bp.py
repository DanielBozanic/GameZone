from flask import Blueprint
from api.contact_and_report_microservice.report_api import *
from utils.routes.contact_and_report_microservice import report_api_routes

report_api_bp = Blueprint('report_api_bp', __name__)
report_api_bp.route(report_api_routes.ADD_REPORT, methods=['POST'])(add_report)
report_api_bp.route(report_api_routes.GET_REPORTS_BY_USER_ID + "/<userId>", methods=['GET'])(get_reports_by_user_id)
