package gbld

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type CommandLog struct {
	args []string
	file *File
}

func (pj *Project) GenerateCompileCommands() {
	var entries []map[string]interface{}

	for _, cmd := range pj.cmd_log {
		entry := make(map[string]interface{})
		entry["directory"], _ = filepath.Abs(".")
		entry["arguments"] = cmd.args

		if cmd.file != nil { // only include entries that compile a specific file
			entry["file"] = cmd.file.path
			entries = append(entries, entry)
		}
	}

	data, _ := json.MarshalIndent(entries, "", "\t")
	os.WriteFile(filepath.Join(pj.out, "compile_commands.json"), data, pj.PERMISSION)
}
