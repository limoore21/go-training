package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var clients = make(map[net.Conn]string)
var mutex sync.Mutex

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Не удалось занять порт: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Println("Сервер запущен и слушает порт 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при приеме соединения:", err)
		}
		fmt.Println("Ура! Кто-то подключился к нам!")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Привет! Как тебя зовут? Напиши имя и нажми Enter:\n"))
	scanner := bufio.NewScanner(conn)
	var clientName string
	if scanner.Scan() {
		clientName = scanner.Text()
	} else {
		return
	}
	mutex.Lock()
	clients[conn] = clientName
	mutex.Unlock()

	fmt.Printf("Клиент представился как: %s\n", clientName)

	for scanner.Scan() {
		msg := scanner.Text()
		mutex.Lock()
		for connection, _ := range clients {
			if connection != conn {
				connection.Write([]byte(clientName + ": " + msg + "\n"))
			}

		}
		mutex.Unlock()
	}

	//response := fmt.Sprintf("Приятно познакомиться, %s! Пока-пока!\n", clientName)
	//conn.Write([]byte(response))
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()

	fmt.Printf("Клиент %s покинул чат.\n", clientName)
}
