package protobuf

import "errors"

type CryptoType string

const (
	ECDSA CryptoType = "ecdsa"
	EDDSA CryptoType = "eddsa"
)

func ParseTransactionType(s string) (CryptoType, error) {
	switch s {
	case string(ECDSA):
		return ECDSA, nil
	case string(EDDSA):
		return EDDSA, nil
	default:
		return "", errors.New("unknown transaction type")
	}
}
