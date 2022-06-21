package sshClient

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type SSHClient struct {
	SSHParams
}

type SSHParams struct {
	KnownHosts string
	Password   string
	User       string
	Addr       string
}

func NewShhClient(params SSHParams) *SSHClient {
	return &SSHClient{SSHParams: params}
}

func (sshC *SSHClient) Connect() (*ssh.Client, error) {
	hostKeyCallback, err := knownhosts.New(sshC.KnownHosts)

	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User:            sshC.User,
		Auth:            []ssh.AuthMethod{ssh.Password(sshC.Password)},
		HostKeyCallback: hostKeyCallback,
	}

	conn, err := ssh.Dial("tcp", sshC.Addr, config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (sshC *SSHClient) NewSession(connection *ssh.Client) (*ssh.Session, error) {
	session, err := connection.NewSession()

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (sshC *SSHClient) ExecCommand(session *ssh.Session, cmd string) (*bytes.Buffer, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	err := session.Run(cmd)

	if err != nil {
		return nil, err
	}
	return &stdoutBuf, nil

}
