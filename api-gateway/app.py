from flask import Flask
from flask_cors import CORS
from bp_routes.business_report_microservice.business_report_api_bp import \
    business_report_api_bp, \
    business_report_api_routes
from bp_routes.comment_and_rating_microservice.product_comment_api_bp import \
    product_comment_api_bp, \
    product_comment_api_routes
from bp_routes.contact_and_report_microservice.report_api_bp import \
    report_api_bp, \
    report_api_routes
from bp_routes.contact_and_report_microservice.ban_api_bp import \
    ban_api_bp, \
    ban_api_routes
from bp_routes.contact_and_report_microservice.contact_api_bp import \
    contact_api_bp, \
    contact_api_routes
from bp_routes.news_microservice.news_article_api_bp import \
    news_article_api_bp, \
    news_article_api_routes
from bp_routes.news_microservice.news_comment_api_bp import \
    news_comment_api_bp, \
    news_comment_api_routes
from bp_routes.news_microservice.news_subscription_api_bp import \
    news_subscription_api_bp, \
    news_subscription_api_routes
from bp_routes.product_microservice.console_api_bp import \
    console_api_bp, \
    console_api_routes
from bp_routes.product_microservice.graphics_card_api_bp import \
    graphics_card_api_bp, \
    graphics_card_api_routes
from bp_routes.product_microservice.hard_disk_drive_api_bp import \
    hard_disk_drive_api_bp, \
    hard_disk_drive_api_routes
from bp_routes.product_microservice.headphones_api_bp import \
    headphones_api_bp, \
    headphones_api_routes
from bp_routes.product_microservice.keyboard_api_bp import \
    keyboard_api_bp, \
    keyboard_api_routes
from bp_routes.product_microservice.monitor_api_bp import \
    monitor_api_bp, \
    monitor_api_routes
from bp_routes.product_microservice.motherboard_api_bp import \
    motherboard_api_bp, \
    motherboard_api_routes
from bp_routes.product_microservice.mouse_api_bp import \
    mouse_api_bp, \
    mouse_api_routes
from bp_routes.product_microservice.power_supply_unit_api_bp import \
    power_supply_unit_api_bp, \
    psu_api_routes
from bp_routes.product_microservice.processor_api_bp import \
    processor_api_bp, \
    processor_api_routes
from bp_routes.product_microservice.product_api_bp import \
    product_api_bp, \
    product_api_routes
from bp_routes.product_microservice.product_purchase_api_bp import \
    product_purchase_api_bp, \
    product_purchase_api_routes
from bp_routes.product_microservice.ram_api_bp import \
    ram_api_bp, \
    ram_api_routes
from bp_routes.product_microservice.solid_state_drive_api_bp import \
    solid_state_drive_api_bp, \
    solid_state_drive_api_routes
from bp_routes.product_microservice.video_game_api_bp import \
    video_game_api_bp, \
    video_game_api_routes
from bp_routes.user_microservice.auth_api_bp import \
    auth_api_bp, \
    auth_api_routes
from bp_routes.user_microservice.user_api_bp import \
    user_api_bp, \
    user_api_routes

app = Flask(__name__)
CORS(app)

app.register_blueprint(business_report_api_bp, url_prefix=business_report_api_routes.API)
app.register_blueprint(product_comment_api_bp, url_prefix=product_comment_api_routes.API)

app.register_blueprint(report_api_bp, url_prefix=report_api_routes.API)
app.register_blueprint(ban_api_bp, url_prefix=ban_api_routes.API)
app.register_blueprint(contact_api_bp, url_prefix=contact_api_routes.API)

app.register_blueprint(news_article_api_bp, url_prefix=news_article_api_routes.API)
app.register_blueprint(news_comment_api_bp, url_prefix=news_comment_api_routes.API)
app.register_blueprint(news_subscription_api_bp, url_prefix=news_subscription_api_routes.API)

app.register_blueprint(console_api_bp, url_prefix=console_api_routes.API)
app.register_blueprint(graphics_card_api_bp, url_prefix=graphics_card_api_routes.API)
app.register_blueprint(hard_disk_drive_api_bp, url_prefix=hard_disk_drive_api_routes.API)
app.register_blueprint(headphones_api_bp, url_prefix=headphones_api_routes.API)
app.register_blueprint(keyboard_api_bp, url_prefix=keyboard_api_routes.API)
app.register_blueprint(monitor_api_bp, url_prefix=monitor_api_routes.API)
app.register_blueprint(motherboard_api_bp, url_prefix=motherboard_api_routes.API)
app.register_blueprint(mouse_api_bp, url_prefix=mouse_api_routes.API)
app.register_blueprint(power_supply_unit_api_bp, url_prefix=psu_api_routes.API)
app.register_blueprint(processor_api_bp, url_prefix=processor_api_routes.API)
app.register_blueprint(product_api_bp, url_prefix=product_api_routes.API)
app.register_blueprint(product_purchase_api_bp, url_prefix=product_purchase_api_routes.API)
app.register_blueprint(ram_api_bp, url_prefix=ram_api_routes.API)
app.register_blueprint(solid_state_drive_api_bp, url_prefix=solid_state_drive_api_routes.API)
app.register_blueprint(video_game_api_bp, url_prefix=video_game_api_routes.API)

app.register_blueprint(auth_api_bp, url_prefix=auth_api_routes.API)
app.register_blueprint(user_api_bp, url_prefix=user_api_routes.API)

if __name__ == '__main__':
    app.run(host="localhost", port=8000, debug=True)
