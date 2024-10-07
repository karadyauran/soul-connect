package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	migrateBinary := "migrate"
	cmd := exec.Command(migrateBinary, "create", "-ext", "sql", "-dir", "db/migration", "-seq", "init_schema")

	cmd.Dir = "./"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error creating migration:", err)
		return
	}

	fmt.Println("Migration created successfully.")
}
