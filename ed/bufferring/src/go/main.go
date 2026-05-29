package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Deque struct {
	data     []int
	front    int
	size     int
	capacity int
}

func NewDeque(inicialCapacidade int) *Deque {
	return &Deque{
		data: make([]int, inicialCapacidade),
		front: 0,
		size: 0,
		capacity: inicialCapacidade,

	}
}

func (d *Deque) Len() int {
	return d.size
}

func (b *Deque) String() string {
	result := []string{}
	for i := range b.size {
		val := b.data[(b.front+i)%b.capacity]
		result = append(result, fmt.Sprint(val))
	}
	return "[" + strings.Join(result, ", ") + "]"
}
func (d *Deque) resize (newCap int) {
	newData := make([]int, newCap)

	for i := 0; i < d.size; i++ {
		oldIndex := (d.front + i) % d.capacity
		newData[i] = d.data[oldIndex]
	}

	d.data = newData
	d.front = 0
	d.capacity = newCap
}

func (d *Deque) getIndex(log int) int {
	return (log + d.front) % d.capacity
}

func (d *Deque) Front() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}
	return d.data[d.front], nil
}

func (d *Deque) Back() (int, error) {
	if d.size == 0 {
		return 0, fmt.Errorf("fail: buffez vazio")
	}
	return d.data[(d.front+d.size-1)%d.capacity], nil
}

func (d *Deque) PushBack(valor int) {
	if d.size == d.capacity {
		d.resize(d.capacity * 2)
	}
	pos := (d.front + d.size) % d.capacity
	d.data[pos] = valor
	d.size += 1
}

func (d *Deque) PushFront(valor int) {
	if d.size == d.capacity {
		d.resize(d.capacity *2)
	}
	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.data[d.front] = valor
	d.size += 1
}

func (d *Deque) PopFront() error {
	if d.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}
	d.front = (d.front + 1) % d.capacity
	d.size -= 1
	return nil
}

func (b *Deque) Debug() string {
	result := make([]string, b.capacity)
	for i := range result {
		result[i] = " _"
		if i == b.front && b.size == 0 {
			result[i] = ">_"
		}
	}
	for i := range b.size {
		index := (b.front + i) % b.capacity
		val := b.data[index]
		prefix := " "
		if i == 0 {
			prefix = ">"
		}
		result[index] = fmt.Sprintf("%s%d", prefix, val)
	}
	return strings.Join(result, " |")
}

func (d *Deque) PopBack() error {
	if d.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}
	d.size -= 1
	return nil
}

func (d *Deque) Clear() {
	d.front = 0
	d.size = 0
	d.data = make([]int, d.capacity)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := &Deque{data: make([]int, 4), capacity: 4}

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(buf.String())
		case "debug":
			fmt.Println(buf.Debug())
		case "size":
			fmt.Println(buf.Len())

		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushFront(num)
			}
		case "pop_back":
			if err := buf.PopBack(); err != nil {
				fmt.Println(err)
			}
		case "pop_front":
			if err := buf.PopFront(); err != nil {
				fmt.Println(err)
			}
		case "front":
			if val, err := buf.Front(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "back":
			if val, err := buf.Back(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "clear":
			buf.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
