package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	us "user-srv/proto/user"

	"github.com/go-log/log"
	"github.com/micro/go-micro/v2/client"
)

var (
	serviceClient us.UserService
)

// Error 错误结构体
type Error struct {
	Code    string `json:"code"`
	Details string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = true

		// 干掉密码返回
		rsp.User.Pwd = ""
		response["data"] = rsp.User
	} else {
		response["success"] = false
		response["error"] = &Error{
			Details: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
