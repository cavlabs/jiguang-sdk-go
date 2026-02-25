// Copyright 2025 cavlabs/jiguang-sdk-go authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jiguang_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/v2/jiguang"
	"github.com/cavlabs/jiguang-sdk-go/v2/third_party/gmsm/sm2"
)

const (
	sm2B64PubKey  = "BKyLwHkGHKW0aUlYciVxrLtMfSUTJIqoYn0mooiIZJbled1+d/IO+JaxD/6PE7eoO84Ko/8rGCD0k7+vZn+tuU8="
	sm2B64PrivKey = "q9DebGDlx84PQe/eNUzyyNB4B1s8kVTTsKchH970goc="
	sm2PrivKeyD   = "abd0de6c60e5c7ce0f41efde354cf2c8d078075b3c9154d3b0a7211fdef48287"
	sm2PubKeyX    = "ac8bc079061ca5b4694958722571acbb4c7d2513248aa8627d26a288886496e5"
	sm2PubKeyY    = "79dd7e77f20ef896b10ffe8f13b7a83bce0aa3ff2b1820f493bfaf667fadb94f"
)

// 打印 SM2 公钥的 Base64 格式
func sprintB64PubKey(pubKey *sm2.PublicKey) string {
	curve, x, y := sm2.P256Sm2(), pubKey.X, pubKey.Y

	// Note: `elliptic.Marshal` has been deprecated since Go 1.21.

	// (0, 0) is the point at infinity by convention. It's ok to operate on it,
	// although IsOnCurve is documented to return false for it. See Issue 37294.
	if (x.Sign() != 0 || y.Sign() != 0) && !curve.IsOnCurve(x, y) {
		panic("gmsm2: attempted operation on invalid point")
	}

	byteLen := (curve.Params().BitSize + 7) / 8

	pubBytes := make([]byte, 1+2*byteLen)
	pubBytes[0] = 4 // uncompressed point

	x.FillBytes(pubBytes[1 : 1+byteLen])
	y.FillBytes(pubBytes[1+byteLen : 1+2*byteLen])

	return base64.StdEncoding.EncodeToString(pubBytes)
}

// 将 SM2 私钥打印为可复现格式
func sprintPrivKey(privKey *sm2.PrivateKey) string {
	dHex := hex.EncodeToString(privKey.D.Bytes())
	xHex := hex.EncodeToString(privKey.X.Bytes())
	yHex := hex.EncodeToString(privKey.Y.Bytes())

	return fmt.Sprintf("私钥 (D): %s\n公钥 (X, Y):\nX: %s\nY: %s\n", dHex, xHex, yHex)
}

// nolint:unused
func rebuildSm2PrivKey() (*sm2.PrivateKey, error) {
	// 解码 Base64 私钥
	privKeyBytes, _ := base64.StdEncoding.DecodeString(sm2B64PrivKey)

	// 将私钥字节数组转换为大整数 D
	d := new(big.Int).SetBytes(privKeyBytes)

	curve := sm2.P256Sm2() // 使用 sm2p256v1 曲线

	// 通过 D 计算公钥坐标 X, Y
	x, y := curve.ScalarBaseMult(d.Bytes())

	// 构造 SM2 私钥
	return &sm2.PrivateKey{
		D: d,
		PublicKey: sm2.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
	}, nil
}

// nolint:unused
func rebuildSm2PrivKeyByDXY() (*sm2.PrivateKey, error) {
	// 将十六进制字符串转换为大整数
	d, _ := new(big.Int).SetString(sm2PrivKeyD, 16)
	x, _ := new(big.Int).SetString(sm2PubKeyX, 16)
	y, _ := new(big.Int).SetString(sm2PubKeyY, 16)

	curve := sm2.P256Sm2() // 使用 sm2p256v1 曲线

	// 构造 SM2 私钥
	return &sm2.PrivateKey{
		D: d,
		PublicKey: sm2.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
	}, nil
}

func TestGmSm2(t *testing.T) {
	privKey, _ := rebuildSm2PrivKey()
	// privKey, _ := rebuildSm2PrivKeyByDXY()
	// 序列化私钥（仅序列化 D 值）
	privKeyBytes := privKey.D.Bytes()
	// 将私钥序列化结果转为 Base64
	b64PrivKey := base64.StdEncoding.EncodeToString(privKeyBytes)
	t.Logf("SM2 私钥 (Base64): %s\n%s", b64PrivKey, sprintPrivKey(privKey))
	t.Logf("SM2 公钥 (Base64): %s", sprintB64PubKey(&privKey.PublicKey))

	// -----------------------------------------------------------------------------------------------------------------

	const plainText = "ABCDEFabdef123456!@#$😄emoji表情😂にちほん"
	t.Log("\n==== 加密 ====")
	cipherB64, err := jiguang.EncryptWithSM2([]byte(plainText))
	if err != nil {
		t.Errorf("加密失败: %v\n", err)
		return
	}
	t.Logf("密文 (Base64): %s\n", cipherB64)
	cipherBytes, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		t.Errorf("Base64 解码错误: %v", err)
		return
	}
	cipherHex := hex.EncodeToString(cipherBytes)
	t.Logf("加密结果: %s\n", cipherHex)

	// -----------------------------------------------------------------------------------------------------------------

	// cipherB64 = "BG16SQPntGtstHFJNHERgkuF5eB/scGQc1XyEZ5XpeL7K2EYXNKKPAzYqb5g39wacEdM5Hbpdb5MqSUVKv/ZGp6G8/Ya6q2FRXeJ4zq4osak9XmAiw8uYc1c3K3ShVnDBXYO4B9yMVV8C5or+odL3kt0AfRsyWSLR6ByxODcP5nl9re5GdmllyIqc5CDV8xCU7mUDmUFuI0T7d4jON8Q4w1RFhQd6K67a9Rwza//l2782tZ2oOgO0uBbknnbEvd8rK2OBIr/Z3ZXmcHp9CW18kkwjnvqtipy0g2y/teJ62wmiHPXupUVOld17hjXUU6FQdIfvvzkQeejFbxABBibZhsQpgXHxQimQJ1Nirk++qWqbS4RRmkq8YunxJ5fP8asJ7TnIGWuoij0J/HfuCwwrH++X+ZtL5pAZXJGRIbwq+G7mZuOYW+auRJVAhZ+T7yVrFNf1VqiVL6QLBgp3sUSsCW2hQU9On5z369WSSF0CZCBoJ3AcSFRsLirMf3/N1VyxFB1J8hLM6gvaPbvS+NauFsaaugmtRqwsQufpFacHH+V7bLoryFpdsZGlr8bDoORO94wPIGSwXisVCr++q/TAc7Wxz5DzeN0C/ldo4e4+MTvOBKCx8qoCBGe/CTVZoVpTUP+aFEXd3Vq927NpLCWUrykt26zeOveSuMSAV3cbY5PaEfd0EQLVLDJGsfDeTABGpggOyIWhL8zMijQbSmM0IYh+yM1yDEva2Ecl4FQA3JWKe9vkWCPDqgJbI+Ckjrh1pWn5f+ZnfQRSNKUWXkrS+J4xaRxQTtmUilHNsKckmi07CPNcCyL3Wa3pQQK/BWQsA=="

	t.Log("\n==== 解密 ====")
	plainBytes, err := jiguang.DecryptWithSM2(cipherB64)
	if err != nil {
		t.Errorf("解密失败: %v\n", err)
		return
	}
	t.Logf("解密结果: %s\n", plainBytes)
}
