import os

SECRET_KEY = "7kRh7cgjun9S83Hu06HhqhB8sGYkZKHrZ7CRkpQJHfOzXTllQPcIWCSn3IcUccq"

basedir = os.path.abspath(os.path.dirname(__file__))

DEBUG = True

SQLALCHEMY_DATABASE_URI = 'mysql+pymysql://root:root@localhost:3306/users-db?charset=utf8mb4'

SQLALCHEMY_TRACK_MODIFICATIONS = False
