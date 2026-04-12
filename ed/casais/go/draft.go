package main
import "fmt"
func main() {
    
    n := 0

    fmt.Scan(&n) //Leio a quantidade de animais

    animais := make([]int, n) //Vetor que vai guardar o número dos animais

    solteiros := make(map[int]int) //Map para contar quantos solteiros tem e etc
    for i := range animais {
        fmt.Scan(&animais[i]) //Preenchendo o vetor com os números dos animais
    }

    numPares := 0 //Contagem para ver quantos casais foram formados, inicia em 0
    for _, animal := range  animais { //Vou percorrer o vetor de animais (Não preciso do indice). copio o valor do vetor animais para fazer a consulta
        if solteiros[-animal] > 0 { //Cria uma gaveta com o valor tal, esse "-", torna o número positivo se for -(-10), por exemplo. Se for maior que zero, significa que tem tem um par
            solteiros[-animal]-- //Tiro decremento da contagem, o valor 
            numPares++ //Incremento no par, achei um casal, obaaa
        } else {
            solteiros[animal]++ // se não tiver par, vou criar um solteiro ou incrementar mais nos solteiros
        }
    }

    fmt.Println(numPares) //Inprime a quantidade de casais
}
