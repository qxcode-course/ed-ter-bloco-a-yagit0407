package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	_ = vet
	if len(vet) == 0 {
		return "[]"
	}
    var rec func (v []int) string
	rec = func (v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return strconv.Itoa(v[0]) + ", " + rec(v[1:])
	}
	return "[" + rec(vet) + "]"

	
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return rec(v[1:]) + ", " + strconv.Itoa(v[0])
	}
	return "[" + rec(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	var rec func(v []int)
	rec = func(v []int) {
		if	 len(v) <= 1 {
			return
		}
		v[0], v[len(v) - 1] = v[len(v) - 1], v[0]
		rec(v[1 : len(v)-1])
	}
	rec(vet)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	var rec func (v[]int) int
	rec = func(v []int) int {
		if len(v) == 1 {
			return v[0]
		}
		return v[0] + rec(v[1:])

	}
	return rec(vet)
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	var rec func(v []int) int
	rec = func(v []int) int {
		if len(v) == 1 {
			return v[0]
		}
		return v[0] * rec(v[1:])
	}
	return rec(vet)
	
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
    if len(vet) == 0 {
        return -1 
    }

    var rec func(v []int, idx int) (int, int) // retorna (valor, índice)
    rec = func(v []int, idx int) (int, int) {
        if len(v) == 1 {
            return v[0], idx
        }

        vResto, idxResto := rec(v[1:], idx+1)

        
        if v[0] <= vResto {
            return v[0], idx
        }
        return vResto, idxResto
    }

    _, indiceMenor := rec(vet, 0)
    return indiceMenor
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
