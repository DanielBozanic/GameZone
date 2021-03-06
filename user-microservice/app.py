from flask import Flask
from flask_migrate import Migrate
from flask_cors import CORS
from db.database import db
from routes.user_bp import user_bp
from routes.auth_bp import auth_bp
from models import user
from models import user_verification

app = Flask(__name__)
app.config.from_object('config')
CORS(app)
db.init_app(app)
migrate = Migrate(app, db)
app.register_blueprint(user_bp, url_prefix='/api/users')
app.register_blueprint(auth_bp, url_prefix='/api/auth')


if __name__ == '__main__':
    app.debug = True
    app.run()
