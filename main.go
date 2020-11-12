package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/ssh"
)

type Device struct {
	Config  *ssh.ClientConfig
	Client  *ssh.Client
	Session *ssh.Session
	Stdin   io.WriteCloser
	Stdout  io.Reader
	Stderr  io.Reader
}

func (d *Device) Connect() error {
	client, err := ssh.Dial("tcp", "10.2.0.11:22", d.Config)
	if err != nil {
		return err
	}
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	sshIn, err := session.StdinPipe()
	if err != nil {
		return err
	}
	sshOut, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	sshErr, err := session.StderrPipe()
	if err != nil {
		return err
	}
	d.Client = client
	d.Session = session
	d.Stdin = sshIn
	d.Stdout = sshOut
	d.Stderr = sshErr
	return nil
}

func (d *Device) SendCommand(cmd string) error {
	if _, err := io.WriteString(d.Stdin, cmd+"\r\n"); err != nil {
		return err
	}
	return nil
}

func (d *Device) SendConfigSet(cmds []string) error {
	for _, cmd := range cmds {
		if _, err := io.WriteString(d.Stdin, cmd+"\n"); err != nil {
			return err
		}
		fmt.Println("here1")
		// time.Sleep(time.Second)
	}
	return nil
}

func (d *Device) PrintOutput() {
	fmt.Println("in printer")
	r := bufio.NewReader(d.Stdout)
	for {
		// TODO: SR OS does not return EOF
		text, err := r.ReadString('\n')
		fmt.Printf("%s", text)
		if err == io.EOF {
			break
		}
	}
}

func (d *Device) PrintErr() {
	r := bufio.NewReader(d.Stderr)
	for {
		text, err := r.ReadString('\n')
		fmt.Printf("%s", text)
		if err == io.EOF {
			break
		}
	}
}

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

	sw := &Device{Config: cfg}

	fmt.Println("Connecting to ", address)
	if err := sw.Connect(); err != nil {
		log.Fatal(err)
	}
	defer sw.Client.Close()
	defer sw.Session.Close()
	defer sw.Stdin.Close()

	if err := sw.Session.Shell(); err != nil {
		log.Fatal(err)
	}
	// TODO: discard MOTD
	// ioutil.ReadAll(sw.Stdout)

	command := "show version"
	if err := sw.SendCommand(command); err != nil {
		log.Fatal(err)
	}

	sw.PrintOutput()
}
