package ishell

import (
	"fmt"
	"strings"

	"parking_lot/schema"
	"parking_lot/utils"
)

func availableCommandsHints() {
	fmt.Println(schema.AllCommendHint)
}

func greeting() {
	fmt.Println("Thanks for trying out")
}

// Process attempts to validate the command input string.
// Checks the command and arguments list are valid or not.
func Process(cmdInputStr string) (*schema.Command, error) {
	var cmd = new(schema.Command)
	cmdInputList, err := utils.SplitCmdArguments(cmdInputStr)
	if err != nil {
		return nil, err
	}
	cmd.Command = cmdInputList[0]
	if len(cmdInputList) > 1 {
		cmd.Arguments = strings.Split(cmdInputList[1], utils.Space)
	}
	// validate the inputs
	if err := cmd.Ok(); err != nil {
		return nil, err
	}

	// set cmd store connection
	setStoreConnection(cmd)

	return cmd, nil
}

func setStoreConnection(command *schema.Command) {
	switch command.Command {
	case string(schema.CMDCreateParkingLot):
		command.Connection = Store.CreateParkingLot()
	case string(schema.CMDPark):
		command.Connection = Store.Park()
	case string(schema.CMDStatus):
		command.Connection = Store.Status()
	case string(schema.CMDHelp):
		command.Connection = Store.Help()
	case string(schema.CMDShellHistory):
		command.Connection = Store.ShellHistory()
	case string(schema.CMDParkingHistory):
		command.Connection = Store.ParkHistory()
	}
}
