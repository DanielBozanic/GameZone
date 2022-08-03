from flask import Flask
from flask_mail import Mail, Message
from flask import request, jsonify
from jinja2 import Environment, FileSystemLoader, select_autoescape
from dotenv import load_dotenv
import os

load_dotenv(os.path.abspath(os.path.dirname(__file__)) + "/env_vars.env")

app = Flask(__name__)

app.config['SECRET_KEY'] = 'secret'
app.config['MAIL_SERVER'] = 'smtp.sendgrid.net'
app.config['MAIL_PORT'] = 587
app.config['MAIL_USE_TLS'] = True
app.config['MAIL_USERNAME'] = 'apikey'
app.config['MAIL_PASSWORD'] = os.getenv('MAIL_PASSWORD')
app.config['MAIL_DEFAULT_SENDER'] = os.getenv('MAIL_DEFAULT_SENDER')

mail = Mail(app)


@app.route("/api/email/sendEmail", methods=['POST'])
def send_email():
    msg = Message(subject=request.json["subject"],
                  sender=app.config['MAIL_DEFAULT_SENDER'],
                  recipients=request.json["recipients"])
    env = Environment(
        loader=FileSystemLoader(os.path.abspath(os.path.dirname(__file__)) + "/templates"),
        autoescape=select_autoescape()
    )
    template = env.get_template(str(request.json["content"]["template"]) + ".html")
    msg.html = template.render(request.json["content"]["params"])
    if "attachment" in request.json.keys() is not None:
        with app.open_resource(os.path.abspath(os.path.dirname(__file__) + request.json["attachment"]["path"])) as fp:
            msg.attach(request.json["attachment"]["name"], request.json["attachment"]["fileType"], fp.read())
    mail.send(msg)
    return jsonify(status_code=200, content={"message": "Email has been sent"})


if __name__ == '__main__':
    app.run(port=5001, debug=True)
