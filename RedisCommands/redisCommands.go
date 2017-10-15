package internals

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var connection redis.Conn

//ConnectRedis connects to redis
func ConnectRedis(address string, db string) error {

	connection, err := redis.Dial("tcp", address)
	if err != nil {
		return err
	}
	//Need to handle password
	// _, err = connection.Do("auth ", "")
	// if err != nil {
	// 	return err
	// }
	dbSelected, err := connection.Do("SELECT", db)
	fmt.Println(dbSelected)
	if err != nil {
		return err
	}
	fmt.Println("Connected", connection)
	return nil
}

//SelectAllKeys lists all keys from the specified db
func SelectAllKeys() error {
	if connection == nil {
		return errors.New("Server disconnected. Connect again")
	}
	keys, err := connection.Do("KEYS *")
	if err != nil {
		return err
	}
	fmt.Println(keys)
	return nil
}
