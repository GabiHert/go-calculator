<p align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=0ABAB5&height=260&section=header&text=Dock&fontSize=90&animation=fadeIn&fontAlignY=38&desc=Tech%20your%20business%20free&descAlignY=56&descAlign=50">

<h1 align="center">Learning Sessions Golang Basics</h1>
</p>

>Nessa primeira sessão iremos abordar conceitos básicos necessários para 
> a criação de um projeto GoLang. Falaremos sobre módulos, pacotes, funções, 
> laços de repetição e dependências. Para demonstrar a usabilidade de cada um dos 
> tópicos de uma forma prática, realizaremos a implementação de uma calculadora.

>1. ### [Instalação](#Instalação)
>2. ### [Módulos](#Módulkos)
>3. ### [Pacotes](#Pacotes)
>4. ### [Funções](#Funções)
>5. ### [Laços](#Laços)
>6. ### [Dependências](#Dependências)
>7. ### [Desafio](#Desafio)

# Instalação
Para realizar a instalação da ferramenta siga os passos descritos [aqui](https://go.dev/doc/install) conforme seu sistema
operacional. Caso não tenha acesso de administrador em sua máquina, opte por criar uma conta no [replit](https://replit.com/signup)
, selecione a linguagem `Go`e inicie sua jornada!
 >Obs: caso utulize o `replit`, seu projeto já iniciará com o módulo go criado, portanto, caso deseje criar um novo exclua o arquivo `go.mod` 

## Módulos
utilizamos módulos para o encapsulamento de um projeto, incluindo os
pacotes e dependências de um projeto, além de especificar o versionamento da 
linguagem.

para criar um módulo utilizamos o comando a seguir, informando o nome do módulo 
a ser criado.
```bash
go mod init calculator
```

após a execução do comando será criado um arquivo denominado `go.mod`, e o mesmo 
terá o seguinte conteúdo:
```
module calculator

go 1.22.2
```

Posteriormente demonstraremos a relação entre o arquivo `go.mod` e o gerenciamento
de dependencias.

## Pacotes

Um projeto Go é composto por pacotes que devem agrupar arquivos 
de acordo com sua funcionalidade. É relacionado diretamente a forma como
o Go lida com importações e exportações entre arquivos. Portanto não
é possível obter arquivos pertencentes a diferentes pacotes em um mesmo
diretório. 

A variáveis, funções e tipos tem escopo público para serem importados por 
outros pacotes quando tem sua letra inicial maiúscula, e são privados quando 
a mesma é minúscula.

A função inicial de todo projeto Go deve ser uma função main dentro de um 
pacote main:
```go
package main

import (
	"calculator/operations"
	"fmt"
)

func main() {
	res := operations.Sum(0.1, 0.2)
	fmt.Printf("Result %s \n", res)
}

```

A função Sum está sendo exportada pelo pacote `operations` e está localizada
no diretório `/operations`
```go
package operations

func Sum(n ...float64) float64 {
	var res float64
	for _, f := range n {
		res += f
	}
	return res

}

```
Se tivéssimos uma variável/tipo/função com a inicial minuscula, ele não poderiaser
importada para outros pacotes, mas estária disponível para uso de outros arquivos
pertencentes ao mesmo pacote e diretório

## Funções
Funções em Go permitem múltiplos retornos. Na função `Divide` abaixo, temos `res` e `err` como retornos,
permitindo que a função retorne tanto o resultado da operação quanto um possível erro. 

Outro recurso importante são os retornos nomeados . Na função `Divide`, `res` e `err` são definidos na assinatura da função. 
Isso permite que esses valores sejam utilizados e retornados diretamente sem a necessidade de serem explicitamente inicializados e
mencionados no retorno final. 

O generics foi introduzido na versão 1.18 do Go. A função `Divide` usa um tipo genérico `T` que pode ser `float64`, `float32`, `int`, `rune` ou `complex64`.
Isso é especificado na assinatura da função com `func Divide[T float64 | float32 | int | rune | complex64](a, b T)`.
Com isso, a função pode ser usada com diversos tipos numéricos, tornando-a mais flexível e reutilizável.

Outra característica do Go é o mecanismo de tratamento de erros. O mesmo é feito de forma explícita. Na função Divide,
se `b` for zero, um erro é criado com `errors.New("can't divide by zero")` e retornado. Se `b` não for zero, a divisão é
realizada e `res` é retornado junto com `nil`, indicando que não houve erro.

```go
package operations

import "errors"

func Divide[T float64 | float32 | int | rune | complex64](a, b T) (res T, err error) {
	//Obs: 'T' is our generic type, it's a type constraint that matches float32, float64, int, and complex64
	//Obs: 'res' and 'err' are named returns, so they aren't explicitly initialized
	if b == 0 {
		err = errors.New("can't divide by zero")
		return //Obs: same as "return err, res" since 'res' and 'err' are named returns
	}

	res = a / b
	return res, nil
}
```

## Laços
Na função `Factorial`, temos um loop for clássico que calcula o fatorial de um número inteiro n. O loop começa com i 
igual a 1 e continua enquanto i for menor ou igual a n. A cada iteração, i é incrementado em 1.

```go
package operations

import "fmt"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("can't calculate a factorial of a negative number")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}
```

Outro exemplo de laço em Go é o for condicional, similar ao while em outras linguagens.
```go
for a < b {
	fmt.Println("while a < b")
}
```

go também suporta loops infinitos como:
```go
for {
	fmt.Println("for(ever)")
}
```

Quando se trabalha com arrays, maps ou slices, o loop for range é muito útil, iiterando sobre cada elemento, proporcionando tanto o índice 
quanto o valor do elemento em cada iteração. 
```go
array := []int{1, 2, 3}
for index, element := range array {
	fmt.Println(index, element)
}
```

A partir da versão 1.22, Go introduziu a possibilidade de usar for range com um número, como demonstrado a seguir:
```go
for range 10 {
	fmt.Println("10 iterations but only for Go v1.22+")
}
```
## Dependências

Em Go, o gerenciamento de dependências é facilitado por várias ferramentas e comandos integrados.
Para adicionar dependências a um projeto, você pode usar o comando `go get <source>`. Este comando busca o pacote especificado e o adiciona ao seu projeto.

Ao adicionar uma dependência, ela é registrada no arquivo `go.mod`, que mantém o controle das dependências do projeto e suas versões. O `go.mod` é essencial para garantir que todos os desenvolvedores do projeto estejam usando as mesmas versões das dependências.
Além disso, o arquivo `go.sum` é gerado para verificar a integridade das dependências. Ele contém hashes das versões dos módulos que foram baixados, garantindo que o código não seja alterado entre instalações.

Para limpar e sincronizar as dependências do seu projeto, o comando `go mod tidy` é utilizado. Ele remove dependências não utilizadas e adiciona quaisquer dependências que estejam faltando, garantindo que seu projeto tenha exatamente o que precisa.

Em nosso código temos uma função denominada `Sum` apresentada em [Pacotes](#Pacotes). Essa função está sujeita a erros de arredondamento. O mesmo pode ser verificado rodando o seguinte treixo de código
na função `main`

```go
package main

import (
	"calculator/operations"
	"fmt"
)

func main() {
	res := operations.Sum(0.1, 0.2)
	fmt.Printf("Result %.20f \n", res)
}

```

Execute a função main com o comando `go run main.go` e será possível analisar o seguinte resultado na saída do terminal:
```
Result 0.30000000000000004441 
```
O resultado esperado seria:
```
Result 0.30000000000000000000 
```

Para que esse resultado seja atingido, vamos fazer uso do pacote [`github.com/shopspring/decimal`](https://github.com/shopspring/decimal).
Para instalar a dependencia utilizamos o comando `go get github.com/shopspring/decimal`. Note que o arquivo `go.mod` agora contém o registro 
da nossa dependência
```
module calculator

go 1.22.2

require github.com/shopspring/decimal v1.4.0 // indirect
```

Para o uso da dependencia foi cirada a função `SafeSum`. 
```go
package operations

import "github.com/shopspring/decimal"

func SafeSum(n ...float64) decimal.Decimal {
	var res decimal.Decimal
	for _, f := range n {
		res = res.Add(decimal.NewFromFloat(f))
	}
	return res
}
```

Para utilizar nossa nova função, devemos editar nossa função `main` da seguinte forma:
```go
package main

import (
	"calculator/operations"
	"fmt"
)

func main() {
	res := operations.SafeSum(0.1, 0.2)
	fmt.Printf("Result %s \n", res.StringFixed(20))
}
``` 

Após executar nosso código podemos verigicar que o resultado obtido foi conforme o esperado:
```
Result 0.30000000000000000000 
```


# Desafio
Crie sua primeira função! adicione um arquivo na pasta `operations` e desenvolva uma função `Multiply`, com a responsabilidade de retornar
o valor resultante da multiplicação de dois ou mais números!
