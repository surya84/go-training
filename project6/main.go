package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	Name  string `json:"S_Name"`
	Id    int    `json:"S_Id"`
	Age   int    `json:"S_Age"`
	Grade int    `json:"S-Grade"`
}

var StudentData = map[uint64]Student{
	184: {
		Name:  "surya",
		Id:    48,
		Age:   22,
		Grade: 80,
	},

	146: {
		Name:  "teja",
		Id:    45,
		Age:   23,
		Grade: 82,
	},
}

func main() {

	http.HandleFunc("/students", GetDetails)
	http.HandleFunc("/students/get", GetIdDetails)
	http.ListenAndServe(":8080", nil)
}

func GetDetails(w http.ResponseWriter, r *http.Request) {

	if len(StudentData) == 0 {
		w.Write([]byte("No data"))
	}

	students := make([]Student, 0)

	for _, val := range StudentData {
		students = append(students, val)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// jsonData, err :=
	json.NewEncoder(w).Encode(students)

}

func GetIdDetails(w http.ResponseWriter, r *http.Request) {
	userIdString := r.URL.Query().Get("user_id")

	// w.Write([]byte(userIdString))

	userId, err := strconv.ParseUint(userIdString, 10, 64)

	if err != nil {
		w.Write([]byte("Error in converting "))
		return
	}

	w.WriteHeader(http.StatusAccepted)

	check, ok := StudentData[userId]

	if !ok {
		w.Write([]byte("User Id Not Found"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(check)

	if err != nil {
		fmt.Println("No data ")
	}
	w.Write(data)
}
