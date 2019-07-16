build:
	GOOS=linux go build -o scribble "github.com/utsavanand2/scribble-diary/scribble-server"
	docker build -t gcr.io/utsav-talks/scribble-diary:0.1 .
	rm -f scribble

run:
	docker run --rm -p 8080:8080 utsavanand2/scribble-diary:0.1