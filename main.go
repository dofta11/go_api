package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	usr_err "go_api/error"
)

type User struct {
	Id        string
	AddressId string
}

var Member_list = []*Member{}

type Member struct {
	User_cd int
	User_nm string
}

const VerifyMessage = "verified"

func main() {
	s := NewServer()

	s.HandleFunc("POST", "/member/login_check", func(c *Context) {

		db, dberr := sql.Open("mysql", "root:zjaaod11@tcp(27.1.238.145:3306)/ments_co_kr")
		if dberr != nil {
			log.Fatal(dberr)
		}
		defer db.Close()

		var password string
		var user_nm string

		dberr = db.QueryRow(
			"SELECT password, user_nm FROM MEMBER_COMMON WHERE user_id = ? AND password = ?", c.Params["user_id"], c.Params["password"]).Scan(&password, &user_nm)
		if dberr != nil {
			log.Fatal(dberr)
		}

		if password != c.Params["password"] {
			var jsonErr = usr_err.HttpError{Code: http.StatusUnauthorized, Text: "Password Not Match"}
			c.RenderJson(jsonErr)
		} else {

			json_temp := map[string]interface{}{
				"code": 200,
				"msg":  "SUCCESS",
				"result": map[string]string{
					"user_nm": user_nm,
				},
			}

			c.RenderJson(json_temp)
		}

	})

	s.HandleFunc("GET", "/dice", func(c *Context) {

		/*db, dberr := sql.Open("mysql", "root:zjaaod11@tcp(27.1.238.145:3306)/ments_co_kr")
		if dberr != nil {
			log.Fatal(dberr)
		}
		defer db.Close()

		rows, dberr := db.Query("SELECT user_nm, user_id FROM MEMBER_COMMON WHERE 1=?", 1)

		if dberr != nil {
			log.Fatal(dberr)
		}
		defer rows.Close()
		var user_id string
		var user_nm string
		for rows.Next(){
			err := rows.Scan(&user_nm, &user_id)
			if err != nil{
				log.Fatal(err)
			}
			fmt.Println(user_id, user_nm)
		}*/

		//멤버 설정
		GenerateMember()

		c.RenderTemplate("/view/dice/main.html", map[string]interface{}{"member_list": Member_list})
	})

	s.HandleFunc("GET", "/", func(c *Context) {
		c.RenderTemplate("/view/index.html",
			map[string]interface{}{"time": time.Now()})
	})

	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	})

	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		u := User{Id: c.Params["id"].(string)}
		c.RenderXml(u)
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJson(u)
	})

	s.HandleFunc("POST", "/users", func(c *Context) {
		c.RenderJson(c.Params)
	})

	s.HandleFunc("GET", "/login", func(c *Context) {
		// "login.html" 렌더링
		c.RenderTemplate("/view/login.html",
			map[string]interface{}{"message": "로그인이 필요합니다"})
	})

	s.HandleFunc("POST", "/login", func(c *Context) {
		// 로그인 정보를 확인하여 쿠키에 인증 토큰 값 기록
		if CheckLogin(c.Params["username"].(string), c.Params["password"].(string)) {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:  "X_AUTH",
				Value: Sign(VerifyMessage),
				Path:  "/",
			})
			c.Redirect("/")
		}
		// id와 password가 맞지 않으면 다시 "/login" 페이지 렌더링
		c.RenderTemplate("/view/login.html",
			map[string]interface{}{"message": "id 또는 password가 일치하지 않습니다"})

	})

	s.Use(AuthHandler)

	s.Run(":9090")
}

func GenerateMember() {
	list := []*Member{
		{1, "이승현"},
		{2, "최원혁"},
		{3, "이민선"},
		{4, "박승훈"},
		{5, "윤석빈"},
		{6, "이지혜"},
		{7, "전교진"},
		{8, "김용태"},
		{9, "백선욱"},
		{10, "황종인"},
	}
	Member_list = list
}
func CheckLogin(username, password string) bool {
	// 로그인 처리
	const (
		USERNAME = "tester"
		PASSWORD = "12345"
	)

	return username == USERNAME && password == PASSWORD
}

func AuthHandler(next HandlerFunc) HandlerFunc {
	ignore := []string{"/login", "view/index.html", "/member/login_check", "/dice"}
	return func(c *Context) {
		// url prefix가 "/", /login", "view/index.html"인 경우 auth를 체크하지 않음
		for _, s := range ignore {
			if strings.HasPrefix(c.Request.URL.Path, s) {
				next(c)
				return
			}
		}

		if v, err := c.Request.Cookie("X_AUTH"); err == http.ErrNoCookie {
			// "X_AUTH" 쿠키 값이 없으면 "/login" 으로 이동
			//c.Redirect("/login")

			u := usr_err.HttpError{306, "Not Auth"}
			c.RenderJson(u)
			return

		} else if err != nil {
			// 에러 처리
			c.RenderErr(http.StatusInternalServerError, err)
			return
		} else if Verify(VerifyMessage, v.Value) {
			// 쿠키 값으로 인증이 확인이 되면 다음 핸들러로 넘어감
			next(c)
			return
		}

		// "/login"으로 이동
		c.Redirect("/login")
	}
}

// 인증 토큰 생성
func Sign(message string) string {
	secretKey := []byte("golang-book-secret-key2")
	if len(secretKey) == 0 {
		return ""
	}
	mac := hmac.New(sha1.New, secretKey)
	io.WriteString(mac, message)
	return hex.EncodeToString(mac.Sum(nil))
}

// 인증 토큰 확인
func Verify(message, sig string) bool {
	return hmac.Equal([]byte(sig), []byte(Sign(message)))
}
