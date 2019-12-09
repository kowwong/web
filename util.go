package web

import (
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
)

func createFailureResponse(code int, name string, err error) *ResponseData {
	data := &ResponseData{
		Success: false,
		Code:    code,
		Errors: ErrorMessages{
			{
				Name: name,
				Errors: []error{
					err,
				},
			},
		},
	}

	return data
}

func createSuccessResponse(result interface{}) *ResponseData {
	data := &ResponseData{
		Success: true,
		Result:  result,
		Errors:  ErrorMessages{},
	}

	return data
}

// newContext return a web.Context
func newContext(w http.ResponseWriter, r *http.Request, params *Params) *Context {

	ctx := &Context{
		ResponseWriter: w,
		Request:        r,
		paramValues:    params,
	}

	return ctx
}

func contentType(val string) string {
	var ctype string

	if strings.ContainsRune(val, '/') {
		ctype = val
	} else {
		if !strings.HasPrefix(val, ".") {
			val = "." + val
		}
		ctype = mime.TypeByExtension(val)
	}

	return ctype
}

func logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func exists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !info.IsDir()
}
