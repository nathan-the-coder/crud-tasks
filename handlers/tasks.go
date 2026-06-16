package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	ctx := r.Context()
	response, err := h.Store.List(ctx)
	if err != nil {
		fmt.Println(err)
		utils.WriteISError(w, fmt.Sprintf("%s", err))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	ctx := r.Context()

	response, err := h.Store.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": fmt.Sprintf("Task with id of '%d' doesn't exist.", id),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	bodyData := utils.ReadBody(r)
	ctx := r.Context()

	var task store.Task
	if err := json.Unmarshal([]byte(bodyData), &task); err != nil {
		utils.WriteISError(w, fmt.Sprintf("%s", err))
		return
	}

	_, err := h.Store.Create(ctx, task.Title, task.Description)
	if err != nil {
		utils.WriteISError(w,fmt.Sprintf("%s", err))
		return
	}
	response := map[string]any{
		"message": "Task Created Successfully",
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)

}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	status := r.PathValue("status")
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)

	ctx := r.Context()

	err := h.Store.Mark(ctx, id, status)
	if err != nil {
		utils.WriteISError(w, fmt.Sprintf("%s", err))
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"message": fmt.Sprintf("Task (%d) updated successfully.", id),
	})
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
