from __future__ import print_function

import string
import threading
import time
import random

import grpc
import sys
from src import server_pb2 as server__pb2
#
from src.server_pb2_grpc import StudentServiceStub

channel = grpc.insecure_channel('localhost:5000')
stub = StudentServiceStub(channel)

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

    for grade in stub.SendSolution(solution, wait_for_ready=True):
        print("grade received:", grade, " for task ", task)
        break

    print("I closed channel after grade received for", solution)


def subscribe():
    print("listening for tasks")
    for task in stub.SignForLab(sub, wait_for_ready=True):
        print("received task{\n", task, "}")
        threading.Thread(target=await_grading, args=[task]).start()
        print("awaiting further tasks")
    print("WHAT HTE FUCK")


threading.Thread(target=subscribe).start()

create_task = server__pb2.Task(type=0, deadline=int(time.time()) + 100, descriptions=["jedna", "druga", "trzecia"])

time.sleep(2)
stub.TESTCreateAssignment(create_task)
