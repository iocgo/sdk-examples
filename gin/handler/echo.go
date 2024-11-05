package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/iocgo/sdk"
)

// @Router(path="/echo")
type EchoHandler struct {
}

// @Inject(lazy="true")
func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

// @GetMapping(path="hi")
func (e *EchoHandler) Hi(gtx *gin.Context) {
	gtx.String(http.StatusOK, "hi ~")
}
