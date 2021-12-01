package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
)

//go:embed input.txt
var input []byte

func main() {
	var depths []int
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		d, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		depths = append(depths, d)
	}

	var count int
	for i := 0; i < len(depths)-1; i++ {
		if depths[i] < depths[i+1] {
			count++
		}
	}

	fmt.Println(count)

	var curr, prev, triple int
	for i := 0; i < len(depths)-2; i++ {
		prev = curr
		curr = depths[i] + depths[i+1] + depths[i+2]
		if prev < curr && prev > 0 {
			triple++
		}
	}
	fmt.Println(triple)
}
