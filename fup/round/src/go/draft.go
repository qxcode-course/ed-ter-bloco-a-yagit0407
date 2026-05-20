    package main
    import "fmt"



    func main() {
        var operacao string
        var numero float64

        _, err := fmt.Scan(&operacao, &numero)
        if err != nil {
            fmt.Scan(&operacao, &numero)
        }

        inteiro := int(numero)

        switch operacao {
            
        case "f" :
            fmt.Println(inteiro)

        case "c" :
            if numero > float64(inteiro) {
                fmt.Println(inteiro + 1)
            } else {
                fmt.Println(inteiro)
            }

        case "r": 
            resultado := int(numero + 0.5)
            fmt.Println(resultado)
        }
    }
