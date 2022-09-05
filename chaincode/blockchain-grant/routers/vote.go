package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/lib"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/utils"
	"strconv"
)

func CreateVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	stipendId := args[0]
	voteId := args[1]
	vote := args[2]
	if stipendId == "" || vote == "" || voteId == "" {
		return shim.Error("参数存在空值")
	}
	if stipendId == voteId {
		return shim.Error("不能给自己评分")
	}
	var formattedVote float64
	if val, err := strconv.ParseFloat(vote, 64); err != nil {
		return shim.Error(fmt.Sprintf("annualHouseholdIncome参数格式转换出错: %s", err))
	} else {
		formattedVote = val
	}
	//判断学生是否存在
	result, err := utils.GetStateByPartialCompositeKeys(stub, lib.ProofMaterialKey, []string{stipendId})
	if err != nil || len(result) != 1 {
		return shim.Error(fmt.Sprintf("学生信息验证失败%s", err))
	}
	voteData := &lib.Vote{
		StipendId: stipendId,
		VoteId:    voteId,
		Vote:      formattedVote,
	}
	// 写入账本
	if err := utils.WriteLedger(voteData, stub, lib.VoteKey, []string{voteData.StipendId, voteData.VoteId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	voteDataByte, err := json.Marshal(voteData)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(voteDataByte)
}

func QueryVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//stipendId := args[0]
	voteList := []lib.Vote{}
	var readId string
	var voteAllData lib.Vote
	var voteAll float64
	var count float64
	var averageVote float64
	var empty []string
	//将存入区块链的数据读取出来
	resultsVote, err := utils.GetStateByPartialCompositeKeys2(stub, lib.VoteKey, empty)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	resultsProof, err := utils.GetStateByPartialCompositeKeys(stub, lib.ProofMaterialKey, empty)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	if voteList != nil {
		fmt.Sprintf("1")
	}
	for _, v := range resultsProof {
		if v != nil {
			var proofData lib.ProofMaterial
			err := json.Unmarshal(v, &proofData)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryVote-反序列化出错: %s", err))
			}
			for _, v2 := range resultsVote {
				var voteData lib.Vote
				//将读取出来的数据转为结构体
				err := json.Unmarshal(v2, &voteData)
				if err != nil {
					return shim.Error(fmt.Sprintf("QueryVote-反序列化出错: %s", err))
				}
				fmt.Println(voteData.StipendId)
				if voteData.StipendId == proofData.StipendId {
					fmt.Println("匹配成功")
					readId = voteData.StipendId
					voteAll = voteAll + voteData.Vote
					count++
				} else {
					fmt.Println("学号不匹配")
				}
			}
			averageVote = voteAll / count
			voteAll = 0
			voteAllData = lib.Vote{
				StipendId:   readId,
				AverageVote: averageVote,
				StuNum:      count,
			}
			fmt.Println(count)
			averageVote = 0
			count = 0
			voteList = append(voteList, voteAllData)
		}
	}
	fmt.Println(voteList)
	voteListByte, err := json.Marshal(voteList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryVote-序列化出错: %s", err))
	}
	return shim.Success(voteListByte)
}

func QueryVoteOnly(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	stipendId := args[0]
	voteList := []lib.Vote{}
	var empty []string
	//将存入区块链的数据读取出来
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.VoteKey, empty)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	if voteList != nil {
		fmt.Sprintf("1")
	}
	for _, v := range results {
		if v != nil {
			var voteData lib.Vote
			//将读取出来的数据转为结构体
			err := json.Unmarshal(v, &voteData)
			if err != nil {
				return shim.Error(fmt.Sprintf("反序列化出错: %s", err))
			}
			if voteData.StipendId == stipendId {
				voteList = append(voteList, voteData)
			} else {
				fmt.Println("学号不匹配")
			}
		}
	}
	voteListByte, err := json.Marshal(voteList)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化出错: %s", err))
	}
	return shim.Success(voteListByte)
}
