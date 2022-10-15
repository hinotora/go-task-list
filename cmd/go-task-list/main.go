package main

import (
	"fmt"

	"github.com/hinotora/go-task-list/internal/app/command"
	"github.com/hinotora/go-task-list/internal/app/list"
	"github.com/hinotora/go-task-list/internal/app/registry"
	"github.com/hinotora/go-task-list/internal/app/task"
)

func main() {

	command := command.ParseCLI()

	registry := registry.MakeRegistry()
	registry.LoadRegistry()

	if command.Action == "all" {

		registry.GetAllLists()

	} else if command.Action == "list" && command.Lid != "" {

		registry.GetList(command.Lid)

	} else if command.Action == "createlist" && command.Lname != "" {

		newList := list.MakeList(command.Lname)

		registry.AddList(newList)

	} else if command.Action == "removelist" && command.Lid != "" {

		registry.RemoveList(command.Lid)

	} else if command.Action == "createtask" && command.Lid != "" && command.Tname != "" {

		newTask := task.NewTask(command.Tname)

		registry.AddNewTask(command.Lid, newTask)

	} else if command.Action == "removetask" && command.Lid != "" && command.Tid != "" {

		registry.RemoveTaskFromList(command.Lid, command.Tid)

	} else if command.Action == "done" && command.Lid != "" && command.Tid != "" {

		registry.SetTaskStatus(command.Lid, command.Tid, true)

	} else {
		fmt.Println("Unknown action, please check your syntax")
		return
	}

	registry.StoreRegistry()
}
