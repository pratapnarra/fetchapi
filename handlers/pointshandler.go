package handlers

import(
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"github.com/pratapnarra/fetchapi/models"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	
	models.MapMutex.Lock()
	pnts, exists := models.PointsMap[id]
	models.MapMutex.Unlock()

	if !exists {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	response := models.GetResponse{
		Points: pnts,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(jsonResponse))

}