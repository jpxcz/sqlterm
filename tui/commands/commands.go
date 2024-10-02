package commands

import tea "github.com/charmbracelet/bubbletea"

type MsgDatabaseSelectionUpdate bool

func CmdDatabaseSelectionUpdate() tea.Msg {
    return MsgDatabaseSelectionUpdate(true)
}
