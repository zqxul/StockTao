package util

import (
	"fmt"

	"github.com/sony/sonyflake"
)

var snowFlake *sonyflake.Sonyflake

func init() {
	snowFlake.NewSonyflake(sonyflake.Settings{})
}

// NextID ==> fetch next id
func NextID() (uint64, error) {
	id, err := snowFlake.NextID()
	if err != nil {
		fmt.Errorf("get next id error:%v", err)
		return nil, err
	}
	return id, err
}
