package main

import "fmt"

type todo struct {
	id     int
	title  string
	status string
}

type allTodos struct {
	todos []todo
}

var box = allTodos{}

func main() {
	todo1 := todo{id: 1, title: "Clean", status: "in progress"}
	todo2 := todo{id: 2, title: "Work out", status: "Done"}
	todo3 := todo{id: 3, title: "Work", status: "Done"}
	todo4 := todo{id: 4, title: "Buy car", status: "in progress"}
	todo5 := todo{id: 5, title: "Buy milk", status: "in progress"}

	box.createTodo(todo1)
	box.createTodo(todo2)
	box.createTodo(todo3)
	box.createTodo(todo4)
	box.createTodo(todo5)
	box.createTodo(todo{id: 6, title: "Read a book", status: "in progress"})

	box.listTodos()

	box.deleteTodo(1)

	box.updateTodo(0, "Clean my room", "Done")

	box.listTodos()

}

func (a *allTodos) createTodo(t todo) {
	a.todos = append(a.todos, t)
}

func (a *allTodos) listTodos() {
	fmt.Println(box.todos)
}

func (a *allTodos) deleteTodo(i int) {
	a.todos = append(a.todos[:i], a.todos[i+1:]...)
}

func (a *allTodos) updateTodo(i int, title string, status string) {
	a.todos[i].title = title
	a.todos[i].status = status
}
