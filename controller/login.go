package controller

import (
	"net/http"
	"io/ioutil"
	"io"
	//"fmt"
	"encoding/json"
	_ "go_api/vo"
	_ "fmt"
	_ "go/types"
	"go_api/vo"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	_ "go_api/model"
)

func LoginCheck(w http.ResponseWriter, r *http.Request) {

	db, dberr := sql.Open("mysql", "root:zjaaod11@tcp(27.1.238.145:3306)/ments_co_kr")
	if dberr != nil {
		log.Fatal(dberr)
	}
	defer db.Close()



	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var request_param vo.LoginApiRequestVo
	if err := json.Unmarshal(body, &request_param); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(request_param.User_id)
	fmt.Println(request_param.Password)
	var user_nm string
	err = db.QueryRow("SELECT user_nm FROM MEMBER_COMMON WHERE user_id = ?", request_param.User_id).Scan(&user_nm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user_nm)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(request_param); err != nil {
		panic(err)
	}



	/*fmt.Print("[BODY] =>")
	fmt.Println(r)*//*
	result := vo.LoginApiRequestVo{"dofta11","zjaaod11"}
	json.NewEncoder(w).Encode(result)*/


}
