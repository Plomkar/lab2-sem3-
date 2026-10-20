package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"project/array"
	"project/set"
)

var gSet set.Set

func loadFromFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 || fields[0] != "SET" {
			continue
		}
		for _, s := range fields[1:] {
			if v, err := strconv.Atoi(s); err == nil {
				gSet.Add(v)
			}
		}
	}
}

func saveToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	var arr array.Array
	gSet.ToArray(&arr)
	fmt.Fprint(f, "SET")
	for i := 0; i < arr.Len(); i++ {
		fmt.Fprintf(f, " %d", arr.Get(i))
	}
	fmt.Fprintln(f)
}

func processQuery(query string) {
	fields := strings.Fields(query)
	if len(fields) == 0 {
		fmt.Println("Unknown command: ")
		return
	}
	cmd := fields[0]
	switch cmd {
	case "SETADD":
		if len(fields) < 2 {
			break
		}
		if x, err := strconv.Atoi(fields[1]); err == nil {
			if gSet.Add(x) {
				fmt.Println("TRUE")
			} else {
				fmt.Println("FALSE")
			}
			return
		}
	case "SETDEL":
		if len(fields) < 2 {
			break
		}
		if x, err := strconv.Atoi(fields[1]); err == nil {
			if gSet.Del(x) {
				fmt.Println("TRUE")
			} else {
				fmt.Println("FALSE")
			}
			return
		}
	case "SET_AT":
		if len(fields) < 2 {
			break
		}
		if x, err := strconv.Atoi(fields[1]); err == nil {
			if gSet.At(x) {
				fmt.Println("TRUE")
			} else {
				fmt.Println("FALSE")
			}
			return
		}
	case "PRINT":
		var arr array.Array
		gSet.ToArray(&arr)
		for i := 0; i < arr.Len(); i++ {
			fmt.Printf("%d ", arr.Get(i))
		}
		fmt.Println()
		return
	}
	fmt.Println("Unknown command:", query)
}

func main() {
	filename, query := "", ""
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--file":
			if i+1 < len(args) {
				filename = args[i+1]
				i++
			}
		case "--query":
			if i+1 < len(args) {
				query = args[i+1]
				i++
			}
		}
	}
	gSet.Init(16)
	if filename != "" {
		loadFromFile(filename)
	}
	if query == "" {
		fmt.Println("Usage: task2_set --file <file.data> --query '<SETADD|SETDEL|SET_AT|PRINT> [x]'")
		return
	}
	processQuery(query)
	if filename != "" {
		saveToFile(filename)
	}
	gSet.Clear()
}
