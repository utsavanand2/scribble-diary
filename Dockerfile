FROM alpine:3.10

LABEL maintainer=utsavanand2@gmail.com

RUN mkdir /scribble-diary
WORKDIR /scribble-diary

ENV IMAGEMAGICK_VERSION=7.0.8.58-r0

# Install ImageMagick
RUN apk update && \
	apk add imagemagick=${IMAGEMAGICK_VERSION}

# ADD ./fonts .
ADD ./fonts ./fonts

# ADD the backend binary
ADD scribble .

# Expose port for the scribble server to listen on
EXPOSE 80

# Run the scribble server
ENTRYPOINT [ "./scribble" ]
