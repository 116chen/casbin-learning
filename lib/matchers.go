package lib

import (
	"fmt"
	"strings"
)

var admin = []string{"admin", "root"}

func InitFunctions() {
	E.AddFunction("methodMatch", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 2 {
			key1, key2 := arguments[0].(string), arguments[1].(string)
			return MethodMatch(key1, key2), nil
		}
		return nil, fmt.Errorf("methodMatch err")
	})

	E.AddFunction("isSuperAdmin", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 1 {
			key := arguments[0].(string)
			return IsSuperAdmin(key), nil
		}
		return nil, fmt.Errorf("isSuperAdmin err")
	})
}

func MethodMatch(key1, key2 string) bool {
	ks := strings.Split(key2, ",")
	for _, v := range ks {
		if v == key1 {
			return true
		}
	}
	return false
}

func IsSuperAdmin(key string) bool {
	for _, v := range admin {
		if v == key {
			return true
		}
	}
	return false
}
