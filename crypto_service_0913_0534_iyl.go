// 代码生成时间: 2025-09-13 05:34:58
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
    "log"
    "net"
)

// CryptoService is the server API for the crypto service.
type CryptoService struct {
    // Contains filtered or unexported fields.
}

// EncryptPassword encrypts the given password using AES-256-GCM algorithm.
func (s *CryptoService) EncryptPassword(ctx net.Context, req *EncryptRequest) (*EncryptResponse, error) {
    key := []byte("your-256-bit-key-here") // Replace with your actual key
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, fmt.Errorf("failed to create cipher: %w", err)
    }
    
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return nil, fmt.Errorf("failed to create GCM: %w", err)
    }
    
    nonce := make([]byte, aesGCM.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, fmt.Errorf("failed to generate nonce: %w", err)
    }
    
    encrypted := aesGCM.Seal(nonce, nonce, req.Password, nil)
    encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
    
    return &EncryptResponse{Password: encryptedBase64}, nil
}

// DecryptPassword decrypts the given encrypted password using AES-256-GCM algorithm.
func (s *CryptoService) DecryptPassword(ctx net.Context, req *DecryptRequest) (*DecryptResponse, error) {
    key := []byte("your-256-bit-key-here") // Replace with your actual key
    encryptedBase64 := req.Password
    encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
    if err != nil {
        return nil, fmt.Errorf("failed to base64 decode: %w", err)
    }
    
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, fmt.Errorf("failed to create cipher: %w", err)
    }
    
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return nil, fmt.Errorf("failed to create GCM: %w", err)
    }
    
    nonceSize := aesGCM.NonceSize()
    if len(encrypted) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    
    nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
    password, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to decrypt: %w", err)
    }
    
    return &DecryptResponse{Password: string(password)}, nil
}

// EncryptRequest is the request message for the EncryptPassword method.
type EncryptRequest struct {
    Password []byte `protobuf:"varint,1,opt,name=password,proto3" json:"password,omitempty"`
}

// EncryptResponse is the response message for the EncryptPassword method.
type EncryptResponse struct {
    Password string `protobuf:"string,1,opt,name=password,proto3" json:"password,omitempty"`
}

// DecryptRequest is the request message for the DecryptPassword method.
type DecryptRequest struct {
    Password string `protobuf:"string,1,opt,name=password,proto3" json:"password,omitempty"`
}

// DecryptResponse is the response message for the DecryptPassword method.
type DecryptResponse struct {
    Password string `protobuf:"string,1,opt,name=password,proto3" json:"password,omitempty"`
}

func main() {
    // Add your server setup and start logic here
    // For example, setup gRPC server, define the service, and start listening on a port
}
