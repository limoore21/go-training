package main

import (
	"errors"
	"fmt"
)

func ConnectDB(host string, port int) error {
	if host == "" {
		return errors.New("Критическая ошибка: хост базы данных не указан")
	}
	if port < 1 || port > 65535 {
		return errors.New("Критическая ошибка: недопустимый пор")
	}
	fmt.Println("[INFO] Подключение к бд успешно")
	return nil
}

func main() {
	err := ConnectDB("127.0.0.1", 5432)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
	}
	err = ConnectDB("", 5432)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
	}
	err = ConnectDB("127.0.0.1", 99999)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
	}

}
