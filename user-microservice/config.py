import os

SECRET_KEY = os.urandom(32)

basedir = os.path.abspath(os.path.dirname(__file__))

DEBUG = True

SQLALCHEMY_DATABASE_URI = 'mysql+pymysql://root:root@localhost:3306/users-db?charset=utf8mb4'

SQLALCHEMY_TRACK_MODIFICATIONS = False
