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

type StipendRankingQueryRequestBody struct {
	Power string `json:"power"` //权限
}

type SortStipendCreateRequestBody struct {
	StipendId string `json:"stipendId"` //申请助学金学生ID
	//StipendScore float64 `json:"stipendScore"`
	//Ranking      int64   `json:"ranking"` //排名
}

type SortStipendQueryRequestBody struct {
	StipendId string `json:"stipendId"` //申请助学金学生ID
	//Ranking   int64  `json:"ranking"`   //排名
}

func QueryAwardList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SortStipendQueryRequestBody)
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if GlobalPower != "180806" {
		appG.Response(http.StatusBadRequest, "失败", "权限密码错误")
		return
	}
	var bodyBytes [][]byte
	//传入区块链的数据
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryAwardList", bodyBytes)
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

func CreateQueryStipendRanking(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(StipendRankingQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if GlobalPower != "180806" {
		appG.Response(http.StatusBadRequest, "失败", "权限密码错误")
		return
	}
	var bodyBytes [][]byte
	if body.Power != "" {
		bodyBytes = append(bodyBytes, []byte(body.Power))
		//bodyBytes = append(bodyBytes, []byte((strconv.FormatFloat(body.StipendScore, 'E', -1, 64))))
		//bodyBytes = append(bodyBytes, []byte(strconv.FormatInt(body.Ranking, 10)))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("createQueryStipendRanking", bodyBytes)
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
	appG.Response(http.StatusOK, "CreateStipendRanking成功", data)
}
