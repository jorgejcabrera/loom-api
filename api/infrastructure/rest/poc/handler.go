package poc

import (
	"encoding/json"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"log"
	"loom-api/api/domain/poc"
	"net/http"
)

type Handler struct {
	TemporalClient client.Client
}

func NewHandler(c client.Client) *Handler {
	return &Handler{TemporalClient: c}
}

func (h *Handler) CreatePocHandler(w http.ResponseWriter, r *http.Request) {
	var p poc.Poc
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        "load-test-workflow-" + GenerateRandomID(),
		TaskQueue: "load-test-task-queue",
	}

	we, err := h.TemporalClient.ExecuteWorkflow(r.Context(), workflowOptions, "LoadTestWorkflow", p)
	if err != nil {
		log.Printf("Error iniciando workflow: %v", err)
		http.Error(w, "Error iniciando workflow", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"workflowID": we.GetID(),
		"runID":      we.GetRunID(),
	})
}

func GenerateRandomID() string {
	return uuid.NewString()
}
