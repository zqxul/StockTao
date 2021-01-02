package util

import (
	"fmt"

	"github.com/sony/sonyflake"
)

var snowFlake *sonyflake.Sonyflake

func init() {
	snowFlake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// NextID ==> fetch next id
func NextID() (uint64, error) {
	id, err := snowFlake.NextID()
	if err != nil {
		fmt.Printf("get next id error:%v\n", err)
		return 0, err
	}
	return id, err
}
