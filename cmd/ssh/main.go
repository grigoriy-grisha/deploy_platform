package main

import (
	"cdcd_platform/pkg/sshClient"
	"fmt"
	"log"
)

func main() {
	client := sshClient.NewShhClient(sshClient.SSHParams{
		Addr:       "localhost:22",
		KnownHosts: "C:\\Users\\programmer\\.ssh\\known_hosts",
		User:       "sshuser",
		Password:   "123",
	})

	connection, err := client.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer connection.Close()

	session, err := client.NewSession(connection)

	if err != nil {
		log.Fatal(err)
	}

	buffer, err := client.ExecCommand(session, "ls -la")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buffer)

	defer session.Close()
}
