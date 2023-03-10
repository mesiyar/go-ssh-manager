package lib

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// NewShell
func NewShell(s *Server) {
	log.Println("connect to server", s.Name)
	dsn := fmt.Sprintf("%s:%d", s.Host, s.Port)
	var config *ssh.ClientConfig

	if s.Type == TYPE_PASSWORD {
		config = &ssh.ClientConfig{
			User: s.Username,
			Auth: []ssh.AuthMethod{ssh.Password(s.Password)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	} else {
		key, err := os.ReadFile(s.PrivateKey)
		HandlerErr(err, fmt.Sprintf("read private key %s", s.PrivateKey))
		signer, err := ssh.ParsePrivateKey(key)
		HandlerErr(err, "parse private key")
		config = &ssh.ClientConfig{
			User: s.Username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	}

	client, err := ssh.Dial("tcp", dsn, config)
	HandlerErr(err, "dial")
	session, err := client.NewSession()
	HandlerErr(err, "ssh session")
	defer session.Close()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		fmt.Println("Error making raw terminal: ", err)
		os.Exit(1)
	}
	defer terminal.Restore(fd, state)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 25, 80, modes)
	HandlerErr(err, "set pty")

	err = session.Shell()
	HandlerErr(err, "shell")

	err = session.Wait()
	HandlerErr(err, "complate")

}
