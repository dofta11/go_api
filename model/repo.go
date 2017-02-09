package model

import (
	"fmt"
	"go_api/vo"
)

var currentId int

var Todos_val vo.Todos

// Give us some seed data
func init() {
	RepoCreateTodo(vo.Todo{Name: "Write presentation"})
	RepoCreateTodo(vo.Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) vo.Todo {
	for _, t := range Todos_val {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return vo.Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t vo.Todo) vo.Todo {
	currentId += 1
	t.Id = currentId
	Todos_val = append(Todos_val, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range Todos_val {
		if t.Id == id {
			Todos_val = append(Todos_val[:i], Todos_val[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
