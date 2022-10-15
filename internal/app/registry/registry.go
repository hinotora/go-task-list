package registry

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/hinotora/go-task-list/internal/app/list"
	"github.com/hinotora/go-task-list/internal/app/task"
	"github.com/hinotora/go-task-list/internal/app/utils"
)

type Registry struct {
	Lists map[string]*list.List
}

func MakeRegistry() *Registry {
	return &Registry{
		Lists: make(map[string]*list.List),
	}
}

func (r *Registry) StoreRegistry() {
	utils.StoreData(utils.Data_registry_file, *r)
}

func (r *Registry) LoadRegistry() {
	dataFile, _ := os.Open(fmt.Sprintf("%s/%s", utils.Prepare_dir(), utils.Data_registry_file))

	dataDecoder := gob.NewDecoder(dataFile)
	dataDecoder.Decode(r)

	dataFile.Close()
}

func (r *Registry) AddList(list *list.List) {
	r.Lists[list.Id] = list

	fmt.Printf("Created new list [%s] with identifier [%s]. \n", list.Name, list.Id)
}

func (r *Registry) RemoveList(id string) {
	delete(r.Lists, id)

	fmt.Printf("Removed list [%s] from memory. \n", id)
}

func (r *Registry) GetAllLists() {
	if len(r.Lists) == 0 {

		fmt.Println("There are no task lists in memory. Please add new tasklist first")

		return
	}

	for _, list := range r.Lists {

		time_created := utils.GetClearDateFromUNIX(list.Date_created)

		output := fmt.Sprintf("\nList id: %s \nList name: %s \nList date created: %s", list.Id, list.Name, time_created)

		fmt.Println(output)
		fmt.Println(utils.Console_delimeter)
	}
}

func (r *Registry) GetList(id string) {
	list, exists := r.Lists[id]

	if exists {
		list.PrintWithTasks()
	} else {
		fmt.Println("List with given identidier don't exist")
	}
}

func (r *Registry) AddNewTask(id string, task *task.Task) {
	r.Lists[id].AddNewTask(task)

	fmt.Printf("Added new task with identifier [%s] to list [%s] \n", task.Id, r.Lists[id].Name)
}

func (r *Registry) RemoveTaskFromList(listId string, taskId string) {

	list, l_exists := r.Lists[listId]

	if !l_exists {
		fmt.Println("List with given id not found")
		return
	}

	_, t_exists := list.Tasks[taskId]

	if !t_exists {
		fmt.Printf("Task with given id not found in list [%s]", list.Id)
		return
	}

	list.RemoveTaskFromList(taskId)

	fmt.Printf("Removed task from list [%s] \n", list.Name)
}

func (r *Registry) SetTaskStatus(listId string, taskId string, status bool) {

	list, l_exists := r.Lists[listId]

	if !l_exists {
		fmt.Println("List with given id not found")
		return
	}

	task, t_exists := list.Tasks[taskId]

	if !t_exists {
		fmt.Printf("Task with given id not found in list [%s]", list.Id)
		return
	}

	task.SetStatus(true)

	fmt.Printf("Task [%s] marked as done \n", task.Name)
}
