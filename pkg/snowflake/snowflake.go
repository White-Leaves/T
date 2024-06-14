package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func GenInt() int64 {
	node, _ = snowflake.NewNode(1)
	return node.Generate().Int64()
}

func GenString() string {
	node, _ = snowflake.NewNode(1)
	id := node.Generate()
	return id.String()
}
