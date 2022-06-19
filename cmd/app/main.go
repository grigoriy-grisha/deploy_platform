package main

import (
	"bytes"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func main() {
	// ssh config
	hostKeyCallback, err := knownhosts.New("C:\\Users\\programmer\\.ssh\\known_hosts")
	//if err != nil {
	//	log.Fatal(err)
	//}
	config := &ssh.ClientConfig{
		User: "sshuser",
		Auth: []ssh.AuthMethod{
			ssh.Password("123"),
		},
		HostKeyCallback: hostKeyCallback,
	}
	// connect ot ssh server
	client, err := ssh.Dial("tcp", "localhost:2022", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run("ls -la"); err != nil {
		panic("Failed to run: " + err.Error())
	}

	if err := session.Wait(); err != nil {
		panic(err.Error())
	}

	session.
		fmt.Println(b.String())

}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
