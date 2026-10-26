package main

import (
	"errors"
	"fmt"
)

func method3() error {
	// Имитируем появление ошибки на самом нижнем уровне
	return errors.New("internal database error")
}

func method2() error {
	err := method3()
	if err != nil {
		// Оборачиваем ошибку нижнего уровня в нашу кастомную ошибку
		return &AppError{
			Err:    err,
			Custom: "something went wrong in method2",
		}
	}
	return nil
}

func method1() error {
	// Просто пробрасываем ошибку наверх
	return method2()
}

func main() {
	err := method1()
	if err != nil {
		fmt.Printf("Финальная ошибка: %v\n", err)

		// Если мы хотим достать самую первоначальную ошибку, вызываем Unwrap
		// (Для этого приводим интерфейс error к нашему конкретному типу *AppError)
		if appErr, ok := err.(*AppError); ok {
			fmt.Printf("Изначальная ошибка (через Unwrap): %v\n", appErr.Unwrap())
		}
		return // Прерываем выполнение
	}

	fmt.Println("Success!")
}
