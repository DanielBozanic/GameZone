from flask import Flask
from flask_mail import Mail, Message
from flask import request, jsonify
from jinja2 import Environment, FileSystemLoader, select_autoescape
import os

app = Flask(__name__)

app.config['MAIL_SERVER'] = 'smtp-mail.outlook.com'
app.config['MAIL_PORT'] = 587
app.config['MAIL_USERNAME'] = 'gamezoneofficial99@hotmail.com'
app.config['MAIL_PASSWORD'] = 'PassWord_FOR_NTP!'
app.config['MAIL_USE_TLS'] = True
app.config['MAIL_USE_AUTH'] = True

mail = Mail(app)


@app.route("/api/email/sendEmail", methods=['POST'])
def send_email():
    msg = Message(subject=request.json["subject"],
                  sender=app.config['MAIL_USERNAME'],
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
