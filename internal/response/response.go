package response

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Paging  *PaginatedResponse
	Error   *JSONErrorResponse
}

type JSONErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Count      int    `json:"count"`
	TotalPages int    `json:"total_pages"`
	Next       string `json:"next"`
	Previous   string `json:"previous"`
}

func Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JSONResponse{Success: true, Data: data})
}

func Pagination(w http.ResponseWriter, list interface{}, page int, perPage int, count int, totalPages int, next string, previous string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var paging *PaginatedResponse
	total := math.Ceil(float64(count) / float64(perPage))
	if page > 0 {
		paging = &PaginatedResponse{
			Page:       page,
			PerPage:    perPage,
			Count:      count,
			TotalPages: int(total),
			Next:       fmt.Sprintf("%s?page=%d", next, page+1),
			Previous:   fmt.Sprintf("%s?page=%d", previous, page-1),
		}
	}

	json.NewEncoder(w).Encode(JSONResponse{Success: true, Data: list, Paging: paging})

}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JSONResponse{Success: false, Error: &JSONErrorResponse{Code: code, Message: message}})
}
