package main

import (
	"fmt"

    "db/LSMTree/commitlog"
)
func main() {
    // TODO: Implement server for requests to the DB

    cl := commitlog.CommitLog{}
    cl.NewCommitLog()
    cl.app
    fmt.Println(cl.Size)
}
