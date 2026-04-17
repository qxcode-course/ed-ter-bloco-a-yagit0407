package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(cap int) *Set {
	s := &Set{}
	if cap <= 0 {
		cap = 10
	}
	s.capacity = cap
	s.data = make([]int, s.capacity)
	s.size = 0
	return s
}
func (s *Set) String() string {
	if s.size == 0 {
		return "[]"
	}
	texto := "["

	for i := 0; i < s.size; i ++ {
		if i > 0 {
			texto += ", "
		}
		texto += strconv.Itoa(s.data[i])
		}
	texto += "]"
	return texto
	}

func (s *Set) reserve (newCapacity int) {
		if newCapacity <= s.capacity {
			return 
		}
		novo := make([]int, newCapacity) 
		for i := 0 ; i < s.size ; i++ {
			novo[i] = s.data[i]
		}
		s.data = novo
		s.capacity = newCapacity
	}

func (s *Set) insert(value int, index int) {
		if s.size == s.capacity {
			s.reserve(s.capacity * 2)
		}

		for i := s.size; i > index; i-- {
			s.data[i] = s.data[i - 1]
		}
		s.data[index] = value
		s.size++
	}

func (s *Set) Insert (value int) {
		index := 0

		for index < s.size && s.data[index] < value {
			index++
		}

		if index < s.size &&  s.data[index] == value {
			return
		}
		s.insert(value, index)
	}

func (s *Set) Contains(value int ) bool {
		for i := 0; i < s.size; i++ {
			if s.data[i] == value {
				return true
			}

		}
		return false
	}
	
func (s *Set) erase(index int) {
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
}

func (s *Set) Erase(value int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			s.erase(i)
			return true
		}
	}
	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
		fmt.Println("value not found")
			}
		case "contains":
			 value, _ := strconv.Atoi(parts[1])
			 fmt.Println(v.Contains(value))
		case "clear":

		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
