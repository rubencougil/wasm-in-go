package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("pretty", pretty())
	<-make(chan bool)
}

func pretty() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		pretty, err := prettyJson(args[0].String())
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}

		jsDoc := js.Global().Get("document")
		jsonOutputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		jsonOutputTextArea.Set("value", pretty)

		return pretty
	})
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
