package list

import (
	"fmt"
	"time"

	"github.com/hinotora/go-task-list/internal/app/task"
	"github.com/hinotora/go-task-list/internal/app/utils"

	"github.com/google/uuid"
)

type List struct {
	Id           string                `json:"id"`
	Name         string                `json:"name"`
	Date_created int64                 `json:"date_created"`
	Tasks        map[string]*task.Task `json:"tasks"`
}

func MakeList(name string) *List {
	now := time.Now().Unix()
	uuid := uuid.New().String()

	return &List{
		Id:           uuid,
		Name:         name,
		Date_created: now,
		Tasks:        make(map[string]*task.Task),
	}
}

func (list *List) AddNewTask(task *task.Task) {
	list.Tasks[task.GetId()] = task
}

func (list *List) RemoveTaskFromList(id string) {
	delete(list.Tasks, id)
}

func (list *List) PrintWithTasks() {
	header := fmt.Sprintf("Name: %s \nDate created: %s", list.Name, utils.GetClearDateFromUNIX(list.Date_created))

	fmt.Println(header)
	fmt.Println(utils.Console_delimeter)

	if len(list.Tasks) == 0 {
		fmt.Println("There are no tasks in this list")
		return
	}

	for _, task := range list.Tasks {

		isDone := "-"

		if task.Done {
			isDone = "+"
		}

		fmt.Printf("[%s] [%s]: %s \n", task.Id, isDone, task.Name)
	}
}
