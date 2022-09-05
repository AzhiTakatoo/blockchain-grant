package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/orangebottle/blockchain-grant/application/blockchain"
	"github.com/orangebottle/blockchain-grant/application/pkg/app"
	"net/http"
)

//type WyuUserIdBody struct {
//	WyuUserId string `json:"wyuUserId"`
//}

type WyuUserRequestBody struct {
	WyuUserId string `json:"wyuUserId"`
	WyuPasswd string `json:"wyuPasswd"`
}

type RegisterRequestBody struct {
	WyuUserId   string `json:"wyuUserId"`   //用户ID
	WyuUserName string `json:"wyuUserName"` //用户名
	WyuPasswd   string `json:"wyuPasswd"`   //用户密码
}

func CreateWyuUser(c *gin.Context) {
	//以下用的postman的form-data
	//wyuUserId := c.PostForm("wyuUserId")
	//wyuUserName := c.PostForm("wyuUserName")
	//wyuPasswd := c.PostForm("wyuPasswd")
	appG := app.Gin{C: c}
	//postman中要raw的json中传参数
	body := new(RegisterRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.WyuUserId))
	bodyBytes = append(bodyBytes, []byte(body.WyuUserName))
	bodyBytes = append(bodyBytes, []byte(body.WyuPasswd))
	//调用智能合约
	resp, err := bc.ChannelExecute("createWyuUser", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryWyuUser(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(WyuUserRequestBody)
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	//传入区块链的数据
	bodyBytes = append(bodyBytes, []byte(body.WyuUserId))
	bodyBytes = append(bodyBytes, []byte(body.WyuPasswd))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryWyuUser", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
