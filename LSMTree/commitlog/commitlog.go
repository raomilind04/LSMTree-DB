package commitlog

import (
    "db/LSMTree/engine"
)

type CommitLog interface {
    append(entry engine.Record)
    readLog() []engine.Record
    size() uint64
    clear()
}
