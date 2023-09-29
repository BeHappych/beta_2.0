package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	//"strconv"
	"log"

	"github.com/gin-gonic/gin"
	//"github.com/swaggo/swag/example/celler/httputil"
	//"github.com/swaggo/swag/example/celler/model"
)

type List struct {
	Id        int
	Full_name string
	Birthday  string
	Address   string
}

var database *sql.DB

func (c *List) ListAccounts(ctx *gin.Context) {
	rows, err := database.Query("select * from Lists")

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	lists := []List{}

	for rows.Next() {

		p := List{
			Id:        0,
			Full_name: "",
			Birthday:  "",
			Address:   "",
		}

		err := rows.Scan(&p.Id, &p.Full_name, &p.Birthday, &p.Address)
		p.Birthday = p.Birthday[0:10]

		if err != nil {
			fmt.Println(err)
			continue
		}

		lists = append(lists, p)
	}
	ctx.JSON(http.StatusOK, lists)
}
