build:
	protoc -I . scribble.proto --go_out=plugins=grpc:.

build-python:
	python3 -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. scribble.proto