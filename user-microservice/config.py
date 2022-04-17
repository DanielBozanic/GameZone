import os

SECRET_KEY = b'\xd5\xce\x99\x88\xf9\xe9\x94\xff\x12\x86\xd3Z\xc0\xacH\x04\x92&G\xbf?y\x7f\xd9\xc9\xfc\x88p\xf1\xf2/E'

basedir = os.path.abspath(os.path.dirname(__file__))

DEBUG = True

SQLALCHEMY_DATABASE_URI = 'mysql+pymysql://root:root@localhost:3306/users-db?charset=utf8mb4'

SQLALCHEMY_TRACK_MODIFICATIONS = False
