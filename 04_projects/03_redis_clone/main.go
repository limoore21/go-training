package main

import (
	"errors"
	"fmt"
)

type Cache struct {
	storage map[string]string
}

func NewCache() *Cache {
	a := Cache{
		storage: make(map[string]string),
	}
	return &a
}

func (c *Cache) Set(key, value string) error {
	if key == "" {
		return errors.New("пустой ключ недопустим")
	}
	c.storage[key] = value
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	val, isFound := c.storage[key]
	if !isFound {
		return "", errors.New("ключ не найден")
	}
	return val, nil
}

func (c *Cache) Delete(key string) error {
	_, isFound := c.storage[key]
	if !isFound {
		return errors.New("ключ для удаления не существует")
	}

	delete(c.storage, key)
	return nil
}

func main() {
	cache := NewCache()
	fmt.Println("--- СТАРТ СИСТЕМЫ ---")

	err := cache.Set("", "123")
	if err != nil {
		fmt.Printf("[ERROR] Запись провалилась: %v\n", err)
	}

	err = cache.Set("token", "super-secret")
	if err == nil {
		fmt.Printf("[INFO] Данные 'token' успешно сохранены.\n")
	}

	tokenVal, err := cache.Get("token")
	if err != nil {
		fmt.Printf("[ERROR] Сбой чтения: %v\n", err)
	} else {
		fmt.Printf("[SUCCESS] Данные получены. Значение: %s\n", tokenVal)
	}

	_, err = cache.Get("admin_password")
	if err != nil {
		fmt.Printf("[ERROR] Поиск отклонен: %v\n", err)
	}

	err = cache.Delete("admin_password")
	if err != nil {
		fmt.Printf("[ERROR] Сбой удаления: %v\n", err)
	}
}
