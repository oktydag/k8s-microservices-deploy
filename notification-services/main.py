from flask import *
import requests
import redis

app = Flask(__name__)


REDIS_PORT = 6379
REDIS_HOST = 'redis' # set docker-compose service name

@app.route("/v1/api/notification/info",methods = ["GET"])
def index():
    r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT)
    # print(r.ping())
    # print(r.keys())
    # print(type(r.get('customer-count')))

    return 'Hi business administrator, current customer count is : ' + r.get('customer-count').decode("utf-8") 

if __name__ == "__main__":
    app.run(port=5006,debug=True)