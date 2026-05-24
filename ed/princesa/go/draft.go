package main
import "fmt"


func imprimir(f []int, e int) {
	fmt.Print("[ ")
	for i := 0; i < len(f); i++ {
		if e == i{
			fmt.Printf("%d> ", f[i])
		} else {
			fmt.Printf("%d ", f[i])
		}
	}
	fmt.Print("]")
	fmt.Println()
}

func main() {
    var n, e int
	fmt.Scan(&n, &e)

	fila := make([]int, 0, n)

	//crio a fila e preencho de 1 até n

	for i := 1; i <= n; i++ {
		fila = append(fila, i)
	}

	//procurar o indice que quem está com a espada
	var indiceEspada int
	for i := 0; i < n; i++ {
		if fila[i] == e {
			indiceEspada = i
			break
		}
	}



	//enquanto minha fila for maior que 0, ou seja, ter alguém, o meu laço vai continuar
	for len(fila) > 0 {
		//imprimo a pessoa que tá na fila e quem possui a espada
		imprimir(fila, indiceEspada)

		//se tiver só uma pessoa eu paro o loop
		if len(fila) == 1 {
			break
		}

		//% dá a ideia de uma fila circular
		morreIdx := (indiceEspada + 1) % len(fila)

		/*O que isso significa ? corto o indice do que morre, na minha fila, e continuo depois do que eu cortei, os "..." é tipo como se eu tivesse indicando que eu quero que continue, mas exclua o morreIdx, o append só cola tudo*/
		fila  = append(fila[:morreIdx], fila[morreIdx+1:]...)

		//atualiza o index de quem fica com a espada, após matar e passar para o próximo
		indiceEspada = morreIdx % len(fila)

	}

}
