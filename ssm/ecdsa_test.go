package ssm

import (
	"fmt"
	"testing"
)

func TestCreateECDSAKeyPair(t *testing.T) {
	priKey, pubKey, cPubKey, _ := CreateECDSAKeyPair()
	fmt.Println("priKey:", priKey)
	fmt.Println("pubKey:", pubKey)
	fmt.Println("cPubKey:", cPubKey)
}

/*
*
priKey: a0e69c67d4e3d9497787cd3991a4c80de6b349fd57ada8c6f864668f37b7aecc
pubKey: 04f35819bb3af4044fdb3fe9d442e2d6d2ae10b8ab8aa990cf2778dc5bca4dc6d5c56cc2f22b4d9a386e9b54e657784e2e17d4640583b25cd43923eb8a347498b2
cPubKey: 02f35819bb3af4044fdb3fe9d442e2d6d2ae10b8ab8aa990cf2778dc5bca4dc6d5
txHash:0x6877498347a58bf169c716d157a503ca85f5f68720d7986a4bd6a9217ad896ca
*/
func TestSignECDSAMessage(t *testing.T) {
	priKey := "a0e69c67d4e3d9497787cd3991a4c80de6b349fd57ada8c6f864668f37b7aecc"
	//hash := sha256.Sum256([]byte("Hello, ECDSA!"))
	message := "0x6877498347a58bf169c716d157a503ca85f5f68720d7986a4bd6a9217ad896ca"
	signature, err := SignECDSAMessage(priKey, message)
	if err != nil {
		t.Errorf("SignECDSAMessage failed: %v", err)
		return
	}
	fmt.Println("signature:", signature)
}

func TestVerifyECDSAMessage(t *testing.T) {
	CompressedPubKey := "02f35819bb3af4044fdb3fe9d442e2d6d2ae10b8ab8aa990cf2778dc5bca4dc6d5"
	//hash := sha256.Sum256([]byte("Hello, ECDSA!"))
	txHash := "6877498347a58bf169c716d157a503ca85f5f68720d7986a4bd6a9217ad896ca"
	signature := "5c972563118e849ad3ba8ec9e2131fc7b9f641af20d6a289df87f5563ad95c8941c901b3bf1fd0d1e1aeb1893af96050c45faabe07815a40942fa789653bf09f00"

	isValid, err := VerifyECDSAMessage(CompressedPubKey, txHash, signature)
	if err != nil {
		t.Errorf("VerifyECDSAMessage failed: %v", err)
		return
	}

	if !isValid {
		t.Error("Signature is invalid")
	}
}
