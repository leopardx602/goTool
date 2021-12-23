package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!") //這個寫入到 w 的是輸出到客戶端的
}

func getData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析參數，預設是不會解析的
	fmt.Println(r.Form)
	for k, v := range r.Form { // v is slice
		fmt.Println(k, v)
	}

	data := map[string]int{"a": 1, "b": 2}
	data2, _ := json.Marshal(data)
	w.Header().Set("Conetent_Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data2)
}

func postData(w http.ResponseWriter, r *http.Request) {
	// "Conetent_Type" = "application/json"
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	fmt.Println(user)    //fmt.Println(user["a"])
	fmt.Fprintf(w, "ok") //這個寫入到 w 的是輸出到客戶端的
}

func main() {
	http.HandleFunc("/", index) //設定存取的路由
	http.HandleFunc("/get", getData)
	http.HandleFunc("/post", postData)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
