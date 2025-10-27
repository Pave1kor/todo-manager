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

// view information about data of task
func (l *List) DataTask(numberTask int) error {
	if numberTask > len(*l) {
		return fmt.Errorf("Task with %q number don`t exist", numberTask)
	}
	for number, task := range *l {
		if numberTask == number {
			fmt.Printf("Tasks: %q\n", task.Task)
			y, m, d := task.CreatedAt.Date()
			fmt.Printf("Data created: %d:%d:%d\n", d, m, y)
			if task.Done {
				y, m, d := task.CompletedAt.Date()
				fmt.Printf("Data completed: %d:%d:%d\n", d, m, y)
			} else {
				fmt.Printf("Item don`t completed!")
			}
		}
	}
	return nil
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

func (l *List) CompleteTasks() error {
	formatted := ""
	count := 1
	for _, tasks := range *l {
		if tasks.Done {
			formatted += fmt.Sprintf("%d: %s\n", count, tasks.Task)
			count++
		}
	}
	if len(formatted) == 0 {
		return fmt.Errorf("Don`t exists completed items")
	}
	fmt.Print(formatted)
	return nil
}
