### Construyendo una web app con Go desde cero - Parte #2

En esta segunda parte:

- Extenderemos el `IndexHandler` para que pueda responder `html` en vez de `text`.
- Aprenderemos sobre el manejo de errores, la estructura de datos `map`

```go
package main

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	html := "<h1>{{.Title}}<h1/>"
	t, err := template.New("index").Parse(html)
	if err != nil {
		log.Printf("failed to parse index template, error: %v", err)
		return
	}
	data := map[string]string{
		"Title": "Store App! :)",
	}
	if err := t.Execute(w, data); err != nil {
		log.Printf("failed to execute index template, error: %v", err)
		return
	}
}

func main() {
	log.Println("server running on :8080")
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

##### Explicación:
```go
import (
	"log"
	"net/http"
	"text/template"
)
```

Agregamos el paquete `text/template` de la biblioteca estándar de Go, el cuál usaremos
para hacer escribir `html` en `w http.ResponseWriter` de `IndexHandler`.

```go
html := "<h1>{{.Title}}<h1/>"
```

Declaramos e inicializamos la variable `html`, el cual tiene solo una etiqueta `h1`
y dentro una variable `.Title` que será reemplazado por un valor arbitrario que
definiremos seguidamente.

> También podemos declarar e inicializar la variable previa de las siguientes formas:
> var html string = "<h1>{{.Title}}<h1/>"
>
> `var html string`
> html = "<h1>{{.Title}}<h1/>"

```go
t, err := template.New("index").Parse(html)
```

Usamos la función `New` para crear un nuevo `template` llamado `index` y
`Parse` para convertir el string “<h1>{{.Title}}<h1/>” a una instancia `Template`.

```go
t, err := template.New("index").Parse(html)
if err != nil {
    log.Printf("failed to parse index template, error: %v", err)
    return
}
```

El segundo valor que retorna `Parse` es del tipo `error`,
usamos la sentencia `if` para saber si contiene un valor,
si este es `nil` quiere decir que no hubo error al ejecutar
la función `Parse`.

Si existe un error, entonces usamos `log.Printf` para escribir
en la terminal el mensaje: “failed to parse index template, error: %v”
(falló cuando se intento analizar el template index, error: [aquí ira el mensaje de error]).

> Usamos `log.Printf` cuando queremos agregar valores a variables que
estan dentro del string que deseamos imprimir en la terminal.

> En cambio usamos `log.Println` si solo deseamos escribir un string
en la terminal.

```go
data := map[string]string{
    "Title": "Store App! :)",
}
```

Usamos la estructura de datos del tipo `map` para crear una variable `data`
que contendrá llaves y valores del tipo `string`.

Usaremos esta variable para ejecutar y escribir el valor
de la variable `html` que definimos previamente.

```go
if err := t.Execute(w, data); err != nil {
    log.Printf("failed to execute index template, error: %v", err)
    return
}
```

Usamos la función `Execute` para escribir en `w http.ResponseWriter` el
resultado de ejecutar el `template` `t`, usamos como context la
variable `data`, las variables previamente definidas dentro de
la variable `html` serán reemplazadas por las variables
del `map` data.

***

Es hora de ejecutar nuestro programa `store`

```go
go run main.go
```

###### Terminal

```bash
2015/11/21 19:07:11 server running on :8080
```

Seguidamente realizamos una petición web a nuestro servidor, para ello usamos el comando `curl`.

##### Terminal

```bash
<h1>Store App! :)<h1/>
```

A continuación el link al código de esta primera parte:
[https://github.com/chris-ramon/go-workshop/blob/step-2/store/main.go](https://github.com/chris-ramon/go-workshop/blob/step-2/store/main.go)
