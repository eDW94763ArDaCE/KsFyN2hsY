// 代码生成时间: 2025-09-10 15:49:49
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "io"
    "log"
    "os"
)

// CryptoService defines the methods for password encryption and decryption.
type CryptoService struct{}

// EncryptPassword encrypts a password using AES-256-GCM.
func (s *CryptoService) EncryptPassword(password string) (string, error) {
    key := []byte("your-secret-key") // Replace with a secure key
    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    encrypted := gcm.Seal(nil, nonce, []byte(password), nil)
    return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptPassword decrypts a password using AES-256-GCM.
func (s *CryptoService) DecryptPassword(encryptedPassword string) (string, error) {
    key := []byte("your-secret-key") // Replace with a secure key
    encrypted, err := base64.StdEncoding.DecodeString(encryptedPassword)
    if err != nil {
        return "", err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonceSize := gcm.NonceSize()
    if len(encrypted) < nonceSize {
        return "", errors.New("invalid encrypted payload")
    }
    nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}

func main() {
    service := CryptoService{}
    // Example usage of EncryptPassword and DecryptPassword.
    password := "mysecretpassword"
    encrypted, err := service.EncryptPassword(password)
    if err != nil {
        log.Fatalf("Error encrypting password: %v", err)
    }
    fmt.Println("Encrypted: ", encrypted)

    decrypted, err := service.DecryptPassword(encrypted)
    if err != nil {
        log.Fatalf("Error decrypting password: %v", err)
    }
    fmt.Println("Decrypted: ", decrypted)
}
