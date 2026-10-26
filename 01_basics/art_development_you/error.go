package main

import (
	"fmt"
)

// AppError — наша кастомная структура для ошибок
type AppError struct {
	Err    error  // Изначальная ошибка
	Custom string // Дополнительное сообщение
}

// Реализуем метод Error() для соответствия интерфейсу error
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Custom, e.Err)
}

// Реализуем метод Unwrap(), чтобы иметь возможность достать исходную ошибку
func (e *AppError) Unwrap() error {
	return e.Err
}
