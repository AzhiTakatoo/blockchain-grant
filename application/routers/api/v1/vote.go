package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/orangebottle/blockchain-grant/application/blockchain"
	"github.com/orangebottle/blockchain-grant/application/pkg/app"
	"net/http"
	"strconv"
)

type VoteRequestBody struct {
	StipendId string  `json:"stipendId"` //被评分者
	VoteId    string  `json:"voteId"`    //评分者
	Vote      float64 `json:"vote"`      //分数
}

type VoteQueryRequestBody struct {
	StipendId string `json:"stipendId"`
}

func CreateVote(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(VoteRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	bodyBytes = append(bodyBytes, []byte(body.VoteId))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.Vote, 'E', -1, 64)))
	fmt.Println(bodyBytes)
	//调用智能合约
	resp, err := bc.ChannelExecute("createVote", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "createVote成功", data)
}

func QueryVote(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(VoteQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	//传入区块链的数据
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryVote", bodyBytes)
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

func QueryVoteOnly(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(VoteQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	//传入区块链的数据
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryVoteOnly", bodyBytes)
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
