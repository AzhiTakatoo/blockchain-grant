package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/lib"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/utils"
	"strconv"
	"time"
)

//查询文件类型哈希
func QueryPhotoMaterialList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var photoMaterialList []lib.PhotoMaterial
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.PhotoMaterialKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var photoMaterial lib.PhotoMaterial
			err := json.Unmarshal(v, &photoMaterial)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			photoMaterialList = append(photoMaterialList, photoMaterial)
		}
	}
	proofMaterialByte, err := json.Marshal(photoMaterialList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-序列化出错: %s", err))
	}
	return shim.Success(proofMaterialByte)
}
func CreatePhotoMaterial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("参数个数不满足")
	}
	wyuUserId := args[0]
	photo := args[1]
	if wyuUserId == "" || photo == "" {
		return shim.Error("参数存在空值")
	}
	//判断学生是否存在
	resultWyuUserId, err := utils.GetStateByPartialCompositeKeys(stub, lib.WyuUserKey, []string{wyuUserId})
	if err != nil || len(resultWyuUserId) != 1 {
		return shim.Error(fmt.Sprintf("学生信息验证失败%s", err))
	}
	photoMaterial := &lib.PhotoMaterial{
		WyuUserId: wyuUserId,
		Photo:     photo,
	}
	// 写入账本
	if err := utils.WriteLedger(photoMaterial, stub, lib.PhotoMaterialKey, []string{photoMaterial.WyuUserId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	photoMaterialByte, err := json.Marshal(photoMaterial)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(photoMaterialByte)
}

// UpdateProofMaterial 更新贫困材料
func UpdateProofMaterial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("参数个数不满足")
	}
	stipendId := args[0]
	annualHouseholdIncome := args[1]
	comprehensiveEvaluation := args[2]
	volunteerTime := args[3]
	//photoMaterial := args[4]
	if stipendId == "" || annualHouseholdIncome == "" || comprehensiveEvaluation == "" || volunteerTime == "" {
		return shim.Error("参数存在空值")
	}
	// 参数数据格式转换
	var formattedAnnualHouseholdIncome float64
	if val, err := strconv.ParseFloat(annualHouseholdIncome, 64); err != nil {
		return shim.Error(fmt.Sprintf("annualHouseholdIncome参数格式转换出错: %s", err))
	} else {
		formattedAnnualHouseholdIncome = val
	}
	var formattedComprehensiveEvaluation float64
	if val, err := strconv.ParseFloat(comprehensiveEvaluation, 64); err != nil {
		return shim.Error(fmt.Sprintf("comprehensiveEvaluation参数格式转换出错: %s", err))
	} else {
		formattedComprehensiveEvaluation = val
	}
	var formattedVolunteerTime int64
	if val, err := strconv.ParseInt(volunteerTime, 10, 64); err != nil {
		return shim.Error(fmt.Sprintf("volunteerTime: %s", err))
	} else {
		formattedVolunteerTime = val
	}
	stipendScore := formattedAnnualHouseholdIncome*0.25 + formattedComprehensiveEvaluation + float64(formattedVolunteerTime)
	resultsUpdate, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProofMaterialKey, []string{stipendId})
	if err != nil || len(resultsUpdate) != 1 {
		return shim.Error(fmt.Sprintf("根据%s获取学生信息失败: %s", stipendId, err))
	}
	var proofMaterial lib.ProofMaterial
	if err = json.Unmarshal(resultsUpdate[0], &proofMaterial); err != nil {
		return shim.Error(fmt.Sprintf("UpdateSellingBySeller-反序列化出错: %s", err))
	}
	if formattedComprehensiveEvaluation >= 120 {
		return shim.Error(fmt.Sprintf("综合测评不能超过120"))

	}
	if formattedVolunteerTime <= 20 {
		return shim.Error(fmt.Sprintf("义工时小于20小时不参与评选助学金"))

	}
	if formattedAnnualHouseholdIncome >= 5000 {
		return shim.Error(fmt.Sprintf("家庭人均年收入超过5000不参与评选助学金"))

	}
	//判断学生是否存在
	resultsStipendId, err := utils.GetStateByPartialCompositeKeys(stub, lib.WyuUserKey, []string{stipendId})
	if err != nil || len(resultsStipendId) != 1 {
		return shim.Error(fmt.Sprintf("学生信息验证失败%s", err))
	}
	//清除原来的的贫困材料信息
	if err := utils.DelLedger(stub, lib.ProofMaterialKey, []string{stipendId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	updateProofMaterial := &lib.ProofMaterial{
		StipendId:               stipendId,
		AnnualHouseholdIncome:   formattedAnnualHouseholdIncome,
		ComprehensiveEvaluation: formattedComprehensiveEvaluation,
		VolunteerTime:           formattedVolunteerTime,
		StipendScore:            stipendScore,
		//PhotoMaterial:           photoMaterial,
	}
	// 写入账本
	if err := utils.WriteLedger(updateProofMaterial, stub, lib.ProofMaterialKey, []string{updateProofMaterial.StipendId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	updateProofMaterialByte, err := json.Marshal(updateProofMaterial)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(updateProofMaterialByte)
}

func CreateProofMaterial(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 4 {
		return shim.Error("参数个数不满足")
	}
	stipendId := args[0]
	annualHouseholdIncome := args[1]
	comprehensiveEvaluation := args[2]
	volunteerTime := args[3]
	var empty []string
	if stipendId == "" || annualHouseholdIncome == "" || comprehensiveEvaluation == "" || volunteerTime == "" {
		return shim.Error("参数存在空值")
	}
	resultsProof, err := utils.GetStateByPartialCompositeKeys(stub, lib.ProofMaterialKey, empty)
	for _, v := range resultsProof {
		if v != nil {
			var proofData lib.ProofMaterial
			err := json.Unmarshal(v, &proofData)
			if err != nil {
				return shim.Error(fmt.Sprintf("CreateProofMaterial-反序列化出错: %s", err))
			}
			if stipendId == proofData.StipendId {
				return shim.Error(fmt.Sprintf("已提交过材料，若需要修改请选择修改申请材料（文本）: %s", err))
			}
		}
	}
	// 参数数据格式转换
	var formattedAnnualHouseholdIncome float64
	if val, err := strconv.ParseFloat(annualHouseholdIncome, 64); err != nil {
		return shim.Error(fmt.Sprintf("annualHouseholdIncome参数格式转换出错: %s", err))
	} else {
		formattedAnnualHouseholdIncome = val
	}
	var formattedComprehensiveEvaluation float64
	if val, err := strconv.ParseFloat(comprehensiveEvaluation, 64); err != nil {
		return shim.Error(fmt.Sprintf("comprehensiveEvaluation参数格式转换出错: %s", err))
	} else {
		formattedComprehensiveEvaluation = val
	}
	var formattedVolunteerTime int64
	if val, err := strconv.ParseInt(volunteerTime, 10, 64); err != nil {
		return shim.Error(fmt.Sprintf("volunteerTime: %s", err))
	} else {
		formattedVolunteerTime = val
	}
	stipendScore := (1200 - formattedAnnualHouseholdIncome*0.2) + formattedComprehensiveEvaluation*1.5 + float64(formattedVolunteerTime)*2

	if formattedComprehensiveEvaluation >= 120 {
		return shim.Error(fmt.Sprintf("综合测评不能超过120"))

	}
	if formattedVolunteerTime <= 20 {
		return shim.Error(fmt.Sprintf("义工时小于20小时不参与评选助学金"))

	}
	if formattedAnnualHouseholdIncome >= 5000 {
		return shim.Error(fmt.Sprintf("家庭人均年收入超过5000不参与评选助学金"))

	}
	//判断学生是否存在
	resultsStipendId, err := utils.GetStateByPartialCompositeKeys(stub, lib.WyuUserKey, []string{stipendId})
	if err != nil || len(resultsStipendId) != 1 {
		return shim.Error(fmt.Sprintf("学生信息验证失败%s", err))
	}
	proofMaterial := &lib.ProofMaterial{
		StipendId:               stipendId,
		AnnualHouseholdIncome:   formattedAnnualHouseholdIncome,
		ComprehensiveEvaluation: formattedComprehensiveEvaluation,
		VolunteerTime:           formattedVolunteerTime,
		StipendScore:            stipendScore,
		//PhotoMaterial:           photoMaterial,
	}
	// 写入账本
	if err := utils.WriteLedger(proofMaterial, stub, lib.ProofMaterialKey, []string{proofMaterial.StipendId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	proofCertify := &lib.ProofCertify{
		StipendId:    stipendId,
		RegisterTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
	local, _ := time.LoadLocation("Local")
	createTimeUnixNano, _ := time.ParseInLocation("2006-01-02 15:04:05", proofCertify.RegisterTime, local)
	//写入账本
	if err := utils.WriteLedger(proofCertify, stub, lib.ProofCertifyKey, []string{proofCertify.StipendId, fmt.Sprintf("%d", createTimeUnixNano.UnixNano())}); err != nil {
		return shim.Error(fmt.Sprintf("将本次注册记录写入账本失败%s", err))
	}

	//将成功创建的信息返回
	proofMaterialByte, err := json.Marshal(proofMaterial)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(proofMaterialByte)
}

func QueryProofCertifyList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var proofCertifyList []lib.ProofCertify
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProofCertifyKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var proofCertify lib.ProofCertify
			err := json.Unmarshal(v, &proofCertify)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			proofCertifyList = append(proofCertifyList, proofCertify)
		}
	}
	proofCertifyListByte, err := json.Marshal(proofCertifyList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-序列化出错: %s", err))
	}
	return shim.Success(proofCertifyListByte)
}

func QueryProofMaterialOnly(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var proofMaterialList []lib.ProofMaterial
	stipendId := args[0]
	if stipendId == "" {
		return shim.Error("参数存在空值")
	}
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.ProofMaterialKey, []string{stipendId})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var proofMaterial lib.ProofMaterial
			err := json.Unmarshal(v, &proofMaterial)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			proofMaterialList = append(proofMaterialList, proofMaterial)
		}
	}
	proofMaterialByte, err := json.Marshal(proofMaterialList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-序列化出错: %s", err))
	}
	return shim.Success(proofMaterialByte)
}

func QueryProofMaterialList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var proofMaterialList []lib.ProofMaterial
	//var empty []string
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.ProofMaterialKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var proofMaterial lib.ProofMaterial
			err := json.Unmarshal(v, &proofMaterial)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			proofMaterialList = append(proofMaterialList, proofMaterial)
		}
	}
	proofMaterialByte, err := json.Marshal(proofMaterialList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-序列化出错: %s", err))
	}
	return shim.Success(proofMaterialByte)
}
