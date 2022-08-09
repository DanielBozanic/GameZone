from flask import Blueprint
from api.product_microservice.product_api import *
from utils.routes.product_microservice import product_api_routes

product_api_bp = Blueprint('product_api_bp', __name__)
product_api_bp.route(product_api_routes.GET_PRODUCT_BY_ID + "/<id>", methods=['GET'])(get_product_by_id)
product_api_bp.route(product_api_routes.SEARCH_BY_NAME, methods=['GET'])(search_by_name)
product_api_bp.route(product_api_routes.GET_NUMBER_OF_RECORDS_SEARCH, methods=['GET'])(get_number_of_records_search)
product_api_bp.route(product_api_routes.GET_MAIN_PAGE_PRODUCTS, methods=['GET'])(get_main_page_products)
product_api_bp.route(product_api_routes.GET_POPULAR_PRODUCTS, methods=['GET'])(get_popular_products)
product_api_bp.route(product_api_routes.ADD_PRODUCT_TO_MAIN_PAGE + "/<productId>",
                     methods=['PUT'])(add_product_to_main_page)
product_api_bp.route(product_api_routes.REMOVE_PRODUCT_FROM_MAIN_PAGE + "/<productId>",
                     methods=['PUT'])(remove_product_from_main_page)
product_api_bp.route(product_api_routes.DELETE_PRODUCT + "/<id>", methods=['DELETE'])(delete_product)
product_api_bp.route(product_api_routes.GET_RECOMMENDED_PRODUCTS, methods=['GET'])(get_recommended_products)
