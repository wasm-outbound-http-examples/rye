// Partially based on https://github.com/refaktor/picorye/blob/35eb4f5fb9e8f33bfebe0a342eeb12a443308c71/main_local.go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/refaktor/picorye/env"
	"github.com/refaktor/picorye/evaldo"
)

var ps *env.ProgramState

func main() {
	ps = env.NewProgramStateNEW()
	ctx := ps.Ctx
	ps.Ctx = env.NewEnv(ctx)

	registerLocalBuiltins()

	const cmd = `"https://httpbin.org/anything" .httpget .print`

	result := evaldo.MinimalEyrEvalString(cmd, ps)

	if result == nil {
		panic("Internal error")
	}
	if ps.ErrorFlag {
		if result.Type() == env.ErrorType {
			fmt.Printf("Error: %s\n", result.Inspect(*ps.Idx))
		} else {
			panic("Evaluation failed")
		}
	}
}

func registerLocalBuiltins() {
	registerLocal := func(word string, builtin env.Builtin) {
		idxwrd := ps.Idx.IndexWord(word)
		ps.Ctx.Set(idxwrd, builtin)
	}

	registerLocal("print", *env.NewBuiltin(
		func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			if arg0 != nil {
				fmt.Println(arg0.Inspect(*ps.Idx))
			}
			return env.NewVoid()
		}, 1, false, false, "Prints a value to stdout"))

	registerLocal("httpget", *env.NewBuiltin(
		func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			url := arg0.Inspect(*ps.Idx)

			resp, err := http.Get(url)
			if err != nil {
				return env.NewError(err.Error())
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)

			return env.NewString(string(body))
		}, 1, false, false, "Makes a HTTP request and returns response body as string"))
}
