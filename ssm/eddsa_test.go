package ssm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestCreateEDDSAKeyPair(t *testing.T) {
	privKey, pubKey, err := CreateEDDSAKeyPair()
	if err != nil {
		t.Fatalf("Failed to create EDDSA key pair: %v", err)
	}
	fmt.Println("privKey=", privKey)
	fmt.Println("pubKey=", pubKey)
}

/**
privKey= a125a195d7faa81d6001966910ddd26cf2814ee212ab86803cc05f36a2d54250
pubKey= 152d09f60bd0e053dc1a282452f633d2efe975e22147f7153c113efd1fd06d7ca125a195d7faa81d6001966910ddd26cf2814ee212ab86803cc05f36a2d54250
*/

func TestSignEDDSAMessage(t *testing.T) {
	privKey := "a125a195d7faa81d6001966910ddd26cf2814ee212ab86803cc05f36a2d54250"
	//message := "6877498347a58bf169c716d157a503ca85f5f68720d7986a4bd6a9217ad896ca"
	signature, err := SignEDDSAMessage(privKey, common.Hash{}.String())
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}
	fmt.Println("Signature: ", signature)
}

func TestVerifyEDDSAMessage(t *testing.T) {
	publicKey := "39f523de37c1218d28ca467a6e0ea0aa0a603064ab402983829513a0feca0039"
	txHash := common.Hash{}
	signature := "9594742341865c897b066714f10bc90f4c2afebba986c3fef500d9d69ea8eb6df19a19160c210dcf65becbe68c1885820915565f3a84277efc34027b73c89608"

	ok := VerifyEDDSAMessage(publicKey, txHash.String(), signature)
	fmt.Println("Signature valid:", ok)

}
