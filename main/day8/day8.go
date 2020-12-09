package main

import (
	helper "../helper"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, err := os.Open("/home/rap/Projects/dev/AdventOfCode/main/day8/input.txt")
	if err != nil {
		println("Error: ", err)
	}
	ops := helper.ReadeFileByLine(r)
	acc := 0
	var addr []byte

	for i := 0; i < len(ops); i++ {
		op := strings.Split(ops[i], " ")[0]
		param, _ := strconv.Atoi(strings.Split(ops[i], " ")[1])
		if op == "acc" {
			acc += param
		}
		if op == "jmp" {
			i += param
		}
		if !bytes.Contains(addr, []byte{byte(i)}) && op != "nop" {
			addr = append(addr, byte(i))
		} else if op != "nop" {
			break
		}
	}
	println("Acc: ", acc)
}
