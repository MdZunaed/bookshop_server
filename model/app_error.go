package model

import "fmt"

type AppError struct {
	Source     string `json:"source"`
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e AppError) Error() string {
	return e.Source + "::" +
		fmt.Sprintf("%d", e.StatusCode) + "::" +
		e.Message + "::" +
		e.Err.Error()
}
