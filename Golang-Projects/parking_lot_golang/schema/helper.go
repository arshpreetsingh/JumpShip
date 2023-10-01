package schema

// IShellWelcome holds welcome message string
var IShellWelcome = `
Parking-Lot Solution in GO.

Usage:
        You can create a parking lot with N slots, and you can park the vehicle
in the alloted slot. Tou can even interact with the shell by using the following
commands to get status of slots, parked vehicle info, shell command histories,etc..

Type 'help' to see all available commands to use.
Type 'exit' or 'Ctrl + c' to exit from the iShell.
`

// AllCommendHint holds all the commands list string
var AllCommendHint = `
Available commands:
    ●   create_parking_lot
            To create a parking lot with N slots.
            'create_parking_lot {no.of slots to create}'
            Eg: 'create_parking_lot 6'
            Eg: 'create_parking_lot help' to get help
    ●   park
            To park a vehicle, the system will allocate parking slot to park.
            'park {registration number} { vehicle colur}'
            Eg: 'park​ KA-01-HH-1234​ ​White'
            Eg: 'park help' to get help
    ●   status
            To get the current status of the all parking slots.
            Eg: 'status'
    ●   help
            To get all the availabe commands to use.
            Eg: 'help'
    ●   shell_history
            To get all the list of commands get used with in shell.
            Eg: 'shell_history'
    ●   park_history
            To get all the list of parking happend.
            Eg: 'park_history'
    ●   exit
            To exit from the current iShell.
            Eg: 'exit'
`

// CMDCreateParkingLotHint holds help message for `create_parking_lot`
var CMDCreateParkingLotHint = `
●   create_parking_lot
        To create a parking lot with N slots.
        'create_parking_lot {no.of slots to create}'
        Eg: 'create_parking_lot 6'
`

// CMDParkHint holds help message for `park`
var CMDParkHint = `
●   park
        To park a vehicle, the system will allocate parking slot to park.
        'park {registration number} { vehicle colur}'
        Eg: 'park​ KA-01-HH-1234​ ​White'
`

// CMDstatusHint holds help message for `status`
var CMDstatusHint = `
●   status
        To get the current status of the all parking slots.
        Eg: 'status'
`

// CMDHelpHint holds help message for `help`
var CMDHelpHint = `
●   help
        To get all the availabe commands to use.
        Eg: 'help'
`

// CMDExitHint holds help message for `exit`
var CMDExitHint = `
●   exit
        To exit from the current iShell.
        Eg: 'exit'
`

// CMDShellHistoryHint holds help message for `shell_history`
var CMDShellHistoryHint = `
●   shell_history
        To get all the list of commands get used with in shell.
        Eg: 'shell_history'
`
