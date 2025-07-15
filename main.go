package main

// Это небольшая программа, которая демонстрирует работу со стандартным вводом 

import (
	"bufio" // буферизированный ввод-вывод
	"fmt"   // для того, чтобы делать форматированный вывод на экран
	"os"    // для того, чтобы получить доступ к стандартному вводу
	"io"	// 
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file is not sorted")
		}
		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
