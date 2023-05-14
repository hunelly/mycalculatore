/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
"os"

_ "embed"
"github.com/goreleaser/goreleaser/cmd"
"super_calculator/cmd"
)

func main() {
	cmd.Execute()
}

