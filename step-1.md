### Construyendo una web app con Go desde cero - Parte #1

Esta es la primera parte de una serie de guías, que describen paso a paso cómo construir una web app usando [Go](http://golang.org/).

La web app es un e-commerce - con el cual puedes aprender:

- Sintaxis, fundamentos y características importantes de Go.
- Cómo resolver problemas comunes en web apps usando Go.

***

Empecemos!

Crea el directorio de la web app llamado: `store` y el archivo principal del mismo llamado: `main.go` :

```bash
mkdir store && cd $_ && touch main.go
```

##### Explicación:

```
mkdir store -- Crea el directorio `store`

cd $_ -- Cambia el directorio actual al recién creado

&& -- Agrupa comandos shell

touch main.go -- Crea el archivo `main.go`
```

***

Agregar el siguiente código a `main.go` :

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

##### Explicación:

El archivo `main.go` pertenece al paquete `main` :

```go
package main
```

> Todos los programas en Go tienen un paquete `package main` y una función `func main` - este último indica el inicio de ejecución del programa.


Importa dos paquetes de la [biblioteca estándar](https://golang.org/pkg/) de Go: [`log`](https://golang.org/pkg/log/) y [`net/http`](https://golang.org/pkg/net/http/):

```go
import (
	"log"
	"net/http"
)
```

Define una función `func main`, `log.Println` escribirá en la `terminal` el string: `server running on :8080` y `log.Fatal` escribirá el potencial error al invocar `http.ListenAndServe`.

Finalmente `http.ListenAndServe` inicia un servidor web en el puerto `8080` - se usa `nil` como último argumento para usar el [HTTP request multiplexer](https://golang.org/pkg/net/http/#ServeMux) por defecto:

```go
func main() {
    log.Println("server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

***


Editar el código en `main.go` :

```go
package main

import (
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Store App!"))
}

func main() {
	log.Println("server running on :8080")
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

##### Explicación:

Define un [`handler`](https://golang.org/pkg/net/http/#Handler) para la raíz `/`

`IndexHandler` se ejecutará cuando existan [`requests`](https://golang.org/pkg/net/http/#Request) del tipo: `http://localhost:8080/` :

```go
http.HandleFunc("/", IndexHandler)
```

Define la función `func IndexHandler` - siendo sus parámetros: 
- `w http.ResponseWriter` -- Para escribir la respuesta al `request`.
- `r *http.Request` -- Para obtener información sobre el `request`.

Finalmente se invoca `w.Write` - está función recibe como argumento un arreglo de tipo `bytes` :

```go
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Store App!"))
}
```

> Se usa el carácter `*` para indicar que la variable es del tipo [`pointer`](https://tour.golang.org/moretypes/1).

***

Hora de ejecutar el programa `store` :

```go
go run main.go
```

##### Terminal
```
2015/11/19 05:10:23 server running on :8080
```

Realiza un `request` al servidor web usando el comando `curl` :

```
curl http://localhost:8080/
```

##### Terminal
```
Store App!
```

> El comando: [`go run`](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program) compila y ejecuta los programas.

***

A continuación el link al código de esta primera parte:

[https://github.com/chris-ramon/go-workshop/blob/step-1/store/main.go](https://github.com/chris-ramon/go-workshop/blob/step-1/store/main.go)
