package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Definindo flags (argumentos)
	file := flag.String("file", "", "Arquivo de log")
	level := flag.String("level", "ERROR", "Filtrar nível")

	flag.Parse()

	//Verificação básica
	if *file == "" {
		fmt.Println("Erro: você precisa passar o arquivo com --file")
		return
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer f.Close()

	counts := make(map[string]int)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var log Log
		err := json.Unmarshal([]byte(line), &log)
		if err != nil {
			fmt.Println("Linha inválida:", err)
			continue
		}

		counts[log.Level]++

		//Comparação de nível
		if strings.EqualFold(log.Level, *level) {
			fmt.Printf("[%s] %s - %s\n", log.Level, log.Time, log.Event)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
	}

	fmt.Println("\nResumo:")

	for level, count := range counts {
		fmt.Printf("%s: %d\n", level, count)
	}
}
