package wanip

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type SshClient struct {
	Host string
}

func connect(host string) (c *ssh.Client) {

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	key, err := os.ReadFile(homedir + ".ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	KnownHostKeyCallback, err := knownhosts.New(homedir + ".ssh/known_hosts")
	if err != nil {
		log.Fatal("could not create hostkeycallback function: ", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: KnownHostKeyCallback,
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	return client
}
