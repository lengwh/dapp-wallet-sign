package hsm

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/api/option"
)

type HsmClient struct {
	Ctx     context.Context
	KeyName string
	Gclient *kms.KeyManagementClient
}

func NewHsmClient(ctx context.Context, keyPath string, keyName string) (*HsmClient, error) {
	apiKey := option.WithCredentialsFile(keyPath)

	client, err := kms.NewKeyManagementClient(ctx, apiKey)
	if err != nil {
		log.Error("NewHsmClient", "NewKeyManagementClient", err)
		return nil, err
	}
	return &HsmClient{
		Ctx:     ctx,
		KeyName: keyName,
		Gclient: client,
	}, nil
}

func (hsm *HsmClient) SignTransaction(hash string) (string, error) {
	hashByte, _ := hex.DecodeString(hash)
	req := &kmspb.AsymmetricSignRequest{
		Name: hsm.KeyName,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: hashByte,
			},
		},
	}
	resp, err := hsm.Gclient.AsymmetricSign(hsm.Ctx, req)
	if err != nil {
		log.Error("HsmClient.SignTransaction", "AsymmetricSign", err)
		return common.Hash{}.String(), err
	}
	return hex.EncodeToString(resp.Signature), nil
}

func (hsm *HsmClient) CreateKeyRing(projectID, locationID, keyRingID, keyID string) (string, error) {
	_, err := hsm.Gclient.CreateKeyRing(hsm.Ctx, &kmspb.CreateKeyRingRequest{
		Parent:    "projects/" + projectID + "/locations/" + locationID,
		KeyRingId: keyRingID,
	})
	if err != nil {
		log.Error("HsmClient.CreateKeyRing", "CreateKeyRing", err)
		return "", err
	}
	return keyRingID, nil

}

func (hsm *HsmClient) CreateKeyPair(projectID, locationID, keyRingID, keyID, method string) (string, error) {
	parent := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, locationID, keyRingID)
	var key *kmspb.CryptoKey
	if method == "ecdsa" {
		key = &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_SIGN,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				Algorithm:       kmspb.CryptoKeyVersion_EC_SIGN_P256_SHA256,
				ProtectionLevel: kmspb.ProtectionLevel_HSM,
			},
		}
	} else {
		key = &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_SIGN,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				Algorithm:       kmspb.CryptoKeyVersion_RSA_SIGN_RAW_PKCS1_4096,
				ProtectionLevel: kmspb.ProtectionLevel_HSM,
			},
		}
	}

	cryptoKey, err := hsm.Gclient.CreateCryptoKey(hsm.Ctx, &kmspb.CreateCryptoKeyRequest{
		Parent:      parent,
		CryptoKeyId: keyID,
		CryptoKey:   key,
	})

	if err != nil {
		log.Error("HsmClient.CreateKeyPair", "CreateCryptoKey", err)
		return "", err
	}
	return cryptoKey.Name, nil
}
