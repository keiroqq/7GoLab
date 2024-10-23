package main

import (
	"fmt"
	"net"
)

// main - точка входа в программу клиента.
func main() {
	// Подключаемся к серверу на порту 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close() // Закрываем соединение по завершении работы клиента

	for {
		var message string
		fmt.Print("Введите сообщение (или 'exit' для выхода): ")
		_, err := fmt.Scanln(&message)       // Читаем ввод от пользователя
		if err != nil || message == "exit" { // Выход при ошибке или вводе "exit"
			break
		}

		_, err = conn.Write([]byte(message)) // Отправляем сообщение на сервер
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			return
		}

		buffer := make([]byte, 1024) // Буфер для чтения ответа
		n, err := conn.Read(buffer)  // Читаем ответ от сервера
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}
		// Печатаем ответ от сервера
		fmt.Printf("Ответ сервера: %s", string(buffer[:n]))
	}
}
