package ssm

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// CreateECDSAKeyPair creates a new ECDSA key pair.
func CreateECDSAKeyPair() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("Failed to generate ECDSA key pair", "error", err)
		return EmptyHexString, EmptyHexString, EmptyHexString, err
	}
	priKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))
	pubKeyStr := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	compressPubkeyStr := hex.EncodeToString(crypto.CompressPubkey(&privateKey.PublicKey))

	return priKeyStr, pubKeyStr, compressPubkeyStr, nil
}

// SignECDSAMessage signs a message using ECDSA.
func SignECDSAMessage(priKeyStr, txMsg string) (string, error) {
	hash := common.HexToHash(txMsg)
	fmt.Println(hash)
	privByte, err := hex.DecodeString(priKeyStr)
	if err != nil {
		log.Error("Failed to decode private key", "error", err)
		return EmptyHexString, err
	}
	privKeyEcdsa, err := crypto.ToECDSA(privByte)
	if err != nil {
		log.Error("Failed to convert private key to ECDSA", "error", err)
		return EmptyHexString, err
	}
	signatureByte, err := crypto.Sign(hash[:], privKeyEcdsa)
	if err != nil {
		log.Error("Failed to sign ECDSA message", "error", err)
		return EmptyHexString, err
	}

	return hex.EncodeToString(signatureByte), nil
}

// VerifyECDSAMessage verifies an ECDSA signature.
func VerifyECDSAMessage(publicKey, txHash, signature string) (bool, error) {
	publicKeyStr, err := hex.DecodeString(publicKey)
	if err != nil {
		log.Error("Failed to decode public key", "error", err)
		return false, err
	}
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		log.Error("Failed to decode transaction hash", "error", err)
		return false, err
	}

	signBytes, err := hex.DecodeString(signature)

	if err != nil {
		log.Error("Failed to decode signature", "error", err)
		return false, err
	}

	return crypto.VerifySignature(publicKeyStr, txHashBytes, signBytes[:64]), nil
}
