package dao

import (
	"fmt"
	"testing"
)

func TestGetMemberInfo(t *testing.T) {
	info, err := GetMemberInfo("", "", "", 0, 100)
	fmt.Println(info, err)
}