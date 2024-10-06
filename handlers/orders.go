package handlers

import (
	"encoding/json"
	"net/http"
	"wb/cache"
)

func GetOrderByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	value, exists := cache.Cache.GetCache(id)
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(value)
}
