package main_test

import (
	"fmt"
	"github.com/orangebottle/blockchain-grant/application/blockchain"
	"testing"
)

func TestInvoke_QueryRegisterCertify(t *testing.T) {
	blockchain.Init()
	response, e := blockchain.ChannelQuery("queryRegisterCertify", [][]byte{})
	if e != nil {
		fmt.Println(e.Error())
		t.FailNow()
	}
	fmt.Println(string(response.Payload))
}
