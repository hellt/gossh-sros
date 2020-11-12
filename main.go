package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {

	const (
		user     = "admin"
		password = "admin"
		address  = "10.2.0.11:22"
	)

	// ssh client config for password base authentication
	cfg := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ignoring host key verification
	}

	client, err := ssh.Dial("tcp", address, cfg)
	if err != nil {
		log.Fatalf("SSH connection failed: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("SSH session failed to open: %v", err)
	}

	// associate Session stdin/out with in-mem buffers
	var stdoutBuf bytes.Buffer
	var stdinBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stdin = &stdinBuf

	// login shell is needed to be spawned for Nokia SR OS
	err = session.Shell()
	if err != nil {
		log.Fatalf("Login shell failed to create: %v", err)
	}

	cmd := "show version"

	// pass the command over SSH
	session.Stdin.Read([]byte(cmd))
	session.Wait()

	// print the stdout contents
	fmt.Println(stdoutBuf.String())
}
