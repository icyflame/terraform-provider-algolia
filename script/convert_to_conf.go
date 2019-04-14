package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getTypedValue(t string) interface{} {
	v, err := strconv.Atoi(t)
	if err == nil {
		return v
	}

	switch t {
	case "true":
		return true
	case "false":
		return false
	default:
		return t
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Must provide path to the state file as first argument")
		os.Exit(1)
	}

	fileName := os.Args[1]

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error while reading the state file: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(b), "\n")

	condensed := map[string]interface{}{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		components := strings.Split(line, " = ")
		if len(components) < 2 {
			continue
		}

		attrName := components[0]
		attrValue := getTypedValue(components[1])

		if strings.HasSuffix(attrName, ".#") {
			attrName = attrName[:len(attrName)-2]

			value := []interface{}{}

			attrLen := attrValue.(int)

			for j := 0; j < attrLen; j++ {
				arrVal := lines[i+j+1]
				components := strings.Split(arrVal, " = ")
				if len(components) < 2 {
					continue
				}
				attrValue := components[1]
				val := getTypedValue(attrValue)

				value = append(value, val)
			}

			condensed[attrName] = value

			i += attrLen
		} else {
			condensed[attrName] = attrValue
		}
	}

	for k, v := range condensed {
		switch v.(type) {
		case int:
			fmt.Printf("%s = %v\n", k, v)
		case string:
			fmt.Printf("%s = \"%v\"\n", k, v)
		default:
			fmt.Printf("%s = %v\n", k, v)
		}
	}
}
