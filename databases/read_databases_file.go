package databases

import (
	"encoding/json"
	"io"
	"os"
	"os/user"
)

type FileDatabases struct {
	Databases []DatabaseCredentials `json:"databases"`
}

func ReadDatabasesJson() ([]DatabaseCredentials, error) {
	currUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open("/Users/" + currUser.Username + "/.config/sqlterm/databases.json")
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var databases FileDatabases
	json.Unmarshal(byteValue, &databases)

	return databases.Databases, nil
}
