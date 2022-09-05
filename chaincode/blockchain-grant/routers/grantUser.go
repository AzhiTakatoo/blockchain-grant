package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/lib"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/utils"
	"time"
)

func CreateWyuUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	wyuUserId := args[0]
	wyuUserName := args[1]
	wyuPasswd := args[2]

	resultsWyuUser, err := utils.GetStateByPartialCompositeKeys(stub, lib.WyuUserKey, args)
	if err != nil || len(resultsWyuUser) == 1 {
		return shim.Error(fmt.Sprintf("学号%s已被注册: %s", wyuUserId, err))
	}

	if wyuUserId == "" || wyuUserName == "" || wyuPasswd == "" {
		return shim.Error("参数存在空值")
	}
	if len([]rune(wyuUserId)) != 10 {
		return shim.Error("学号应为10位数字，输入位数有错")
	}
	//var wyuUser lib.WyuUser
	//不能将名字注册为管理员
	if wyuUserName == "管理员" {
		return shim.Error("用户名字不能为管理员")
	}
	//判断用户是否已存在，不能重复注册，未实现
	registerWyuUser := &lib.WyuUser{
		WyuUserId:   wyuUserId,
		WyuUserName: wyuUserName,
		WyuPasswd:   wyuPasswd,
	}
	// 写入账本
	if err := utils.WriteLedger(registerWyuUser, stub, lib.WyuUserKey, []string{registerWyuUser.WyuUserId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	registerCertify := &lib.RegisterCertify{
		RegisterId:   wyuUserId,
		RegisterName: wyuUserName,
		RegisterTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
	local, _ := time.LoadLocation("Local")
	createTimeUnixNano, _ := time.ParseInLocation("2006-01-02 15:04:05", registerCertify.RegisterTime, local)
	//写入账本
	if err := utils.WriteLedger(registerCertify, stub, lib.RegisterCertifyKey, []string{registerCertify.RegisterId, fmt.Sprintf("%d", createTimeUnixNano.UnixNano())}); err != nil {
		return shim.Error(fmt.Sprintf("将本次注册记录写入账本失败%s", err))
	}
	registerWyuUserByte, err := json.Marshal(registerWyuUser)
	//registerCertifyByte, err := json.Marshal(registerCertify)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(registerWyuUserByte)
}

func QueryRegisterCertify(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var registerList []lib.WyuUser
	var empty []string
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.WyuUserKey, empty)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var register lib.WyuUser
			err := json.Unmarshal(v, &register)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryDonatingList-反序列化出错: %s", err))
			}
			registerList = append(registerList, register)
		}
	}
	registerListByte, err := json.Marshal(registerList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDonatingList-序列化出错: %s", err))
	}
	return shim.Success(registerListByte)
}

// QueryWyuUser 查询用户是否存在且密码正确
func QueryWyuUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var wyuUserList []lib.WyuUser
	//读取后端传入的数据
	wyuUserId := args[0]
	wyuPasswd := args[1]
	wyuUserList := []lib.WyuUser{}
	err1 := 0
	err2 := 0
	err111 := 0
	err222 := 0
	var emptyargs []string
	//将存入区块链的数据读取出来
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.WyuUserKey, emptyargs)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	if wyuUserList != nil {
		fmt.Sprintf("1")
	}

	for _, v := range results {
		if v != nil {
			var wyuUser lib.WyuUser
			//将读取出来的数据转为结构体
			err := json.Unmarshal(v, &wyuUser)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryWyuUser-反序列化出错: %s", err))
			}
			if wyuUser.WyuUserId != wyuUserId {
				err1 = 111
			} else if wyuUser.WyuUserId == wyuUserId && wyuUser.WyuPasswd != wyuPasswd {
				err1 = 0
				err2 = 222
				break
			} else if wyuUser.WyuUserId == wyuUserId && wyuUser.WyuPasswd == wyuPasswd {
				err111 = 111
				err222 = 222
				wyuUserList = append(wyuUserList, wyuUser)
				break
			}
		}
	}
	if err1 == 111 && err111 != 111 {
		err1 = 0
		err111 = 0
		return shim.Error("系统中没有该用户，需要先注册")
	} else if err2 == 222 && err222 != 222 {
		err2 = 0
		err222 = 0
		return shim.Error("用户密码错误")
	}
	grantUserListByte, err := json.Marshal(wyuUserList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryWyuUser-序列化出错: %s", err))
	}
	return shim.Success(grantUserListByte)
}
