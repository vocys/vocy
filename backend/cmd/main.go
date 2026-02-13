package main

import (
	"fmt"
	"os"
	_ "time/tzdata"

	"github.com/vocys/vocy/internal/cmds"
	"github.com/vocys/vocy/internal/common"
)

// @title vocy API
// @version 1.0
// @description.markdown

func main() {
	if err := common.ValidateEnvConfig(&common.EnvConfig); err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(1)
	}
	cmds.Execute()
}
