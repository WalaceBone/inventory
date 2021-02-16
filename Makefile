NAME=server.go


all: clean build run

run:
	go run ./$(NAME)

clean:

fclean: clean all

generate:
	go run github.com/99designs/gqlgen generate

build:
	go build $(NAME)




