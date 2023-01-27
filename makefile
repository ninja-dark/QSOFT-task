build:
	go build -o qsoft-task cmd/qsoft_task/main.go

run: build
		./qsoft-task