/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package jiguang

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	"github.com/calvinit/jiguang-sdk-go/third_party/gmsm/sm2"
)

const (
	sm2B64PubKey  = "BPj6Mj/T444gxPaHc6CDCizMRp4pEl14WI2lvIbdEK2c+5XiSqmQt2TQc8hMMZqfxcDqUNQW95puAfQx1asv3rU="
	sm2B64PrivKey = "EjRWeJA=" // N/A: 在客户端 SDK 中不会使用私钥，因此不提供私钥，私钥应该由极光服务端保存。
)

var (
	sm2PubKey  *sm2.PublicKey
	sm2PrivKey *sm2.PrivateKey
)

func init() {
	var err error

	// 初始化 SM2 公钥
	sm2PubKey, err = desm2PubKey(sm2B64PubKey)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize SM2 public key: %v", err))
	}

	// 初始化 SM2 私钥
	sm2PrivKey, err = desm2PrivKey(sm2B64PrivKey)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize SM2 private key: %v", err))
	}
}

// 从 Base64 编码解析 SM2 公钥。
func desm2PubKey(b64PubKey string) (*sm2.PublicKey, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(b64PubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64 SM2 public key: %w", err)
	}

	// Note: `elliptic.Unmarshal` has been deprecated since Go 1.21.

	// 使用 sm2p256v1 曲线
	curve := sm2.P256Sm2()

	// 检查公钥长度
	byteLen := (curve.Params().BitSize + 7) / 8
	if len(pubKeyBytes) != 1+2*byteLen {
		return nil, fmt.Errorf("invalid SM2 public key length: expected %d, got %d", 1+2*byteLen, len(pubKeyBytes))
	}
	// 检查公钥格式
	if pubKeyBytes[0] != 4 { // uncompressed form
		return nil, fmt.Errorf("invalid SM2 public key format: expected uncompressed form (0x04), got 0x%02x", pubKeyBytes[0])
	}

	// 解析公钥坐标
	p := curve.Params().P
	x := new(big.Int).SetBytes(pubKeyBytes[1 : 1+byteLen])
	y := new(big.Int).SetBytes(pubKeyBytes[1+byteLen:])
	// 检查坐标范围
	if x.Cmp(p) >= 0 || y.Cmp(p) >= 0 {
		return nil, fmt.Errorf("invalid SM2 public key coordinates: x or y out of range")
	}
	// 检查坐标是否在曲线上
	if !curve.IsOnCurve(x, y) {
		return nil, fmt.Errorf("invalid SM2 public key coordinates: point not on curve")
	}

	// 返回解析后的公钥
	return &sm2.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}, nil
}

// 从 Base64 编码解析 SM2 私钥。
func desm2PrivKey(b64PrivKey string) (*sm2.PrivateKey, error) {
	privKeyBytes, err := base64.StdEncoding.DecodeString(b64PrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64 SM2 private key: %w", err)
	}

	// 使用 sm2p256v1 曲线
	curve := sm2.P256Sm2()
	privKey := new(sm2.PrivateKey)
	privKey.D = new(big.Int).SetBytes(privKeyBytes)                               // 将字节数组转换为大整数
	privKey.PublicKey.Curve = curve                                               // 设置公钥曲线
	privKey.PublicKey.X, privKey.PublicKey.Y = curve.ScalarBaseMult(privKeyBytes) // 计算公钥坐标

	return privKey, nil
}

// 使用 SM2 公钥加密数据并返回 Base64 编码字符串。
func EncryptWithSM2(data []byte) (string, error) {
	if sm2PubKey == nil {
		return "", errors.New("SM2 public key not initialized")
	}

	cipherBytes, err := sm2.Encrypt(sm2PubKey, data, rand.Reader, sm2.C1C2C3)
	if err != nil {
		return "", fmt.Errorf("SM2 encryption error: %w", err)
	}

	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// 使用 SM2 私钥解密 Base64 编码的加密数据。
func DecryptWithSM2(data string) ([]byte, error) {
	if sm2PrivKey == nil {
		return nil, errors.New("SM2 private key not initialized")
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	plainBytes, err := sm2.Decrypt(sm2PrivKey, cipherBytes, sm2.C1C2C3)
	if err != nil {
		return nil, fmt.Errorf("SM2 decryption error: %w", err)
	}

	return plainBytes, nil
}
