package todo

import (
    "context"
    todo "microservice/proto/todo"
)

type TodoServiceServer struct {
    todo.UnimplementedTodoServiceServer
    todos []*todo.Todo
}

func NewTodoServiceServer() *TodoServiceServer {
    return &TodoServiceServer{}
}

func (s *TodoServiceServer) ListTodos(ctx context.Context, _ *todo.Empty) (*todo.TodoList, error) {
    return &todo.TodoList{Todos: s.todos}, nil
}

func (s *TodoServiceServer) CreateTodo(ctx context.Context, req *todo.TodoRequest) (*todo.TodoResponse, error) {
    newTodo := &todo.Todo{
        Title:       req.Title,
        Description: req.Description,
    }
    s.todos = append(s.todos, newTodo)
    return &todo.TodoResponse{Message: "Todo created"}, nil
}
