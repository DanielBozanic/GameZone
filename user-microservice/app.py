from flask import Flask
from flask_migrate import Migrate
from models.user import db
from routes.user_bp import user_bp
from routes.auth_bp import auth_bp

app = Flask(__name__)
app.config.from_object('config')
db.init_app(app)
migrate = Migrate(app, db)
app.register_blueprint(user_bp, url_prefix='/api/users')
app.register_blueprint(auth_bp, url_prefix='/api/auth')


if __name__ == '__main__':
    app.debug = True
    app.run()
