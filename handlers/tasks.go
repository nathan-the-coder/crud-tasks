package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": fmt.Sprintf("Task with id of '%s' doesn't exist.", id),
		})
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	bodyData := utils.ReadBody(r)

	var task store.Task
	if err := json.Unmarshal([]byte(bodyData), &task); err != nil {
		fmt.Println(err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal server error. Please try again later.",
		})
	}

	_, err := h.Store.Create(task.Title, task.Description)
	if err != nil {
		fmt.Println(err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": "Internal Server Error.",
		})
	}

	response := map[string]any{
		"message": "Task Created Successfully",
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
