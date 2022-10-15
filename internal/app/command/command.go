package command

import "flag"

type Command struct {
	Action string

	Lname string
	Lid   string

	Tname string
	Tid   string
}

func ParseCLI() *Command {
	command := Command{}

	flag.StringVar(&command.Action, "action", "all", " ACTION ")

	flag.StringVar(&command.Lname, "lname", "", " LIST NAME ")

	flag.StringVar(&command.Lid, "lid", "", " LIST ID ")

	flag.StringVar(&command.Tname, "tname", "", " TASK NAME ")

	flag.StringVar(&command.Tid, "tid", "", " TASK ID ")

	flag.Parse()

	return &command
}
