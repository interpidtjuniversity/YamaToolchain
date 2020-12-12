package main

import (
	"../structural"
)

func main() {
	// opfunc.InitDB()
	commit := structural.Commit{}
	structural.ComputeSha256(&commit)
}
