package utils

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func GenSnowflake() int64 {

	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		fmt.Println("generate snow id failed")
	}
	return int64(id)
}
