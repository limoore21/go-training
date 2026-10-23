package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	keyBytes, err := os.ReadFile("/home/max/.ssh/id_rsa_test")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		log.Fatalf("Ошибка парсинга ключа: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "fedora",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "176.109.109.162:22", config)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer client.Close()

	fmt.Println("Готово! Соединение установлено.")
}
