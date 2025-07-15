package main

// Это небольшая программа, которая демонстрирует работу со стандартным вводом 

import (
	"bufio" // буферизированный ввод-вывод
	"fmt"   // для того, чтобы делать форматированный вывод на экран
	"os"    // для того, чтобы получить доступ к стандартному вводу
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var prev string
	for in.Scan() {
		txt := in.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			panic("file not sorted")
		}

		prev = txt

		fmt.Println(txt)
	}
}
