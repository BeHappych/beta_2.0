package controller

import (
	"database/sql"
	//"fmt"
	"net/http"

	//"strconv"
	//"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"github.com/swaggo/swag/example/celler/model"
	//"github.com/swaggo/swag/example/celler/model"
)

type List struct {
	Id        int
	Full_name string
	Birthday  string
	Address   string
}

var database *sql.DB

func (c *Controller) ListAccounts(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	accounts, err := model.AccountsAll(q)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
