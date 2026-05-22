package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	counts := make(map[int]int)
	for _, val := range vet {
		absVal := val
		if val < 0 {
			absVal = -val
		}
		counts[absVal]++
	} 
	var keys []int
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var pairs []Pair
	
	for _, k := range keys {
		pairs = append(pairs, Pair{One: k, Two: counts[k]})
	}
	return pairs
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}
	var res []Pair
	currentVal := vet[0]
	count := 1

	for i := 1 ; i < len(vet); i++ {
		if vet[i] == currentVal {
			count++
		} else {
			res = append(res, Pair {One: currentVal, Two: count})
			currentVal = vet[i]
			count = 1
		}
		
	}
	res = append(res, Pair{One: currentVal, Two: count})
	return res
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))
	for i, val := range vet {
		if val > 0 {
			hasWomanNeighbor := false
			if i > 0 && vet[i - 1] < 0 {
				hasWomanNeighbor = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				hasWomanNeighbor = true
			} 
			if hasWomanNeighbor {
				res[i] = 1
			} else {
				res[i] = 0
			} 
		} else {
			res[i] = 0
		}
	}
	return res
}
	func alone(vet []int) []int {
	res := make([]int, len(vet))
	for i, val := range vet {
		if val > 0 {
			hasWomanNeighbor := false
			if i > 0 && vet[i-1] < 0 {
				hasWomanNeighbor = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				hasWomanNeighbor = true
			}
			if !hasWomanNeighbor {
				res[i] = 1
			} else {
				res[i] = 0
			}
		} else {
			res[i] = 0
		}
	}
	return res
}



func couple(vet []int) int {
	count := 0
	available := make(map[int]int)

	for _, val := range vet {
		
		partner := -val

		if available[partner] > 0 {
			count++
			available[partner]-- 
		} else {
			
			available[val]++
		}
	}
	return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	toRemove := make(map[int]bool)
	for _, pos := range posList {
		toRemove[pos] = true
	}

	var res []int
	for i, val := range vet {
		if !toRemove[i] {
			res = append(res, val)
		}
	}
	return res
}

func clear(vet []int, value int) []int {
	var res []int
	for _, val := range vet {
		if val != value {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
