package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Stringer interface {
	String() string
}

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// view tasks list in stdout
func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X"
		}
		formatted += fmt.Sprintf("%s %d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

// add new tasks
func (l *List) Add(taskName string) {
	t := item{
		Task:        taskName,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Mark about tasks todo
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Items %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

// delete tasks with key i
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Items %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// save list tasks in file
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

// get list task from file
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
