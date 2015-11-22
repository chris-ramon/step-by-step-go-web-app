### Construyendo una web app con Go desde cero - Parte #2

En esta segunda parte:

- Extendemos `IndexHandler` para que emitir `html` en vez de `text`.
- Aprendemos sobre el control de errores y la estructura de datos `map`.

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

Agregamos el paquete `text/template` de la biblioteca estándar de Go, usaremos
varias funciones de este paquete para que `IndexHandler` pueda emitir
`html` en vez de `text`.

```go
html := "<h1>{{.Title}}<h1/>"
```

Declaramos e inicializamos la variable `html`, la cual tiene solo una etiqueta `h1`
y dentro una variable `.Title` que será reemplazada por un valor arbitrario que
definiremos luego.

> También podemos declarar e inicializar la variable previa de las siguientes formas:

> `var html string = "<h1>{{.Title}}<h1/>"`
>
> `var html string`
>
> `html = "<h1>{{.Title}}<h1/>"`

```go
t, err := template.New("index").Parse(html)
```

Usamos la función `New` para crear un nuevo `template` llamado `index` y
`Parse` para convertir el string `<h1>{{.Title}}<h1/>` a una instancia del tipo `Template`.

```go
t, err := template.New("index").Parse(html)
if err != nil {
    log.Printf("failed to parse index template, error: %v", err)
    return
}
```

El segundo valor que retorna `Parse` es del tipo `error`,
usamos la sentencia `if` para saber si contiene un valor,
si `err` es diferente de `nil` significa que algo falló y
debemos manejar u omitir el error.

> Si deseamos omitir el error usamos `_`:
> `t, _ := template.New(“index”).Parse(html)`

Si existe un error, entonces usamos `log.Printf` para escribir
en la terminal el mensaje: “failed to parse index template, error: %v”
(falló cuando se intento analizar el template index, error: [aquí ira el mensaje de error]).

> Usamos:
> `log.Println` para escribir en el `standard error` por defecto.
> `log.Printf` similar a `log.Println` más la posibilidad de dar formato al string
> y reemplazar variables por valores.

```go
data := map[string]string{
    "Title": "Store App! :)",
}
```

Usamos la estructura de datos del tipo `map` para crear una variable `data`
que contendrá llaves y valores del tipo `string`.

```go
if err := t.Execute(w, data); err != nil {
    log.Printf("failed to execute index template, error: %v", err)
    return
}
```

Usamos la función `Execute` para escribir en `w http.ResponseWriter` el
resultado de ejecutar el `template` `t`, usamos como contexto de la ejecución
la variable `data`, entonces `{{.Title}}` será reemplazado por `Store App! :)`.

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

```
curl http://localhost:8080/
```

###### Terminal

```bash
<h1>Store App! :)<h1/>
```

A continuación el link al código de esta segunda parte:
[https://github.com/chris-ramon/go-workshop/blob/step-2/store/main.go](https://github.com/chris-ramon/go-workshop/blob/step-2/store/main.go)
