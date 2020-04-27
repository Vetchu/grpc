from __future__ import print_function

import string
import threading
import time
import random

import grpc
import sys

from src.api.v1 import server_pb2 as server__pb2
#
from src.api.v1.server_pb2_grpc import StudentServiceStub

connection_array = []
channel = grpc.insecure_channel('localhost:5000')
connection_array.append(channel)
stub = StudentServiceStub(channel)
connection_array.append(stub)
args = sys.argv
print(args)


def randomString(stringLength=8):
    letters = string.ascii_lowercase
    return ''.join(random.sample(letters, stringLength))


args.append(randomString())
args.append("tekst")
args.append("tekst2")
sub = server__pb2.Subscribe(user=args[1], lab=[args[2], args[3]])

print(sub)


def await_grading(task):
    print("task solution sent, listening for grades for task{\n", task, "}")
    solution = server__pb2.Solution(id=str(time.time()), comment="done", solution=b'')
    try:
        for grade in connection_array[1].SendSolution(solution, wait_for_ready=True):
            print("grade received:", grade, " for task ", task)
            break
        print("I closed channel after grade received for", solution)
    except Exception as ex:
        print("channel closed while waiting for grade", ex)


def subscribe(array):
    print("listening for tasks")
    while True:
        try:
            for task in array[1].SignForLab(sub, wait_for_ready=True):
                print("received task{\n", task, "}")
                threading.Thread(target=await_grading, args=[task]).start()
                print("awaiting further tasks")
        except Exception as ex:
            print("listening for task failed with ", ex)
            array[0].close()
        try:
            print("retrying connection")
            array[0] = grpc.insecure_channel('localhost:5000')
            array[1] = StudentServiceStub(array[0])
        except Exception as ex:
            print("server is not ready yet", ex)
        print("connection will be reestablished when server wakes up")


threading.Thread(target=subscribe, args=(connection_array,)).start()

create_task = server__pb2.Task(type=0, deadline=int(time.time()) + 100, descriptions=["jedna", "druga", "trzecia"])

# connection_array[1].TESTCreateAssignment(create_task)
