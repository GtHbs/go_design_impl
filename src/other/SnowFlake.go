package other

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func SnowFlakeMain() {

	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println(id)
	}
}
