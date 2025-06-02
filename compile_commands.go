package gbld

import "path/filepath"

type CompileCommand struct {
	Arguments []string `json:"arguments"`
	Directory string   `json:"directory"`
	File      string   `json:"file"`
}

func (pj *Project) GetCompileCommands() []CompileCommand {
	var commands []CompileCommand

	for _, mod := range pj.modules {
		for _, cmd := range mod.command_log {
			if cmd.file.IsReal() {
				commands = append(commands, CompileCommand{
					cmd.GetArgList(),
					filepath.Dir(cmd.file.Path()),
					cmd.file.Path(),
				})
			}
		}
	}

	return commands
}
