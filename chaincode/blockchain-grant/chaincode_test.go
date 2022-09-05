package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/lib"
	"testing"
)

func initTest(t *testing.T) *shim.MockStub {
	scc := new(BlockChainGrant)
	stub := shim.NewMockStub("ex01", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	return stub
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) pb.Response {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	return res
}

// 测试链码初始化
func TestBlockChainRealEstate_Init(t *testing.T) {
	initTest(t)
}

// 测试获取用户信息
func Test_QueryWyuUser(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("2、测试获取3118002204数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryWyuUser"),
			[]byte("3118002204"),
			[]byte("zp43"),
		}).Payload)))
}

// 测试创建用户
func Test_CreateWyuUser(t *testing.T) {
	stub := initTest(t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("createWyuUser"),
		[]byte("3118004651"), //学号
		[]byte("朱海城"),        //名字
		[]byte("zhc"),        //密码
	})
	////注册名字不能为管理员
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createWyuUser"),
	//	[]byte("1234567890"), //注册学号
	//	[]byte("管理员"),        //名字
	//	[]byte("8888"),       //密码
	//})
	////参数个数不满足
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createWyuUser"),
	//	[]byte("3118004651"), //学号
	//	[]byte("朱海城"),        //名字
	//})
}

// 测试创建申请材料
func Test_CreateProofMaterial(t *testing.T) {
	stub := initTest(t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("createProofMaterial"),
		[]byte("0000000000"), //学号
		[]byte("1234"),       //收入
		[]byte("48"),         //综测
		[]byte("82"),         //义工时
	})
	////参数个数不满足
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("1234"),       //收入
	//	[]byte("48"),         //综测
	//})
	////参数存在空值
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte(""),   //学号
	//	[]byte("0"),  //收入
	//	[]byte("12"), //综测
	//	[]byte("48"), //义工
	//})
	////综合测评超过120
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("2712"),       //收入
	//	[]byte("122"),        //综测
	//	[]byte("48"),         //义工
	//})
	////义工时小于20小时不参与评选助学金
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("2313"),       //收入
	//	[]byte("115"),        //综测
	//	[]byte("18"),         //义工
	//})
	////家庭人均年收入超过5000不参与评选助学金
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("7313"),       //收入
	//	[]byte("115"),        //综测
	//	[]byte("29"),         //义工
	//})
	////学生不存在
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte("0000012345"), //学号
	//	[]byte("2313"),       //收入
	//	[]byte("115"),        //综测
	//	[]byte("29"),         //义工
	//})
}

//手动创建一些申请材料
func checkCreateProofMaterial(stub *shim.MockStub, t *testing.T) []lib.ProofMaterial {
	var proofMaterialList []lib.ProofMaterial

	var proofMaterial lib.ProofMaterial
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createProofMaterial"),
		[]byte("3118002204"), //学号
		[]byte("3000"),       //收入
		[]byte("50"),         //综合测评
		[]byte("60"),         //义工时
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &proofMaterial)
	proofMaterialList = append(proofMaterialList, proofMaterial)
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createProofMaterial"),
		[]byte("0000000000"), //学号
		[]byte("2000"),       //收入
		[]byte("30"),         //综合测评
		[]byte("90"),         //义工时
	})
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &proofMaterial)
	proofMaterialList = append(proofMaterialList, proofMaterial)
	return proofMaterialList
}

// 测试获取申请材料
func Test_QueryProofMateriallist(t *testing.T) {
	stub := initTest(t)
	proofMaterialList := checkCreateProofMaterial(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryProofMaterialList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryProofMaterialOnly"),
			[]byte(proofMaterialList[0].StipendId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryProofMaterialList"),
			[]byte("0"),
		}).Payload)))
}

// 测试更改申请材料
func Test_UpdateProofMaterial(t *testing.T) {
	stub := initTest(t)
	proofMaterialList := checkCreateProofMaterial(stub, t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("updateProofMaterial"),
		[]byte(proofMaterialList[0].StipendId),
		[]byte("1234"), //收入
		[]byte("48"),   //综测
		[]byte("82"),   //义工时
	})
	////参数个数不满足
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("updateProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("1234"),       //收入
	//	[]byte("48"),         //综测
	//})
	////参数存在空值
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("updateProofMaterial"),
	//	[]byte(""),   //学号
	//	[]byte("0"),  //收入
	//	[]byte("12"), //综测
	//	[]byte("48"), //义工
	//})
	////综合测评超过120
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("updateProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("2712"),       //收入
	//	[]byte("122"),        //综测
	//	[]byte("48"),         //义工
	//})
	////义工时小于20小时不参与评选助学金
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("updateProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("2313"),       //收入
	//	[]byte("115"),        //综测
	//	[]byte("18"),         //义工
	//})
	////家庭人均年收入超过5000不参与评选助学金
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("updateProofMaterial"),
	//	[]byte("0000000000"), //学号
	//	[]byte("7313"),       //收入
	//	[]byte("115"),        //综测
	//	[]byte("29"),         //义工
	//})
	////学生不存在
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createProofMaterial"),
	//	[]byte("0000012345"), //学号
	//	[]byte("2313"),       //收入
	//	[]byte("115"),        //综测
	//	[]byte("29"),         //义工
	//})
}

//手动创建一些申请材料
func checkCreateQueryStipendRanking(stub *shim.MockStub, t *testing.T) []lib.ProofMaterial {
	var rankList []lib.ProofMaterial

	var rank lib.ProofMaterial
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createQueryStipendRanking"),
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &rank)
	rankList = append(rankList, rank)
	return rankList
}

//// 测试排序
//func Test_CreateStipendRanking(t *testing.T) {
//	stub := initTest(t)
//	rankList := checkCreateQueryStipendRanking(stub, t)
//	//成功
//	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n",
//		rankList))
//}

//测试创建评分
func Test_CreateVote(t *testing.T) {
	stub := initTest(t)
	proofMaterialList := checkCreateProofMaterial(stub, t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("createVote"),
		[]byte(proofMaterialList[0].StipendId),
		[]byte("0000000000"), //评分人
		[]byte("4"),          //评分
	})
	////参数存在空值
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createVote"),
	//	[]byte(""),           //被评分人
	//	[]byte("0000000000"), //评分人
	//	[]byte(""),           //评分
	//})
	////参数个数不满足
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createVote"),
	//	[]byte("3118002204"), //被评分人
	//	[]byte("0000000000"), //评分人
	//})
	////评分人与被评分人为同一人
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createVote"),
	//	[]byte("3118002204"), //被评分人
	//	[]byte("3118002204"), //评分人
	//	[]byte("4"),          //评分
	//})
	////学生用户不存在
	//checkInvoke(t, stub, [][]byte{
	//	[]byte("createVote"),
	//	[]byte("3118002201"), //被评分人
	//	[]byte("0000000000"), //评分人
	//	[]byte("4"),          //评分
	//})
}

//手动创建一些评分
func checkCreateVote(stub *shim.MockStub, t *testing.T) []lib.Vote {
	proofMaterialList := checkCreateProofMaterial(stub, t)
	var voteList []lib.Vote
	var vote lib.Vote
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createVote"),
		[]byte(proofMaterialList[0].StipendId),
		[]byte("0000000000"), //评分人
		[]byte("3"),          //评分
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &vote)
	voteList = append(voteList, vote)
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createVote"),
		[]byte("3118002204"), //被评分人
		[]byte("1234567890"), //评分人
		[]byte("5"),          //评分
	})
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &vote)
	voteList = append(voteList, vote)
	resp3 := checkInvoke(t, stub, [][]byte{
		[]byte("createVote"),
		[]byte("0000000000"), //被评分人
		[]byte("3118002204"), //评分人
		[]byte("3.5"),        //评分
	})
	json.Unmarshal(bytes.NewBuffer(resp3.Payload).Bytes(), &vote)
	voteList = append(voteList, vote)
	return voteList
}

// 测试获取申请材料
func Test_QueryVote(t *testing.T) {
	stub := initTest(t)
	voteList := checkCreateVote(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryVote"),
			[]byte(voteList[0].StipendId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryVote"),
			[]byte(voteList[0].StipendId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryVote"),
			[]byte("0"),
		}).Payload)))
}

// 测试获取申请材料
func Test_QueryVoteOnly(t *testing.T) {
	stub := initTest(t)
	voteList := checkCreateVote(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取3118002204数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryVoteOnly"),
			[]byte(voteList[0].StipendId),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取0000000000数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryVoteOnly"),
			[]byte(voteList[2].StipendId),
		}).Payload)))
}
