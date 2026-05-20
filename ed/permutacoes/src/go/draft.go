package main
import (
        "fmt"
        "sort"
)


func backtrack(chars []byte, used [] bool, current [] byte, result *[]string) {
    if len(current) == len(chars) {
        *result = append(*result, string(current))
        return
    }

    for i := 0; i < len(chars); i++ {

        if used[i] {
            continue
        }

        used[i] = true

        current = append (current, chars[i])

        backtrack(chars, used, current, result)

        current = current[:len(current) - 1]

        used [i] = false
    }
}

func main() {
    
    var str string

    _, err := fmt.Scan(&str)
    if err != nil {
        return
    }

    chars := []byte(str)

    sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

    used := make([]bool, len(chars))
    var result []string

    backtrack(chars, used, []byte{}, &result)

    for _, p := range result {
        fmt.Println(p)
    }

}
