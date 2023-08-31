package mutex

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int = 0 

    withdrawls = []transaction {
        {"withdraw", 500},
        {"withdraw", 200},
    }

    deposits = []transaction {
        {"deposit", 500},
        {"deposit", 200},
    }

)

type transaction struct {
    transactionType string
    amount int 
}

func MutexMain() {
    
    var wg sync.WaitGroup

    for _, trx := range deposits {
        wg.Add(1)
        go func(amount int) {
            defer wg.Done()
            for i := 0; i < 1000; i++ {
                deposit(amount)
            }
        }(trx.amount)
    }

    for _, trx := range withdrawls {
        wg.Add(1)
        go func(amount int) {
            defer wg.Done()
            for i := 0; i < 1000; i++ {
                withdraw(amount)
            }
        }(trx.amount)
    }

    wg.Wait()

    fmt.Println("----------------------------------------------------------")
    fmt.Printf("New Balance %d\n", balance)
}

func deposit(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	balance += value
	fmt.Printf("[Deposit]\t | %d\t | to account\t | with balance %d\n", value, balance)
}

func withdraw(value int) {
	mutex.Lock()
    defer mutex.Unlock()
	balance -= value
	fmt.Printf("[Withdrawing]\t | %d\t | from account\t | with balance %d\n", value, balance)
}
