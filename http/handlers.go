package http

import (
	"RestApi/todo"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// структура для работы с хенделерами
type HTTPHandlers struct {
	toDoList *todo.List
}

func NewHttpHandler(toDolist *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		toDoList: toDolist,
	}
}

/*
pattern: /tasks
method: POST
info: JSON in HTTP request body

succes:
-status code: 201 Created
-response body: JSON represent created task

failed:
staus code: 400 Bad Request, 409,500
response body: JSON with message error + time

*/

func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		errorDto := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errorDto.ToString(), http.StatusBadRequest)
	}

	if err := taskDto.ValidateForCreate(); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
	}

	toDotask := todo.NewTask(taskDto.Title, taskDto.Description)
	if err := h.toDoList.AddTask(toDotask); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskAlreadyExist) {
			http.Error(w, errorDTO.ToString(), 409)
		} else {
			http.Error(w, errorDTO.ToString(), http.StatusInternalServerError)
		}

	}

	b, err := json.MarshalIndent(toDotask, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response: ", err)
	}
	

}

/*
pattern: /tasks/{title}
method: GET
info: pattern

succeed:
status code: 200 OK
responde body: JSON represented found task

failed:
status code: 400,404,500
response body: JSON with message error + time


*/

func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks
method: GET
info: -

succeed:
status code: 200 OK
responde body: JSON represented found tasks

failed:
status code: 400,500
response body: JSON with message error + time

*/

func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks?completed=true
method: GET
info: query params

succeed:
status code: 200 OK
response body: JSON represented found tasks

failed:
status code: 400,404,500
response body: JSON with message error + time
*/

func (h *HTTPHandlers) HandlerGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: tasks/{title}
method: PATCH
info: pattern + JSON request body

succced:
status code: 200 OK
response body: JSON represented changed tasks

failed:
status code: 400,404,500
response body: JSON with message error + time
*/
func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: tasks/{title}
method: DELETE
info: pattern

succced:
status code: 204 No content
response body: JSON represented delete task

failed:
status code: 400,404,500
response body: JSON with message error + time
*/
func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {

}
