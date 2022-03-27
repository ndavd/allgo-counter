//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"io"
	"net/http"
	"syscall/js"
)

const message string = "Go all things!"

func getCount() string {
	res, _ := http.Get("/inc")
	bytes, _ := io.ReadAll(res.Body)
	return string(bytes)
}

func main() {
	// Don't let the program exit
	c := make(chan bool)

	document := js.Global().Get("document")

	h := document.Call("createElement", "h1")
	h.Set("innerText", message)

	styles := document.Call("createElement", "style")
	styles.Set(
		"innerHTML",
		`
    body {
      font-family: monospace;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }

    h1 {
      font-size: 2em;
    }

    p {
      font-size: 3em;
      color: #007d9c;
      font-weight: bold;
    }

    button {
      font-size: 1.2em;
      color: #fff;
      background: #007d9c;
      font-weight: bold;
      padding: 0.5em 3em;
      cursor: pointer;
    }`,
	)

	count := document.Call("createElement", "p")
	count.Set("innerText", getCount())
	count.Set("id", "counter")

	inc := document.Call("createElement", "button")
	inc.Set("innerText", "+1")

	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) any {
		go func() {
			http.Post("/inc", "", bytes.NewBuffer([]byte{}))
			count.Set("innerText", getCount())
		}()
		return nil
	})

	inc.Call("addEventListener", "click", callback)

	document.Get("head").Call("appendChild", styles)
	document.Get("body").Call("appendChild", h)
	document.Get("body").Call("appendChild", count)
	document.Get("body").Call("appendChild", inc)

	<-c
}
