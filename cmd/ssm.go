package cmd

import (
	"fmt"
)

func printVersion() error {
	fmt.Println("Version: ", ssmName)
	return nil
}
