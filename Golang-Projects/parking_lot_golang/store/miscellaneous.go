package store

import (
	"fmt"
	"strings"

	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
)

// Park History
type parkHistoryStore struct {
	*store
}

// NewParkHistoryStore returns the store object
func NewParkHistoryStore(st *store) *parkHistoryStore {
	pl := &parkHistoryStore{st}
	return pl
}

func (pl *parkHistoryStore) Execute(cmd *schema.Command) (string, error) {
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	if len(ParkingLot.ParkHistory) == 0 {
		return "", errors.ErrNoHistoyFound("parking")
	}
	var parkHistory = []string{fmt.Sprintf("%-5s%-10s%-25s%-10s%-10s", "No.", "Slot No", "Registration Number", "Colour", "CreatedAt")}
	for i, history := range ParkingLot.ParkHistory {
		parkHistory = append(parkHistory, fmt.Sprintf("%-5d%-10d%-25s%-10s%-10s",
			i+1, history.SlotID, history.RegistrationNumber, strings.Title(history.Colour), utils.FormatDateTime(history.CreatedAt)))
	}
	return strings.Join(parkHistory, utils.NewLineDelim), nil
}

// help interface
type helpStore struct {
	*store
}

// NewHelpStore returns the store object
func NewHelpStore(st *store) *helpStore {
	h := &helpStore{st}
	return h
}

func (h *helpStore) Execute(cmd *schema.Command) (string, error) {
	return schema.AllCommendHint, nil
}

// shell History interface
type shellHistoryStore struct {
	*store
}

// NewShellHistoryStore returns the store object
func NewShellHistoryStore(st *store) *shellHistoryStore {
	pl := &shellHistoryStore{st}
	return pl
}

func (pl *shellHistoryStore) Execute(cmd *schema.Command) (string, error) {
	if len(cmd.ShellHistory) == 0 {
		return "", errors.ErrNoHistoyFound("shell")
	}
	var slotHistory = []string{fmt.Sprintf("%-5s%-50s\t%-10s", "No.", "Command", "CreatedAt")}
	for i, history := range cmd.ShellHistory {
		slotHistory = append(slotHistory, fmt.Sprintf("%-5d%-50s\t%-10s", i+1, history.Command, utils.FormatDateTime(history.CreatedAt)))
	}
	return strings.Join(slotHistory, utils.NewLineDelim), nil
}
