package commands

import tea "github.com/charmbracelet/bubbletea"

type MsgDatabaseSelectionUpdate struct {
    value string
    selected bool
}

func CmdDatabaseSelectionUpdate() tea.Msg {
    return MsgDatabaseSelectionUpdate(true)
}
