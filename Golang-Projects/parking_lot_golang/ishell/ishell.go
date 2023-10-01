package ishell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"parking_lot/schema"
	"parking_lot/store"
	"parking_lot/utils"
)

var (
	shell *schema.IShell
	// Store holds the store conn interface
	Store store.Store
)

// InitIshell inits the conn
func InitIshell() {
	Store = store.NewStore()
	shell = &schema.IShell{}
}

// Start attempts to create new interactive session
// the process reads the commands until it finds `exit` command
//
func Start() {
	if shell == nil {
		InitIshell()
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		cmdInputStr, err := reader.ReadString(utils.EndLineDelim)
		if err != nil {
			break
		}
		cmdInputStr = strings.TrimRight(cmdInputStr, utils.NewLineDelim)
		if strings.TrimSpace(cmdInputStr) == "" {
			continue
		}

		shell.RecordHistory(cmdInputStr)
		// process the commands
		cmd, err := Process(cmdInputStr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// map shell history with command
		cmd.RecordShellHistory(shell.History)
		if cmd.IsExit() {
			greeting()
			break
		}
		// Execute the command
		response, err := cmd.Connection.Execute(cmd)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(response)
	}
}
