package main

import (
	"errors"
	"fmt"
)

type Manager struct {
	services map[string]bool
}

func NewManager() *Manager {
	a := Manager{
		services: make(map[string]bool),
	}
	a.services["nginx"] = false
	a.services["postgres"] = true
	return &a
}

func (a *Manager) Start(name string) error {
	status, isFound := a.services[name]
	if !isFound {
		return errors.New(fmt.Sprintf("[FATAL] сервис %s не найден", name))
	}

	if status == true {
		return errors.New(fmt.Sprintf("[WARNING] Сервис %s уже работает", name))
	}

	a.services[name] = true
	return nil
}

func main() {
	manager := NewManager()
	if err := manager.Start("nginx"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("[SUCCESS] nginx запущен")
	}

	if err := manager.Start("postgres"); err != nil {
		fmt.Println(err)
	}

	if err := manager.Start("redis"); err != nil {
		fmt.Println(err)
	}
}
