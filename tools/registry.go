package tools

import (
	"fmt"

	"github.com/jwhenry28/LLMUtils/model"
)

type Tool interface {
	Brief() string
	Usage() string
	Help() string
	Match() bool
	Invoke() string
}

var Registry = make(map[string]func(model.ToolInput) Tool)

func RegisterTool(name string, constructor func(model.ToolInput) Tool) error {
	if _, ok := Registry[name]; ok {
		return fmt.Errorf("tool already registered: %s", name)
	}

	Registry[name] = constructor
	return nil
}

func RunTool(input model.ToolInput) string {
	constructor, ok := Registry[input.GetName()]
	if !ok {
		return fmt.Sprintf("unregistered tool: %s. use 'help' tool to view registered tools", input.GetName())
	}

	tool := constructor(input)
	if !tool.Match() {
		return fmt.Sprintf("improper usage of tool: %s\n%s", input.GetName(), tool.Help())
	}

	return tool.Invoke()
}
