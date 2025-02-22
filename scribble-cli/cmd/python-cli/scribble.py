# Copyright © 2019 Tanishka Bhardwaj <tanishkab99@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
import grpc
import sys
import os
import click

# Get the current directory
current_dir = os.path.dirname(os.path.realpath(__file__))

# sys.path.append(PATH) will look for python files in "PATH" you import in your project 
# Use the current_dir to find the file created by protoc for python placed in the api folder with sys.path.append() 
sys.path.append(current_dir + "/../../../api/python/")
import scribble_pb2 
import scribble_pb2_grpc as scribble

@click.group(invoke_without_command=True)
@click.pass_context
def cli(ctx):
    if ctx.invoked_subcommand is None:
        click.echo("scribble is a cli tool that generates captioned images 🖼")

@cli.command()
@click.option('-c', '--caption', default='My Scribble Diary')
@click.option('-t', '--textsize', default=70)
@click.option('-i', '--imgsize', default=720)
@click.option('-s', '--server', default="scribble.kumarutsavanand.com")
def create(caption, textsize, imgsize, server):
    # create a channel and stub to server's address and port
    address = server
    # a channel as the name suggests is a channel for the requests and responses 
    # as RPCs in this case to traverse through
    channel = grpc.insecure_channel(address)
    # stubs are programs that live on the client side that forwards the requests to the server as RPCs
    # such that the methods and function call feel like to be called locally.
    stub = scribble.TextToImageStub(channel)

    # convert is a method on stub that is defined as the RPC in pur scribble.proto file that converts
    # text into an image. The body of the function is written in Go and lives in a server
    #
    # The response object holds the Image data in the form of raw bytes as defined in 
    # the protocol buffer file scribble.proto.
    # The rest of the code is self explainatory
    response = stub.convert(scribble_pb2.ImageSpec(text=caption, fontsize=textsize, imgsize=imgsize))
    data = response.Image
    filename = "image.png"

    # As the open function expects strings as its argument when opened with 'w' mode.
    # We need to open our file for writing in write as bytes mode for our file.
    # As the Image is defined as a response from the RPC as a stream of bytes in the scribble.proto definition. 
    imagefile = open(filename, 'wb')
    imagefile.write(data)
    # Never forget to close the file after you're done reading or writing
    imagefile.close()

    # TODO: implement the cli to make this program work as a CLI tool just like it go counterpart

    # try experimenting with the parameters to the convert function.


if __name__ == '__main__':
    # calls the main function
    cli()