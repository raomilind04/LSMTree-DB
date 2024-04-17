package main

import (
	"db/LSMTree/commitlog"
	"fmt"
)
func main() {
    // TODO: Implement server for requests to the DB

    cl := commitlog.CommitLog{}
    cl.NewCommitLog()
    fmt.Println(cl.Size)
}
