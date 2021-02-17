package goerr

import (
	"net/http"

	"github.com/vx416/gox/resperr"
)

var (
	ErrUnprocessableEntity  = resperr.NewRespErr(http.StatusUnprocessableEntity)
	ErrParitalUnprocessable = resperr.NewRespErr(422001, http.StatusText(http.StatusUnprocessableEntity))
)
