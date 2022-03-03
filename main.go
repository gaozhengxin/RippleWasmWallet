package main

import (
	"fmt"
	"log"
	"syscall/js"
)

func goSetRemote(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goImportAccountFromSeed(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()

	arg1 := args[0].String()
	arg2 := args[1].Int()
	fmt.Printf("%v, %v\n", arg1, arg2)

	w := GetWallet()
	address, err := w.ImportAccountFromSeed(arg1, uint(arg2))
	if err != nil {
		log.Printf("%v\n", err.Error())
		panic(err)
	}

	res = map[string]interface{}{
		"address": address,
	}
	args[len(args)-1:][0].Invoke(js.Null(), res)
	return
}

func goRemoveAccount(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goListAccounts(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goGetBalance(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goBuildPayment(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goCommitPayment(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goCheckTx(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goCheckAddress(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

func goIsValidAddress(_ js.Value, args []js.Value) (res interface{}) {
	res = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(fmt.Sprintf("%v", r))
			args[len(args)-1:][0].Invoke(errorObject)
		}
	}()
	return nil
}

// 添加监听事件
func registerCallbacks() {
	js.Global().Set("importAccountFromSeed", js.FuncOf(goImportAccountFromSeed))
	js.Global().Set("removeAccount", js.FuncOf(goRemoveAccount))
	js.Global().Set("listAccounts", js.FuncOf(goListAccounts))
	js.Global().Set("getBalance", js.FuncOf(goGetBalance))
	js.Global().Set("buildPayment", js.FuncOf(goBuildPayment))
	js.Global().Set("commitPayment", js.FuncOf(goCommitPayment))
	js.Global().Set("checkTx", js.FuncOf(goCheckTx))
	js.Global().Set("checkAddress", js.FuncOf(goCheckAddress))
	js.Global().Set("isValidAddress", js.FuncOf(goIsValidAddress))
}

func main() {
	c := make(chan struct{}, 0)
	println("Go WebAssembly Initialized!")
	registerCallbacks()

	<-c
}
