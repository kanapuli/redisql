package cmdparser

import (
	"errors"
	"strings"

	"github.com/Athavankanapuli/redisql/RedisCommands"
)

//ParseCommand parses the os.Stdin commands
func ParseCommand(cmd string) (interface{}, error) {
	//Split the input command into words
	parsedSyntax := strings.Split(cmd, " ")
	if len(parsedSyntax) > 0 {
		switch parsedSyntax[0] {
		//syntax is connect 127.0.0.1:6379 0
		case "connect":
			internals.ConnectRedis(parsedSyntax[1], parsedSyntax[2])
		case "select":
			if parsedSyntax[1] == "*" {
				internals.SelectAllKeys()
			} else {
				return nil, errors.New("Not an Valid Query")
			}
		}
	} else {
		return nil, errors.New("Not an Valid Query")
	}
	return nil, nil
}
