// 代码生成时间: 2025-10-05 02:39:23
package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "fmt"
    "io"
)

// Sign signs the given data with the provided private key.
func Sign(data []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
# 添加错误处理
    // Compute the SHA-256 hash of the data.
    hash := sha256.Sum256(data)

    // Sign the hash with the private key.
    r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
    if err != nil {
        return nil, err
    }

    // Combine the two signature parts into a single byte slice.
    signature := append(r.Bytes(), s.Bytes()...)
    return signature, nil
}

// Verify checks if the given signature is valid for the provided data and public key.
func Verify(data []byte, signature []byte, publicKey *ecdsa.PublicKey) bool {
    // Compute the SHA-256 hash of the data.
    hash := sha256.Sum256(data)
# 添加错误处理

    // Parse the signature into its two components.
# FIXME: 处理边界情况
    r := big.NewInt(0).SetBytes(signature[:(len(signature) + 1) / 2])
    s := big.NewInt(0).SetBytes(signature[(len(signature) + 1) / 2:])

    // Verify the signature with the public key.
    return ecdsa.Verify(publicKey, hash[:], r, s)
}

func main() {
    // Generate a new ECDSA key pair.
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
# TODO: 优化性能
    if err != nil {
        fmt.Printf("Failed to generate private key: %s
", err)
        return
# 添加错误处理
    }
    fmt.Println("Private key generated.")

    // Extract the public key from the private key.
    publicKey := &privateKey.PublicKey
    fmt.Println("Public key extracted.")

    // Example data to sign.
    data := []byte("This is a sample data for digital signature.")

    // Sign the data.
    signature, err := Sign(data, privateKey)
    if err != nil {
        fmt.Printf("Failed to sign data: %s
", err)
# 添加错误处理
        return
    }
# 扩展功能模块
    fmt.Printf("Data signed. Signature: %x
", signature)

    // Verify the signature.
    if Verify(data, signature, publicKey) {
        fmt.Println("Signature is valid.")
    } else {
# NOTE: 重要实现细节
        fmt.Println("Signature is invalid.")
    }
}
# 添加错误处理
