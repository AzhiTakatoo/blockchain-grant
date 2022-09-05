package main

//链码启动必须通过调用 shim 包中的 Start 函数，传递一个类型为 Chaincode 的参数。
//该参数是一个接口类型，有两个重要的函数 Init 与 Invoke 。
import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/lib"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/routers"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/utils"
	"time"
)

type BlockChainGrant struct {
}

// Init 链码初始化，实例化/升级链码时被自动调用
// 在该方法中实现链码初始化或升级时的处理逻辑
// 编写时可灵活使用stub中的API
func (t *BlockChainGrant) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// println 函数的输出信息会出现在链码容器的日志中
	fmt.Println("链码init")

	timeLocal, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//修改
	var wyuUserIds = [3]string{
		"3118002204",
		"1234567890",
		"0000000000",
	}
	var wyuUserNames = [3]string{"曾治评", "管理员", "测试员"}
	var wyuPasswd = [3]string{"zp43", "8888", "0000"}
	//初始化账号数据
	//修改
	for i, val := range wyuUserIds {
		wyuUser := &lib.WyuUser{
			WyuUserId:   val,
			WyuUserName: wyuUserNames[i],
			WyuPasswd:   wyuPasswd[i],
		}
		// 写入账本
		if err := utils.WriteLedger(wyuUser, stub, lib.WyuUserKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainGrant) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "queryRegisterCertify":
		return routers.QueryRegisterCertify(stub, args)
	case "createWyuUser":
		return routers.CreateWyuUser(stub, args)
	case "queryWyuUser":
		return routers.QueryWyuUser(stub, args)
	case "createProofMaterial":
		return routers.CreateProofMaterial(stub, args)
	case "queryProofMaterialList":
		return routers.QueryProofMaterialList(stub, args)
	case "queryProofMaterialOnly":
		return routers.QueryProofMaterialOnly(stub, args)
	case "queryProofCertifyList":
		return routers.QueryProofCertifyList(stub, args)
	case "createPhotoMaterial":
		return routers.CreatePhotoMaterial(stub, args)
	case "queryPhotoMaterialList":
		return routers.QueryPhotoMaterialList(stub, args)
	case "updateProofMaterial":
		return routers.UpdateProofMaterial(stub, args)
	case "createQueryStipendRanking":
		return routers.CreateQueryStipendRanking(stub, args)
	case "queryAwardList":
		return routers.QueryAwardList(stub, args)
	case "createVote":
		return routers.CreateVote(stub, args)
	case "queryVote":
		return routers.QueryVote(stub, args)
	case "queryVoteOnly":
		return routers.QueryVoteOnly(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainGrant))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
