package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

type Client struct {
	Name  string
	Color string
}

var clients = make(map[net.Conn]Client)
var mutex sync.Mutex

var chatColors = []string{
	"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m",
	"\033[91m", "\033[92m", "\033[93m", "\033[94m", "\033[95m", "\033[96m",
	"\033[1;33m", "\033[1;35m", "\033[1;36m",
}

func getRandomColor() string {
	return chatColors[rand.Intn(len(chatColors))]
}

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
			continue
		}
		fmt.Println("Кто-то подключился")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("\033[H\033[2J"))
	conn.Write([]byte("\033[1;36m========================================\033[0m\n"))
	conn.Write([]byte("\033[1;32m       Добро пожаловать в MaxiChat!      \033[0m\n"))
	conn.Write([]byte("\033[1;36m========================================\033[0m\n"))
	conn.Write([]byte("\033[33m- Как тебя зовут? Напиши имя и нажми Enter: \033[0m"))

	scanner := bufio.NewScanner(conn)
	var clientName string
	if scanner.Scan() {
		clientName = scanner.Text()
	} else {
		return
	}

	mutex.Lock()
	clientColor := getRandomColor()
	clients[conn] = Client{Name: clientName, Color: clientColor}

	timestamp := time.Now().Format("15:04:05")
	joinMsg := fmt.Sprintf("\033[90m[%s]\033[0m \033[1;32m[+]\033[0m \033[1;32m%s вошел в чат\033[0m\n", timestamp, clientName)
	for connection := range clients {
		if connection != conn {
			connection.Write([]byte(joinMsg))
		}
	}
	mutex.Unlock()

	conn.Write([]byte("\033[32mУспешно! Ты подключен к чату.\033[0m\n\n"))
	fmt.Printf("Клиент представился как: %s\n", clientName)

	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "" {
			continue
		}

		mutex.Lock()
		msgTime := time.Now().Format("15:04:05")
		chatMsg := fmt.Sprintf("\033[90m[%s]\033[0m %s%s\033[0m: %s\n", msgTime, clientColor, clientName, msg)
		for connection := range clients {
			if connection != conn {
				connection.Write([]byte(chatMsg))
			}
		}
		mutex.Unlock()
	}

	mutex.Lock()
	delete(clients, conn)
	exitTime := time.Now().Format("15:04:05")
	leaveMsg := fmt.Sprintf("\033[90m[%s]\033[0m \033[1;31m[-]\033[0m \033[1;31m%s покинул чат\033[0m\n", exitTime, clientName)
	for connection := range clients {
		connection.Write([]byte(leaveMsg))
	}
	mutex.Unlock()

	fmt.Printf("Клиент %s покинул MaxiChat.\n", clientName)
}
