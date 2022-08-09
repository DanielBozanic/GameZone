from flask import Blueprint
from api.product_microservice.product_purchase_api import *
from utils.routes.product_microservice import product_purchase_api_routes

product_purchase_api_bp = Blueprint('product_purchase_api_bp', __name__)
product_purchase_api_bp.route(product_purchase_api_routes.CHECK_IF_PRODUCT_IS_PAID_FOR,
                              methods=['GET'])(check_if_product_is_paid_for)
product_purchase_api_bp.route(product_purchase_api_routes.CONFIRM_PURCHASE, methods=['POST'])(confirm_purchase)
product_purchase_api_bp.route(product_purchase_api_routes.SEND_PURCHASE_CONFIRMATION_MAIL,
                              methods=['POST'])(send_purchase_confirmation_mail)
product_purchase_api_bp.route(product_purchase_api_routes.GET_PRODUCT_ALERT_BY_PRODUCT_ID_AND_USER_ID,
                              methods=['GET'])(get_product_alert_by_product_id_and_user_id)
product_purchase_api_bp.route(product_purchase_api_routes.ADD_PRODUCT_ALERT, methods=['POST'])(add_product_alert)
product_purchase_api_bp.route(product_purchase_api_routes.NOTIFY_PRODUCT_AVAILABILITY,
                              methods=['GET'])(notify_product_availability)
product_purchase_api_bp.route(product_purchase_api_routes.GET_PURCHASE_HISTORY, methods=['GET'])(get_purchase_history)
product_purchase_api_bp.route(product_purchase_api_routes.GET_NUMBER_OF_RECORDS_PURCHASE_HISTORY,
                              methods=['GET'])(get_number_of_records_purchase_history)
product_purchase_api_bp.route(product_purchase_api_routes.CONFIRM_PAYMENT, methods=['PUT'])(confirm_payment)
product_purchase_api_bp.route(product_purchase_api_routes.SEND_PURCHASED_DIGITAL_VIDEO_GAMES,
                              methods=['POST'])(send_purchased_digital_video_games)
