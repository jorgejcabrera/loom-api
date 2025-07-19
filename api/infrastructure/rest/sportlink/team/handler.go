package team

import (
	"encoding/json"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"log"
	"net/http"
)

type Handler interface {
	TeamCreationScenario(w http.ResponseWriter, r *http.Request)
}

type RestHandler struct {
	TemporalClient client.Client
}

func NewHandler(c client.Client) Handler {
	return &RestHandler{TemporalClient: c}
}

func (h *RestHandler) TeamCreationScenario(w http.ResponseWriter, r *http.Request) {
	workflowOptions := client.StartWorkflowOptions{
		ID:        "sportlink-team-creation" + GenerateRandomID(),
		TaskQueue: "sportlink-task-queue",
	}

	we, err := h.TemporalClient.ExecuteWorkflow(r.Context(), workflowOptions, "InvokeCreationWorkflow")
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
