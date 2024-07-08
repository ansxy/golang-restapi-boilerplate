package request

import (
	"net/http"
	"strconv"

	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
)

type BaseQuery struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Keyword string `json:"keyword"`
	Sort    string `json:"sort"`
}

func NewPaginationQuery(r *http.Request) BaseQuery {
	page := constant.DefaultPage
	perPage := constant.DefaultPerPage

	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))

	}

	if r.URL.Query().Get("per_page") != "" {
		perPage, _ = strconv.Atoi(r.URL.Query().Get("per_page"))
	}

	return BaseQuery{
		Page:    page,
		PerPage: perPage,
		Keyword: r.URL.Query().Get("keyword"),
		Sort:    r.URL.Query().Get("sort"),
	}

}
