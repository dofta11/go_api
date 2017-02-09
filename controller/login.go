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
)

func LoginCheck(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	/*fmt.Print("[BODY] =>")
	fmt.Println(r)*/
	result := vo.LoginApiRequestVo{"dofta11","zjaaod11"}
	json.NewEncoder(w).Encode(result)


}
