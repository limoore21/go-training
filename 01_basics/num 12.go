package main

import (
	"errors"
	"fmt"
)

type SecurityManager struct {
	blacklist map[string]string
}

func NewSecurity() *SecurityManager {
	return &SecurityManager{
		blacklist: make(map[string]string),
	}
}

func (s *SecurityManager) BlockIP(ip string, reason string) error {
	if ip == "" {
		return errors.New("IP-адрес не может быть пустым")
	} else if reason == "" {
		return errors.New("причина блокировки обязательна")
	}
	return nil
}

func (s *SecurityManager) CheckAccess(ip string) (bool, string) {
	var reason bool
	if reason == true {
		return false, ""
	} else {
		return true, ""
	}
}

func main() {
	manager := NewSecurity()

	err := manager.BlockIP("192.168.1.50", "DDoS attack")
	if err != nil {
		fmt.Println("Неожиданная ошибка:", err)
	}

	err = manager.BlockIP("", "Suspicious activity")
	if err != nil {
		fmt.Println("Получена ожидаемая ошибка:", err)
	}

	allowed, reason := manager.CheckAccess("192.168.1.50")
	if !allowed {
		fmt.Printf("Доступ запрещен: %s\n", reason)
	} else {
		fmt.Println("Доступ разрешен")
	}

	allowed, reason = manager.CheckAccess("10.0.0.1")
	if !allowed {
		fmt.Printf("Доступ запрещен: %s\n", reason)
	} else {
		fmt.Println("Доступ разрешен")
	}
}
