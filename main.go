package main

import (
	"fmt"

    "db/LSMTree/commitlog"
    "db/LSMTree/engine"
)
func main() {
    // TODO: Implement server for requests to the DB

    var cl commitlog.CommitLog
    cl = commitlog.CommitLog{}
    cl.NewCommitLog()
    e := engine.NewRecord("key", "value")
    fmt.Println(cl.Size)

    for i := 0; i < 10; i++ {
        cl.Append(e)
    }
    fmt.Println(cl.Size)

    list := cl.ReadLog()
    fmt.Println(len(list))
    fmt.Printf("%s %s", list[0].Key, list[0].Value)

    cl.Clear()
    fmt.Println(cl.Size)
}
