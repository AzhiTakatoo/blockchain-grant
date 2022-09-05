package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/lib"
	"github.com/orangebottle/blockchain-grant/chaincode/blockchain-grant/utils"
	"sort"
)

type Interface interface {
	Len() int           // 数组长度
	Less(i, j int) bool //两个元素的大小比较
	Swap(i, j int)      // 交换两个元素
}

type SortStipendList []lib.ProofMaterial

func (array SortStipendList) Len() int {
	return len(array)
}
func (array SortStipendList) Less(i, j int) bool {
	return array[i].StipendScore > array[j].StipendScore //从小到大， 若为大于号，则从大到小
}
func (array SortStipendList) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

//查询学生排名
func QueryAwardList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var sortList []lib.ProofMaterial
	var gardeList []lib.SortStipend
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProofMaterialKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var sortStipend lib.ProofMaterial
			err := json.Unmarshal(v, &sortStipend)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			sortList = append(sortList, sortStipend)
		}
	}
	//排序
	sort.Sort(SortStipendList(sortList))
	gardeList = utils.StipendGarde(sortList)
	sortStipendListByte, err := json.Marshal(gardeList)
	if err != nil {
		return shim.Error(fmt.Sprintf("sortStipendListByte-序列化出错: %s", err))
	}
	return shim.Success(sortStipendListByte)
}

//QuerySortStipendList 奖学金等级排序
func CreateQueryStipendRanking(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var sortList []lib.ProofMaterial
	var gardeList []lib.SortStipend
	var empty []string
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProofMaterialKey, empty)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var sortStipend lib.ProofMaterial
			err := json.Unmarshal(v, &sortStipend)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			sortList = append(sortList, sortStipend)
		}
	}
	//排序
	sort.Sort(SortStipendList(sortList))
	gardeList = utils.StipendGarde(sortList)
	sortStipendListByte, err := json.Marshal(gardeList)
	if err != nil {
		return shim.Error(fmt.Sprintf("sortStipendListByte-序列化出错: %s", err))
	}
	return shim.Success(sortStipendListByte)
}
