package gbld

import (
	"fmt"
	"os/exec"
	"strconv"
	"sync"
)

type Command struct {
	mod *Module

	name  string
	file  File
	args  []string
	flags map[string]interface{}
}

func parse_flag_value(value interface{}) (string, bool) {
	switch value := value.(type) {
	case string:
		return value, true
	case bool:
		if value {
			return "", true
		} else {
			return "", false
		}
	case int:
		return strconv.Itoa(value), true
	default:
		return fmt.Sprint(value), true
	}
}

func (pj *Project) SetCommandFlag(name string, value interface{}) {
	pj.command_flags[name] = value
}

func (mod *Module) SetCommandFlag(name string, value interface{}) {
	mod.command_flags[name] = value
}

func (mod *Module) NewCommand() *Command {
	cmd := new(Command)
	cmd.mod = mod

	cmd.flags = make(map[string]interface{})
	cmd.args = make([]string, 0)

	mod.command_log = append(mod.command_log, cmd) // for logging/generating compile_commands.json

	return cmd
}

func (cmd *Command) SetName(name string) {
	cmd.name = name
}

func (cmd *Command) SetFile(file File) {
	cmd.file = file
}

func (cmd *Command) SetArg(arg string) {
	cmd.args = append(cmd.args, arg)
}

func (cmd *Command) SetArgs(args ...string) {
	cmd.args = append(cmd.args, args...)
}

func (cmd *Command) SetFlag(name string, value interface{}) {
	cmd.flags[name] = value
}

func (cmd *Command) GetFile() File {
	return cmd.file
}

func (cmd *Command) GetArgList() []string {
	var list []string
	list = append(list, cmd.name)

	// collect all the flags
	flags := make(map[string]interface{})
	for name, value := range cmd.mod.pj.command_flags {
		flags[name] = value
	}
	for name, value := range cmd.mod.command_flags {
		flags[name] = value
	}
	for name, value := range cmd.flags {
		flags[name] = value
	}

	for name, value := range flags {
		v, ok := parse_flag_value(value)
		if ok {
			list = append(list, name, v)
		}
	}

	for _, arg := range cmd.args {
		list = append(list, arg)
	}

	return list
}

func (cmd *Command) Exec(wd string, error_cb func(output []byte)) {
	args := cmd.GetArgList()
	c := exec.Command(args[0], args[1:]...)
	c.Dir = wd

	output, e := c.CombinedOutput()

	if e != nil {
		error_cb(output)
	}
}

func (cmd *Command) ExecAsync(wg *sync.WaitGroup, wd string, error_cb func(output []byte)) {
	wg.Add(1)
	go func() {
		cmd.Exec(wd, error_cb)
		wg.Done()
	}()
}
