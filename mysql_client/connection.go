package mysqlclient

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)


func ExecMySqlClient(username string, host string, password string) {
    log.Printf("connecting to %v host\n", host)
	cmd := exec.Command("mysql", "-u"+username, "-h"+host, "-p"+password)
	var out, stderr bytes.Buffer

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err))
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}
}
