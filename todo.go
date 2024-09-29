package todo

import (
	"errors"
	"time"
)

type item struct {
	Task      string
	Done      bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {

	todo := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)

}

func (t *Todos) Completed (index int) error {
	
	if index >= len(*t) || index < 0 {
		return errors.New("Invalid index")
	}
	
	(*t)[index].CompletedAt = time.Now()
	(*t)[index].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {

	if index >= len(*t) || index < 0 {
		return errors.New("Invalid index")
	}

	*t = append((*t)[:index], (*t)[index+1:]...)

	return nil
}