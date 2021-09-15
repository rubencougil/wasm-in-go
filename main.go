package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func main() {
	js.Global().Set("pretty", pretty())
	js.Global().Set("time", time())
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

func time() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			time := getTime()
			jsDoc := js.Global().Get("document")
			jsonOutputTextArea := jsDoc.Call("getElementById", "jsoninput")
			jsonOutputTextArea.Set("value", time)
		}()

		return nil
	})
}

func getTime() string {
	fmt.Printf("getting time... \n")
	res, err := http.Get("https://worldtimeapi.org/api/timezone/Europe/Madrid")
	if err != nil {
		fmt.Printf("unable to get time: %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("unable to read time from response: %s\n", err)
	}

	return string(body)
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
