### Construyendo una web app con Go desde cero - Parte #1

Esta es la primera parte de una serie de guías, que describen paso a paso como construir una aplicación web con Go.

La app web que construiremos es una e-commerce, la cual contiene varias características esenciales para aprender:

- Sintaxis, fundamentos y características importantes de Go.
- Como resolver problemas comunes en web apps usando Go.

Empecemos!
***

Creamos el directorio de nuestro proyecto, llamado `store` (tienda) y seguidamente creamos el archivo principal de nuestro programa, el cual llamaremos `main.go`, este contendrá el código de nuestro programa que se ejecutará primero.
```
mkdir store && cd $_ && touch main.go
```
##### Explicación:
```
mkdir store -- Crea el directorio `store`

&& -- Agrupa comandos shell

touch main.go -- Crea el archivo `main.go`
```
***

A continuación creamos el servidor web principal de nuestra aplicación, agregamos el siguiente código al archivo `main.go`.
```
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
```
package main
```

Cada archivo `*.go` debe pertenecer a un paquete, en este caso `main.go`
pertenece al paquete `main`.

> Todos los programas en Go tienen un paquete `main` (`package main`) y dentro una función `main` (`func main`), los cuales indican el inicio de ejecución de nuestro programa.

```
import (
	"log"
	"net/http"
)
```

Importamos otros paquetes usando el keyword `import`, en este caso importamos dos paquetes de la biblioteca estándar de Go: `log` y `net/http`.


```
func main() {
    log.Println("server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```


Dentro de la función main, primero usamos la función `log.Println` para imprimir en la `terminal` un mensaje de información: `servidor ejecutándose en :8080`.

Seguidamente usamos la función `log.Fatal` para escribir en la terminal el error que posiblemente retorne la función `http.ListenAndServe`.

Finalmente usamos `http.ListenAndServe` para iniciar el servidor web de nuestra app `store`, en este caso indicamos por el momento específicamente que debe atender peticiones en el puerto `8080`, y usamos `nil` como último argumento porque usamos el [HTTP request multiplexer](https://golang.org/pkg/net/http/#ServeMux) por defecto, el cual puede ser reemplazado luego, por ahora es suficiente, ya que nos ayuda a definir las rutas (urls) de manera simple.

***

A continuación agregamos la ruta principal de nuestro servidor web y la función que se encarga de resolver la misma.

```
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

###### Explicación:
```
http.HandleFunc("/", IndexHandler)
```

Usamos la función `http.HandleFunc` para definir la ruta y la función que se ejecuta cuando existe una petición para la misma, definimos como primer argumento `"/"` que indica la raíz del servidor web y como segundo argumento `IndexHandler` que es la función que se ejecutará cuando existan peticiones del tipo: `http://localhost:8080/` .

```
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Store App!"))
}
```

Definimos la función `IndexHandler`, el tiene como parámetros: 
- `w http.ResponseWriter` -- Lo usamos para escribir la respuesta a una petición.
- `r *http.Request` -- Lo usaremos para obtener información sobre una petición.

Finalmente, la función `w.Write` recibe como argumento `bytes`, deseamos escribir el texto `Store App` como respuesta de la petición, entonces necesitamos convertir este texto (`string`) a `bytes`, por eso usamos `[]byte("Store App")`.

> Usamos el carácter `*` para indicar que la variable será de tipo pointer (puntero).

***

Es hora de ejecutar nuestro programa `store`
```
go run main.go
```

###### Terminal
```
2015/11/19 05:10:23 server running on :8080
```

Seguidamente realizamos la primera petición web a nuestro servidor, para ello usamos el comando `curl`.

```
curl http://localhost:8080/
```

###### Terminal
```
Store App!
```

> Usamos el comando `go run` para compilar y ejecutar nuestros programas.

***

A continuación el link al código de esta primera parte:
[https://github.com/chris-ramon/go-workshop/blob/step-1/store/main.go](https://github.com/chris-ramon/go-workshop/blob/step-1/store/main.go)
