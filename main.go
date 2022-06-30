package main

import "os"

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Нет аргументов")
		return
	}

	for i := 1; i < len(args); i++ {
		fileData, err := os.ReadFile(args[i])
		if err != nil {
			println("Ошибка: ", err.Error())
			return
		}
		println(string(fileData))
	}
}
