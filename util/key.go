package util

import (
	"fmt"
	"strconv"
)

var Key = "ScoreRank"

func GetKey(uid uint) string {
	return fmt.Sprintf("user:%s", strconv.Itoa(int(uid)))
}
