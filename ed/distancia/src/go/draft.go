package main
import "fmt"

func valido(s []rune, pos int, valor int, L int) bool {

	n := len(s)

	for i := pos - L; i <= pos+L; i++ {

		if i < 0 || i >= n || i == pos {
			continue
		}

		if s[i] != '.' && int(s[i]-'0') == valor {
			return false
		}
	}

	return true
}

func resolver(s []rune, L int, pos int) bool {
    if pos == len(s) {
        return true
    }
    if s[pos] != '.' {
        return resolver(s, L, pos + 1)
    }
    for d := 0; d <= L; d++ {

        if valido(s, pos, d, L) {
            s[pos] = rune(d + '0')

            if resolver(s, L, pos + 1) {
                return true
            }
            s[pos] = '.'
        }
    }
    return false
}

func main() {

	var entrada string
	var L int

	fmt.Scan(&entrada)
	fmt.Scan(&L)

	s := []rune(entrada)

    resolver(s, L, 0)

    fmt.Println(string(s))

	
}