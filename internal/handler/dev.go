package handler

import (
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"abdullayev13/timeup/internal/utill"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

type Dev struct {
	Service *service.Service
}

func (h *Dev) EskizSetData(c *gin.Context) {
	type EskizData struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	data := new(EskizData)
	err := c.ShouldBind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	if data.Email == "" {
		c.ShouldBindQuery(data)
		if data.Email == "" {
			marshal, _ := json.Marshal(data)
			response.Fail(c, "data not found: "+string(marshal))
			return
		}
	}

	err = utill.SetEskizData(data.Email, data.Password)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "done")
}

func (h *Dev) EskizRefreshToken(c *gin.Context) {
	err := utill.EskizRefreshToken()
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "done")
}

func (h *Dev) DbQuery(c *gin.Context) {
	var query string
	err := c.Bind(query)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	if query == "" {
		all, _ := io.ReadAll(c.Request.Body)
		query = string(all)
	}

	res := make([]map[string]any, 0, 13)
	rows, err := h.Service.Users.Repo.Users.DB.
		Raw(query).Rows()
	if err != nil {
		response.FailErr(c, err)
		return
	}

	res, _ = rowsToMaps(rows)
	_ = rows
	response.Success(c, res)
}

func rowsToMaps(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})

		for i, col := range columns {
			var v interface{}
			val := values[i]

			// Handle NULL values
			if val != nil {
				v = val
			} else {
				v = nil
			}

			entry[col] = v
		}

		result = append(result, entry)
	}

	return result, nil
}
