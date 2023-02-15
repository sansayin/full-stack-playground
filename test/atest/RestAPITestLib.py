#!/usr/bin/python3
#import json
from pprint import pprint
import requests
from robot.api import logger
#from robot.libraries.BuiltIn import BuiltIn
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
class RestAPITestLib():

    def rest_post(self, url, data):
        #logger.console(data)
        response = requests.post(url, json=data,headers=headers)
        #logger.console(response.content)
        return response.json()

    def rest_get(self,url):
        #logger.console(url)
        response = requests.get(url)
        #logger.console(response.content)
        return response.json() 

    def rest_put(self, url, data):
        #todo = {"userId": 1, "title": "Wash car", "completed": True}
        response = requests.put(url, data)
        return response.json()

    def rest_delete(self, url):
        response = requests.delete(url)
        return response.json()

if __name__ == "__main__":
    API =RestAPITestLib()
    resp = API.rest_get("http://172.22.255.1:8081/users/1")
    pprint(resp)
