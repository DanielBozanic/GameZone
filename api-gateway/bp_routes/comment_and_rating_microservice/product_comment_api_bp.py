from flask import Blueprint
from api.comment_and_rating_microservice.product_comment_api import *
from utils.routes.comment_and_rating_microservice import product_comment_api_routes

product_comment_api_bp = Blueprint('product_comment_api_bp', __name__)
product_comment_api_bp.route(product_comment_api_routes.GET_ALL, methods=['GET'])(get_all)
product_comment_api_bp.route(product_comment_api_routes.GET_BY_ID + "/<id>", methods=['GET'])(get_by_id)
product_comment_api_bp.route(product_comment_api_routes.GET_BY_PRODUCT_ID + "/<productId>",
                             methods=['GET'])(get_by_product_id)
product_comment_api_bp.route(product_comment_api_routes.ADD_COMMENT, methods=['POST'])(add_comment)
product_comment_api_bp.route(product_comment_api_routes.EDIT_COMMENT, methods=['PUT'])(edit_comment)
product_comment_api_bp.route(product_comment_api_routes.DELETE_COMMENT + "/<id>",
                             methods=['DELETE'])(delete_comment)
product_comment_api_bp.route(product_comment_api_routes.GET_BY_USER_ID + "/<userId>",
                             methods=['GET'])(get_by_user_id)
