package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/1/p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	part2(s)
}

func part1(s *bufio.Scanner) {
	var acc int64 = 50
	zc := 0
	for s.Scan() {
		l := s.Text()
		dir := l[0:1]
		var mult int64 = 1
		if dir == "L" {
			mult = -1
		}
		mag, err := strconv.ParseInt(l[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		acc += mag * mult
		acc %= 100
		if acc < 0 {
			acc = 100 + acc
		}
		if acc == 0 {
			zc += 1
		}
	}
	fmt.Println(zc)
}

func part2(s *bufio.Scanner) {
	var acc int64 = 50
	var zc int64 = 0
	for s.Scan() {
		l := s.Text()
		before := acc

		// fmt.Printf("b:%d, l:%s ", before, l)
		dir := l[0:1]
		var mult int64 = 1
		if dir == "L" {
			mult = -1
		}
		mag, err := strconv.ParseInt(l[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		acc += mag * mult
		acc %= 100
		if acc < 0 {
			acc = 100 + acc
		}
		rots := mag / 100
		rem := mag % 100
		zc += rots
		if before != 0 {
			if dir == "L" && rem >= before {
				zc += 1
			} else if dir == "R" && rem >= (100-before) {
				zc += 1
			}
		}

		// fmt.Printf("rots:%d, rem:%d, zc:%d\n", rots, rem, zc)
	}
	fmt.Println(zc)
}
