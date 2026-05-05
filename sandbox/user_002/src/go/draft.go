package main

import "fmt"

func classificarNota(nota float64) string {
    if nota >= 9 {
        return "Excelente"
    } else if nota >= 7.0 && nota < 9.0 {
        return "Bom"
    } else if nota >= 5.0 && nota < 7.0 {
        return "Regular"
    } else {
        return "Insuficiente"
    }
}

func main() {
    

    notas := []float64 {9.5, 8.0, 6.5, 4.0}

    for i := range notas {
        fmt.Printf("Nota %.1f: %s\n", notas[i], classificarNota(notas[i]))
    }

}