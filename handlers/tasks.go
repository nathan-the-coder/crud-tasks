package handlers

import (
	"encoding/json"
	"log"
	"maps"
	"net/http"

	"github.com/nathan-the-coder/crud-tasks/store"
	"github.com/nathan-the-coder/crud-tasks/utils"
)

type TaskHandler struct {
	Store *store.TaskStore
}

func NewTaskHandler(store *store.TaskStore) *TaskHandler {
	return &TaskHandler{
		Store: store,
	}
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	response := h.Store.List()
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	response, err := h.Store.Get(id)
	if err != nil {
		log.Fatalf("Failed to get task %s", id)
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	bodyData := utils.ReadBody(r)

	var task store.Task
	if err := json.Unmarshal([]byte(bodyData), &task); err != nil {
		log.Fatalf("Failed to read request body: %s", err)
	}

	_, err := h.Store.Create(task.Title, task.Description)
	if err != nil {
		log.Fatalf("Failed to create task: %s", err)
	}

	response := map[string]any {
		"message": "Task Created Successfully",
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
