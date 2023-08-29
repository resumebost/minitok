package rpc

import "minitok/cmd/user/pkg/snowflake"

func InitForUser() {
	snowflake.Init("2023-08-25", 1)
}
