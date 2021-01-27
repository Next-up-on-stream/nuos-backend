# API in golang

## Installation

[brew](https://docs.brew.sh/Installation)

**Note**: don't hesitate to reboot your computer. For me closing all terminals and opening a new one wasn't enough.

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
```

## Spec validation

Create your spec and validate it with:

  ```bash
swagger validate ./swagger.yml
  ```

## Top-down generation

Following this [tutorial](https://goswagger.io/tutorial/todo-list.html):

In `/examples/tutorials/todo-list/server-complete`, I used the spec `swagger.yml` and generated with 

```bash
swagger generate server -A TodoList -f ./swagger.yml
```

After disabling l295-308 in `/examples/tutorials/todo-list/server-complete/restapi/server.go`, and replacing `/restapi/configure_todo_list.go` with [this file from a link in the tutorial](https://github.com/go-swagger/go-swagger/blob/master/examples/tutorials/todo-list/server-complete/restapi/configure_todo_list.go), I can...

## Run api

```bash
cd /examples/tutorials/todo-list/server-complete/cmd/todo-list-server/
go install
go build
./todo-list-server --example1="something"
```

