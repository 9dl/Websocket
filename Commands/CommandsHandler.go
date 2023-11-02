package Commands

type Command struct {
	Name        string
	Description string
	Run         func(args []string)
}

var CommandRegistry = make(map[string]Command)

func RegisterCommand(name, description string, runFunc func(args []string)) {
	CommandRegistry[name] = Command{Name: name, Description: description, Run: runFunc}
}
