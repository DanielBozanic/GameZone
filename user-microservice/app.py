
from flask import Flask
from flask_migrate import Migrate
from models.user import db
from routes.user_bp import user_bp

app = Flask(__name__)
app.config.from_object('config')
db.init_app(app)
migrate = Migrate(app, db)
app.register_blueprint(user_bp, url_prefix='/users')


if __name__ == '__main__':
    app.debug = True
    app.run()
