build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -r '\.go$' make run