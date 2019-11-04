from locust import HttpLocust, TaskSet, task
import random
import string


class GreetingTaskSet(TaskSet):
    def on_start(self):
        self.client.headers = {
            'Content-Type': 'application/json; charset=utf8'}

    @task
    def fetch_ping(self):
        self.client.get('/ping')


class GreetingLocust(HttpLocust):
    task_set = GreetingTaskSet

    min_wait = 100
    max_wait = 1000
