package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Получаем контекст из запроса
	ctx := r.Context()

	// Добавляем значения в контекст
	ctx = context.WithValue(ctx, "key", "value")

	// Запускаем асинхронную задачу с использованием контекста
	go func() {
		// Внутри горутины можно использовать контекст для отмены задачи
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Async task completed")
		case <-ctx.Done():
			fmt.Println("Task canceled")
		}
	}()

	// Имитация долгой операции
	time.Sleep(2 * time.Second)

	// Использование значения из контекста
	if value, ok := ctx.Value("key").(string); ok {
		fmt.Println("Value from context:", value)
	}

	// Отправка ответа
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
