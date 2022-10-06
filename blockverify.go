package blockverify

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	Nonce    int
	Trans    string
	Prevhash string
	Hash     string
}

func Newblock(n int, t string) *Block {
	s := new(Block)
	s.Nonce = n
	s.Trans = t
	return s
}

type Blocklist struct {
	list []*Block
}

func ListBlocks(obj *Blocklist) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	for i := 0; i < len(obj.list); i++ {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("Nonce:    ", obj.list[i].Nonce)
		fmt.Println("Transaction ID:   ", obj.list[i].Trans)
		fmt.Println("Previous Hash:   ", obj.list[i].Prevhash)
		fmt.Println("Current Hash:   ", obj.list[i].Hash)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func (ls *Blocklist) AddBlock(n int, t string) *Block {
	st := Newblock(n, t)

	if VerifyChain(ls) {
		ls.list = append(ls.list, st)
		CalculateHash(ls)
		fmt.Println("Block Added")
		return st
	} else {
		return nil
	}
}

func (s *Block) GetString() string {

	var r = ""
	r += strconv.Itoa(s.Nonce)
	r += s.Trans + s.Prevhash
	return r
}

func CalculateHash(stud *Blocklist) {
	for i := 0; i < len(stud.list); i++ {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		stud.list[i].Hash = fmt.Sprintf("%x", sum)
		if i < len(stud.list)-1 {
			stud.list[i+1].Prevhash = fmt.Sprintf("%x", sum)
		}
	}
}

func ChangeBlock(stud *Blocklist, n int, t string) {

	for i := 0; i < len(stud.list); i++ {
		if n == stud.list[i].Nonce {

			stud.list[i].Trans = t
			fmt.Println("Changes Done")
			return
		}
	}
	fmt.Println("Block Not Found!")
}

func VerifyChain(stud *Blocklist) bool {
	var st = ""
	for i := 0; i < len(stud.list); i++ {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		st = fmt.Sprintf("%x", sum)

		if st != stud.list[i].Hash {
			fmt.Printf("Block Tempered, Block #. %d\n", i)
			return false
		}
	}
	fmt.Println("Blocks are Valid")
	return true
}
