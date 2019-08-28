# Copyright © 2019 Tanishka Bhardwaj <tanishkab99@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
import grpc
import sys
import os

# Get the current directory
current_dir = os.path.dirname(os.path.realpath(__file__))

# sys.path.append(PATH) will look for python files in "PATH" you import in your project 
# Use the current_dir to find the file created by protoc for python placed in the api folder with sys.path.append() 
sys.path.append(current_dir + "/../../../api/python/")
import scribble_pb2 
import scribble_pb2_grpc as scribble

def main():

    # create a channel and stub to server's address and port

    # a channel as the name suggests is a channel for the requests and responses 
    # as RPCs in this case to traverse through

    # stubs are programs that live on the client side that forwards the requests to the server as RPCs
    # such that the methods and function call feel like to be called locally.

    # convert is a method on stub that is defined as the RPC in pur scribble.proto file that converts
    # text into an image. The body of the function is written in Go and lives in a server
    #
    # The response object holds the Image data in the form of raw bytes as defined in 
    # the protocol buffer file scribble.proto.
    # The rest of the code is self explainatory

    # TODO: implement the cli to make this program work as a CLI tool just like it go counterpart

    # try experimenting with the parameters to the convert function.


    # As the open function expects strings as its argument when opened with 'w' mode.
    # We need to open our file for writing in write as bytes mode for our file.
    # As the Image is defined as a response from the RPC as a stream of bytes in the scribble.proto definition. 

    # Never forget to close the file after you're done reading or writing

if __name__ == '__main__':
    # calls the main function
    main()