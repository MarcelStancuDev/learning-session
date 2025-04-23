# Golang Chat App

## How to run

1.  Clone the repository using `git clone https://github.com/MarcelStancuDev/learning-session`
2.  Navigate into the cloned directory with `cd golang/chat-app`
3.  Generate go modules `go mod init`
4.  Download modules `go mod tidy`
5.  Run the application using `go run cmd/main.go`
6.  Run the docker build command `docker build -t chat-app .`
7.  Run the docker container using `docker run -p 8080:8080 chat-app`