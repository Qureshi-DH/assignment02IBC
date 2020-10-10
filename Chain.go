package assignment02IBC

import (
	crypto "crypto/sha256"
	format "fmt"
)

// CalculateBalance ...
func CalculateBalance(userName string, chainHead *Block) int {
	currentAddress, balance := chainHead, 0
	for currentAddress != nil {

		balance += currentAddress.Receiver[userName]
		balance -= currentAddress.Spender[userName]

		currentAddress = currentAddress.PrevPointer
	}
	return balance
}

// CalculateHash ...
func CalculateHash(inputBlock *Block) string {
	if inputBlock == nil {
		return ""
	}

	serialized := ""

	serialized += format.Sprint(inputBlock.PrevHash)
	serialized += format.Sprint(inputBlock.Receiver)
	serialized += format.Sprint(inputBlock.Spender)

	return format.Sprintf("%x", crypto.Sum256([]byte(serialized)))
}

// InsertBlock ...
func InsertBlock(spendingUser string, receivingUser string, miner string, amount int, chainHead *Block) *Block {
	if miner != rootUser {
		format.Println("Error 4532: Only " + rootUser + " can mine blocks")
		return chainHead
	}

	block := Block{}

	if chainHead != nil {
		block.PrevHash = chainHead.CurrentHash
		block.PrevPointer = chainHead
	} else {
		block.PrevHash = ""
		block.PrevPointer = nil
	}

	block.Spender[spendingUser] -= amount
	block.Receiver[receivingUser] += amount

	block.CurrentHash = CalculateHash(&block)

	return &block
}

// ListBlocks ...
func ListBlocks(chainHead *Block) {
	currentAddress := chainHead
	for {
		format.Print(*currentAddress)
		if currentAddress.PrevPointer != nil {
			format.Print(" --> ")
			currentAddress = currentAddress.PrevPointer
		} else {
			format.Println()
			break
		}
	}
}

// VerifyChain ...
func VerifyChain(chainHead *Block) {
	currentAddress, compromised := chainHead, false

	for currentAddress != nil {
		if CalculateHash(currentAddress.PrevPointer) != currentAddress.PrevHash {
			compromised = true
			break
		}
		currentAddress = currentAddress.PrevPointer
	}
	if !compromised {
		format.Print("Chain Valid")
	} else {
		format.Print("Chain Compromised")
	}
}
