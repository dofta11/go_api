package api

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
	usr_err "go_api/error"
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
	var test string
	json.Unmarshal(body, test)
	fmt.Println(test)
	if err := json.Unmarshal(body, &request_param); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	var password string
	var user_nm string
	err = db.QueryRow("SELECT password, user_nm FROM MEMBER_COMMON WHERE user_id = ?", request_param.User_id).Scan(&password, &user_nm)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if password != request_param.Password {
		var jsonErr = usr_err.HttpError{Code: http.StatusUnauthorized, Text: "Password Not Match"}
		if err := json.NewEncoder(w).Encode(jsonErr); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusCreated)

		json_temp := map[string]interface{}{
			"code"	: 200,
			"msg"	: "SUCCESS",
			"result": map[string]string{
				"user_nm"	:	user_nm,
			},
		}


		if err := json.NewEncoder(w).Encode(json_temp); err != nil {
			panic(err)
		}
	}

}
