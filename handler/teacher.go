package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sbdb-teacher/model"
	"sbdb-teacher/service"
	"strconv"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	teacherId := r.URL.Query().Get("id")
	userId, _ := strconv.ParseUint(teacherId, 10, 64)
	teacher, err := model.Get(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(teacher)
	_, _ = w.Write(resp)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var toCreate struct {
		model.Teacher
		Username string `json:"username"`
		Password string `json:"password"`
	}
	_ = json.Unmarshal(body, &toCreate)
	id, err := service.SignIn(toCreate.Username, toCreate.Password)
	if err != nil {
		log.Println("Cannot create user!")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	toCreate.Teacher.Id = id
	result, err := model.Create(toCreate.Teacher)
	if err != nil {
		log.Println("Create teacher failed")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		log.Println("Teacher ", result.Name, "created")
	}
	response, err := json.Marshal(result)
	_, _ = w.Write(response)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	teacherId := r.URL.Query().Get("id")
	userId, _ := strconv.ParseUint(teacherId, 10, 64)
	err := model.Delete(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	case "DELETE":
		deleteHandler(w, r)
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	all, err := model.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var body []byte
	if len(all) != 0 {
		body, _ = json.Marshal(all)
	} else {
		body = []byte("[]")
	}
	_, _ = w.Write(body)
}
