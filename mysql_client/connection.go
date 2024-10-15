package mysqlclient

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecMySqlClient(username string, host string, password string, formatAsTable bool) {
	fmt.Printf("connecting to %s host\n", host)

	args := []string{"-u" + username, "-h" + host, "-p" + password}

	if formatAsTable {
		args = append(args, "-t")
	}

	cmd := exec.Command("mysql", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Exiting: Error connecting to mysql. Command Output: %+v\n", err)
		os.Exit(1)
	}
}
