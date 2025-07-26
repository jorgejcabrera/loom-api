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
	we, err := h.TemporalClient.ExecuteWorkflow(r.Context(), client.StartWorkflowOptions{
		ID:        "sportlink-team-creation" + GenerateRandomID(),
		TaskQueue: "sportlink-task-queue",
	}, "InvokeCreationWorkflow")
	if err != nil {
		log.Printf("Error iniciando workflow: %v", err)
		http.Error(w, "Error iniciando workflow", http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(map[string]string{
		"workflowID": we.GetID(),
		"runID":      we.GetRunID(),
	})

	if err != nil {
		log.Printf("Error serializando respuesta: %v", err)
		http.Error(w, "Error interno al serializar respuesta", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	if _, err := w.Write(responseBody); err != nil {
		log.Printf("Error writing responseBody: %v", err)
		http.Error(w, "Error interno al serializar respuesta", http.StatusInternalServerError)
	}
}

func GenerateRandomID() string {
	return uuid.NewString()
}
