package dto

import "time"

type APIResponse struct {
	Status    int                    `json:"status"`
	Message   string                 `json:"message"`
	Timestamp time.Time              `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}
