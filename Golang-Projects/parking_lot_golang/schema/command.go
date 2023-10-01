package schema

import "parking_lot/errors"

// CMDStore interface holds the execute method
type CMDStore interface {
	Execute(cmd *Command) (string, error)
}

// Command holds the command properties
type Command struct {
	Command      string
	Arguments    []string
	Connection   CMDStore
	ShellHistory []*IShellHistory
}

// GetName returns the command name
func (cmd *Command) GetName() string {
	return cmd.Command
}

// GetArguments returns the command arguments
func (cmd *Command) GetArguments() []string {
	return cmd.Arguments
}

// IsExit checks if the command is `exit` or not
func (cmd *Command) IsExit() bool {
	return cmd.GetName() == string(CMDExit)
}

// Ok validate the command object by name and arguments list
func (cmd *Command) Ok() error {
	if _, ok := ValidCommandsByName[cmd.Command]; !ok {
		return errors.ErrInvalidCommand(cmd.Command)
	}
	if len(cmd.Arguments) != CMDArgumentLength[cmd.Command] {
		return errors.ErrInvalidArguments(cmd.Command, CMDArgumentLength[cmd.Command], len(cmd.Arguments))
	}

	return nil
}
