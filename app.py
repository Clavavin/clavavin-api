from flask import Flask
from flask_restful import Api
from resources.api import HelloWorld

app = Flask(__name__)
api = Api(app)

api.add_resource(HelloWorld, '/')

if __name__ == "__main__":
    app.run(debug=True)
