package controller

import (
	"net/http"
	"api/services"
)

// APIResponse api回傳模型
type APIResponse struct {
	SysCode int         `json:"sysCode"`
	SysMsg  string      `json:"sysMsg"`
	Data    interface{} `json:"data"`
}

// GetRespones 取得測試回傳
func GetRespones(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response := APIResponse{200, "連結成功", nil}
	services.ResponseWithJSONgo(w, http.StatusOK, response)
}