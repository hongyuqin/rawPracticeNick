package main

import (
	"core"
	"fmt"
	"strconv"
)

func main() {
	bc := core.NewBlockchain()             //初始化区块链，创建第一个区块
	bc.AddBlock("Send 1 BTC to Ivan")      //加入一个区块，发送一个比特币给ivan
	bc.AddBlock("Send 2 more BTC to Ivan") //加入一个区块，发送更多比特币给ivan

	for _, block := range bc.Blocks {
		fmt.Printf("Prevhash : %x\n", block.PrevBlockHash) //前一个区块的哈希值
		fmt.Printf("Data :%s\n", block.Data)
		fmt.Printf("Hash :%x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("Pow %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println("===========")
	}
	/**
	工作量证明：
	1.工作的结果作为数据加入区块链成为一个区块
	2.完成这个工作的人也会获得奖励(这也就是通过挖矿获得比特币)
	3.整个 努力工作并进行证明 的机制，就叫工作量证明
	*/
}
