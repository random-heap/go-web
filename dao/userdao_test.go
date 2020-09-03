package dao

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	err := SaveUser("admin", "123456", "admin@126.com")
	if err != nil {
		fmt.Println(err)
	}
}
