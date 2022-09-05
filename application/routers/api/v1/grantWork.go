package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
	bc "github.com/orangebottle/blockchain-grant/application/blockchain"
	"github.com/orangebottle/blockchain-grant/application/pkg/app"
	"net/http"
	"os"
	"strconv"
)

var GlobalPower string

type ProofMaterialRequestBody struct {
	StipendId               string  `json:"stipendId"`               //申请助学金学生ID
	AnnualHouseholdIncome   float64 `json:"annualHouseholdIncome"`   //家庭人均年收入
	ComprehensiveEvaluation float64 `json:"comprehensiveEvaluation"` //综合测评
	VolunteerTime           int64   `json:"volunteerTime"`           //义工时长
	StipendScore            float64 `json:"stipendScore"`            //助学金评定得分
	//PhotoMaterial           string  `json:"photoMaterial"`           //照片认证材料
}

type ProofMaterialQueryRequestBody struct {
	StipendId string `json:"stipendId"` //申请助学金学生ID}
}

type ProofCertifyQueryRequestBody struct {
	StipendId string `json:"stipendId"`
}

type PhotoMaterialRequestBody struct {
	WyuUserId     string `json:"wyuUserId"`     //用户ID
	PhotoMaterial string `json:"photoMaterial"` //照片文件
}

type PhotoMaterialQueryRequestBody struct {
	WyuUserId string `json:"wyuUserId"` //用户ID
}

type UpdateProofMaterialRequestBody struct {
	StipendId               string  `json:"stipendId"`               //申请助学金学生ID
	AnnualHouseholdIncome   float64 `json:"annualHouseholdIncome"`   //家庭人均年收入
	ComprehensiveEvaluation float64 `json:"comprehensiveEvaluation"` //综合测评
	VolunteerTime           int64   `json:"volunteerTime"`           //义工时长
	StipendScore            float64 `json:"stipendScore"`            //助学金评定得分
	//PhotoMaterial           string  `json:"photoMaterial"`           //照片认证材料
}

type UpdatePowerRequestBody struct {
	Power string `json:"power"`
}

func UpdateProofMaterial(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdateProofMaterialRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.StipendId == "" || body.ComprehensiveEvaluation == 0 || body.VolunteerTime == 0 || body.AnnualHouseholdIncome == 0 {
		appG.Response(http.StatusBadRequest, "失败", "参数不能存在空值或0值")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.AnnualHouseholdIncome, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.ComprehensiveEvaluation, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatInt(body.VolunteerTime, 10)))
	//调用智能合约
	resp, err := bc.ChannelExecute("updateProofMaterial", bodyBytes)
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

func UpdatePower(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdatePowerRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	GlobalPower = body.Power
	if GlobalPower == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}
	appG.Response(http.StatusOK, "成功", body.Power)
}

func SetPower(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, "成功", GlobalPower)
}

func QueryPower(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, "成功", GlobalPower)
}

func QueryPhotoMaterial(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(PhotoMaterialQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.WyuUserId != "" {
		bodyBytes = append(bodyBytes, []byte(body.WyuUserId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryPhotoMaterialList", bodyBytes)
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
func CreatePhotoMaterial(c *gin.Context) {
	wyuUserId := c.PostForm("wyuUserId")
	//prePhotoMaterial := c.PostForm("prePhotoMaterial")
	appG := app.Gin{C: c}
	//postman中要raw的json中传参数
	body := new(ProofMaterialQueryRequestBody)
	fmt.Println(c)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	//获取文件头
	file, err := c.FormFile("uploadPhotoMaterial")
	fmt.Println(file)
	if err != nil {
		c.String(http.StatusBadRequest, "获取文件请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	fmt.Println("文件名：", fileName)
	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	//var err1 error
	sh := shell.NewShell("localhost:5001")
	photoMaterial, err1 := sh.AddDir(fileName)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err1)
		os.Exit(1)
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(wyuUserId))
	bodyBytes = append(bodyBytes, []byte(photoMaterial))
	fmt.Println(bodyBytes)
	//调用智能合约
	resp, err := bc.ChannelExecute("createPhotoMaterial", bodyBytes)
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

func CreateProofMaterial(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProofMaterialRequestBody)
	fmt.Println(body.VolunteerTime)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.ComprehensiveEvaluation >= 120 {
		appG.Response(http.StatusBadRequest, "失败", "综合测评不能超过120")
		return
	}
	if body.VolunteerTime <= 20 {
		appG.Response(http.StatusBadRequest, "失败", "义工时小于20小时不参与评选助学金")
		return
	}
	if body.AnnualHouseholdIncome >= 5000 {
		appG.Response(http.StatusBadRequest, "失败", "家庭人均年收入超过5000不参与评选助学金")
		return
	}
	if body.StipendId == "" || body.ComprehensiveEvaluation == 0 || body.VolunteerTime == 0 || body.AnnualHouseholdIncome == 0 {
		appG.Response(http.StatusBadRequest, "失败", "参数不能存在空值或0值")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.AnnualHouseholdIncome, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.ComprehensiveEvaluation, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatInt(body.VolunteerTime, 10)))

	//调用智能合约
	resp, err := bc.ChannelExecute("createProofMaterial", bodyBytes)
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

func QueryProofCertify(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProofCertifyQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.StipendId != "" {
		bodyBytes = append(bodyBytes, []byte(body.StipendId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryProofCertifyList", bodyBytes)
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

func QueryProofMaterial(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProofMaterialQueryRequestBody)
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
	//调用智能合约
	resp, err := bc.ChannelQuery("queryProofMaterialList", bodyBytes)
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

func QueryProofMaterialOnly(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ProofMaterialQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.StipendId))
	//调用智能合约
	resp, err := bc.ChannelQuery("queryProofMaterialOnly", bodyBytes)
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
