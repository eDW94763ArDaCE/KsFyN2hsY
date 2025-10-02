// 代码生成时间: 2025-10-03 03:23:23
package main

import (
    "crypto/rand"
    "fmt"
    "log"
    "math/big"
    "time"
)

// CryptoWallet represents a cryptocurrency wallet
type CryptoWallet struct {
    // Balance stores the balance of the wallet
    Balance *big.Int
    // WalletID uniquely identifies the wallet
    WalletID string
    // Timestamp of the last transaction
    LastTransactionTime time.Time
}

// NewCryptoWallet creates a new cryptocurrency wallet with a unique ID and initial balance
func NewCryptoWallet() *CryptoWallet {
    id := fmt.Sprintf("wallet_%d", time.Now().UnixNano())
    return &CryptoWallet{
        Balance: big.NewInt(0),
        WalletID: id,
        LastTransactionTime: time.Now(),
    }
}

// Deposit adds funds to the wallet
func (w *CryptoWallet) Deposit(amount *big.Int) error {
    if amount.Cmp(big.NewInt(0)) <= 0 {
        return fmt.Errorf("invalid amount: %s", amount.String())
    }
    w.Balance.Add(w.Balance, amount)
    w.LastTransactionTime = time.Now()
    return nil
}

// Withdraw removes funds from the wallet
func (w *CryptoWallet) Withdraw(amount *big.Int) error {
    if amount.Cmp(big.NewInt(0)) <= 0 {
        return fmt.Errorf("invalid amount: %s", amount.String())
    }
    if w.Balance.Cmp(amount) < 0 {
        return fmt.Errorf("insufficient funds: %s", w.Balance.String())
    }
    w.Balance.Sub(w.Balance, amount)
    w.LastTransactionTime = time.Now()
    return nil
}

// GetBalance returns the current balance of the wallet
func (w *CryptoWallet) GetBalance() *big.Int {
    return new(big.Int).Set(w.Balance)
}

// GetLastTransactionTime returns the timestamp of the last transaction
func (w *CryptoWallet) GetLastTransactionTime() time.Time {
    return w.LastTransactionTime
}

func main() {
    // Create a new wallet
    wallet := NewCryptoWallet()
    fmt.Printf("New wallet created with ID: %s
", wallet.WalletID)

    // Deposit some funds
    depositAmount := big.NewInt(1000)
    if err := wallet.Deposit(depositAmount); err != nil {
        log.Fatalf("Failed to deposit funds: %s
", err)
    }
    fmt.Printf("Deposited %s, new balance: %s
", depositAmount.String(), wallet.GetBalance().String())

    // Withdraw some funds
    withdrawAmount := big.NewInt(500)
    if err := wallet.Withdraw(withdrawAmount); err != nil {
        log.Fatalf("Failed to withdraw funds: %s
", err)
    }
    fmt.Printf("Withdrew %s, new balance: %s
", withdrawAmount.String(), wallet.GetBalance().String())
}
