package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

type MultiSet struct {
	data []int
	size int
	capacity int
}
func (ms *MultiSet) BetterSearch (value int) (bool, int) {
	left := 0
	right := ms.size - 1

	for left <= right {
		mid := (left + right) / 2
		if ms.data[mid] == value {
			return true, mid
		}
		if value < ms.data[mid] {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}
	return false, left
} 
func NewMultiSet(capacity int) *MultiSet {
	if capacity <= 0 {
		capacity = 10
	}
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}

}
func (ms *MultiSet) grow() {
	ms.capacity *= 2
	newData := make([]int, ms.capacity)

	for i := 0; i < ms.size; i++ {
		newData[i] = ms.data[i]
	}

	ms.data = newData
}

func (ms *MultiSet) Insert (value int) {
	if ms.size == ms.capacity {
		ms.grow()
	}

	_, pos := ms.BetterSearch(value)

	for i := ms.size; i > pos; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[pos] = value
	ms.size++
}



func (ms *MultiSet) String() string {
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func (ms *MultiSet) Contains (value int) bool {
	found, _ := ms.BetterSearch(value)
	return found
}

func (ms *MultiSet) Erase(value int) bool {
	found, pos := ms.BetterSearch(value)

	if !found {
		return false
	}

	for i := pos; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.size--
	return true
}
func (ms *MultiSet) Count(value int) int {
	found, pos := ms.BetterSearch(value)

	if !found {
		return 0
	}

	count := 0

	for i := pos; i >= 0 && ms.data[i] == value; i-- {
		count++
	}

	for i := pos + 1; i < ms.size && ms.data[i] == value; i++ {
		count++
	}

	return count
}
func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}

	count := 1

	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}

	return count
}
func (ms *MultiSet) Clear() {
	ms.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if !ms.Erase(value) {
			fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
