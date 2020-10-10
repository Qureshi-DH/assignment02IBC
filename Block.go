package assignment02IBC

// Block ...
type Block struct {
	Spender     map[string]int
	Receiver    map[string]int
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}
