# About this repo

This repo has my code following the course Go (Golang): Explorando a Linguagem do Google | Udemy.

Resources:

- [My first Golang program](#first-program)
- [Commands](#commands)
- [If else](#if_else)
- [Else if](#else_if)

# My first Golang program

Go programs need a "main" package to be executable.
This will print "PrimeiroPrograma!" in our console.

```go
package main

import "fmt"

func main() {
	fmt.Print("Primeiro")
	fmt.Print("Programa!")
}
```

The "fmt" package has some functions to print stuff in our console.

## Commands

- go get: works as a package manager
- go help: tells what a command does
- go vet: does some linting and warns

## Const and variables

```go
package main

import "fmt"
import "math"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * math.Pow(raio, 2)
	fmt.Println("Area = ", area);
	const (
		a = 2
		b = 1
	)
	fmt.Println(a,b)
	g,h,i := 2, 3, "epa"
	fmt.Print(g,h,i)
}
```

In golang we can't have a variable that is not used, this raises an <strong>error</strong> not a warning as it is more common.

# Conditionals

## If else

In go the if and else conditionals must be used with an expression:

```go
package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota ", nota)
	} else {
		fmt.Println("Reprovado com nota ", nota)
	}
}

func main() {
	imprimirResultado(8)
	imprimirResultado(5)
}
```

A wrong conditional would be:

```go
if nota >= 6
	fmt.Println("Hello")
```

## Else if

Else if works just like in normal programming languages:

```go
package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 3 && n < 5 {
		return "C"
	} else {
		return "E"
	}

}
func main() {
	fmt.Println(notaParaConceito(9.7))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(5))
}

```

## For loop

In golang we don't have while loop, only the for loop.
Here are some examples:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print("par ")
		} else {
			fmt.Print("impar")
		}
	}

	// infinite foor loop
	for {
		fmt.Println(" para sempre...")
		time.Sleep(time.Second)
	}
}
```
