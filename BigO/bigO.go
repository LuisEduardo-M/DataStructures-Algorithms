package main

import "fmt"

func find_sum(my_array []int) int {
	// Metrica
	var sum int
	for _, item := range my_array {
		sum += item
	}
	return sum
	//Big (O) = O(n)
	/*Motivo:
	O 'n' representa o numero de elementos dentro do
	my_array, ou seja, quanto maior o numero de elementos
	dentro do my_array, maior sera o tempo de execucao
	deste algoritmo.
	*/
}

//------------------------------------------------------------------------
/* Big (O) - Tempo Constante - O(1)
O tempo de execucao do algoritmo nao depende dos dados de entrada,
possuindo uma complexidade de tempo constante.
*/

func constante_1() {
	fmt.Println("Hello, world!")
}

func constante_2(n int) int {
	return n
}

func constante_3(data []string) string {
	//Apenas retorna o valor
	return data[0]
}

/* Obs:
Se nao estiver acessando ou modificando um dado de entrada,
muito provavelmente ele tera tempo constante, ou seja, O(1).
*/

//

//------------------------------------------------------------------------
/* Big (O) - Tempo Linear - O(n)
O tempo de execucao do algoritmo depende do valor de entrada.
*/

/* Neste caso, quanto maior a quantidade de itens dentro
de "my_array", ou seja, n, maior sera o tempo de execucao
do algoritmo.
*/
func linear_1(my_array []string) {
	for _, item := range my_array {
		fmt.Println(item)
	}
}

func linear_2(my_array []int) int {
	var max int                     // constante
	for _, item := range my_array { // quanto mais itens, maior o tempo de exec
		if item > max { // constante
			max = item
		}
	}
	return max // constante
}

func linear_3(my_array []int) int {
	var sum int
	for _, item := range my_array {
		sum += item
	}
	return sum
}

/* Vai iterar por todos os itens da array, entao de novo,
quanto maior o numero de itens dentro da array, maior sera
o tempo de execucao do O(n) - Tempo Linear.
*/

//

//------------------------------------------------------------------------
/* Big (O) - Tempo Logaritmico - O(log n)
Um algoritmo tem tempo logarítmico quando ele reduz o tamanho dos dados de entrada a cada iteração. Desta forma, ele não precisa olhar todos os dados de entrada para concluir sua tarefa.
- Busca binaria
*/

//

//------------------------------------------------------------------------
/* Big (O) - Tempo Quadratico - O(n²)
Um algoritmo tem tempo qudratico quando o mesmo possui a necessidade de executar
uma operacao linear para cada valor dentro dos dados de entrada.
*/

func quadratico_1(list1, list2 []string) []string {
	result := []string{}

	for _, item1 := range list1 { // O(n)
		for _, item2 := range list2 { // O(n) = O(n²)
			// 1 item vai passar por todos da lista2 etc...
			if item1 == item2 {
				result = append(result, item1)
			}
		}
	}
	return result
}

// Dica: Normalmente, quando tem um "for" dentro de outro.
//------------------------------------------------------------------------

//

/* Big (O) - Tempo Exponencial - O(2^n)
Um algoritmo tem tempo exponencial quando o crescimento do tempo dobra para cada item adicionado aos dados de entrada.
Exemplo: fibonacci
*/

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

/* Quando o algoritmo resolve o problema de tamanho n recursivamente,
com dois sub-problemas de tamanho n-1, ou seja, no exemplo de fibonacci,
a funcao retorna ela mesma 2 vezes por isso possui o tempo exponencial
O(2^n), caso retornasse ela 3 vezes seria O(3^n).
*/
//------------------------------------------------------------------------

//

/* Big (O) - Tempo Fatorial - O(n!)
Um algoritmo tem tempo fatorial quando cresce de forma fatorial
de acordo com os dados de entrada.
Exemplo de fatorial caso n fosse 8: 8*7*6*5*4*3*2*1
*/

func fatorial_1(max_number int) {
	for i := 0; i < max_number; i++ {
		fmt.Println(i)
		fatorial_1(max_number - 1)
	}
}

// 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1
//------------------------------------------------------------------------
