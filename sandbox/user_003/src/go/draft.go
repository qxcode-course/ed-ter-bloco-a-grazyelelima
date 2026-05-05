package main

import (
    "errors"
    "fmt"
)

type ContaBancaria struct {
    titular string
    saldo   float64
}

// Método com receiver pointer (c *ContaBancaria)
// Retorna apenas error, pois o saldo é alterado diretamente na struct
func (c *ContaBancaria) Depositar(valor float64) error {
    if valor <= 0 {
        return errors.New("valor de depósito inválido")
    }
    c.saldo += valor
    return nil
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor <= 0 {
        return errors.New("valor de saque inválido")
    }
    if valor > c.saldo {
        return errors.New("saldo insuficiente")
    }
    c.saldo -= valor
    return nil
}

func main() {
    conta := ContaBancaria{"Grazy", 300.0}

    // Operação 1: depósito inválido
    err := conta.Depositar(-20)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("Saldo atual: R$ %.2f\n", conta.saldo)
    }

    // Operação 2: depósito válido
    err = conta.Depositar(100)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("Saldo atual: R$ %.2f\n", conta.saldo)
    }

    // Operação 3: saque inválido (maior que saldo)
    err = conta.Sacar(500)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("Saldo atual: R$ %.2f\n", conta.saldo)
    }

    // Operação 4: saque válido
    err = conta.Sacar(50)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Printf("Saldo atual: R$ %.2f\n", conta.saldo)
    }
}