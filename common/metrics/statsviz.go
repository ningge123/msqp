package metrics

import (
	"github.com/arl/statsviz"
	"net/http"
)

func Serve(add string) error {
	mux := http.NewServeMux()
	err := statsviz.Register(mux)
	if err != nil {
		return err
	}

	if err := http.ListenAndServe(add, mux); err != nil {
		return err
	}

	return nil
}
