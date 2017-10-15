package main

import (
	"bufio"
	"fmt"
	"os"

	cmdparser "github.com/Athavankanapuli/redisql/CmdParser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "q" || cmd == "quit" {
			os.Exit(0)
		}
		redisReply, err := cmdparser.ParseCommand(cmd)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(redisReply)
		}

	}
}
