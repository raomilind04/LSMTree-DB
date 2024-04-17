package commitlog

import (
	"bufio"
	"log"
	"os"

	"db/LSMTree/engine"
)

const FILE_PATH = "log.txt"

type CommitLogInterface interface {
    append(entry engine.Record)
    readLog() []engine.Record
    size() uint64
    clear()
}

type CommitLog struct {
    Filepath string 
    Size     uint64
}

func (cl *CommitLog) initialiseCommitLog() {
    cl.Filepath = FILE_PATH
    cl.Size = 0
}

func (cl *CommitLog) NewCommitLog() {
   cl.initialiseCommitLog()

    _, err := os.Stat(cl.Filepath)
    if !os.IsNotExist(err) {
        log.Printf("Log exists, calculating size \n")
        cl.Size = calculateLogSize(cl.Filepath)
    } else {
        log.Printf("Log doesn't exist, creating new file %s", cl.Filepath)
        createLogFile(cl.Filepath)
    }
}

func calculateLogSize(filepath string) uint64 {
    file, err := os.Open(filepath)
    if err != nil {
        // TODO: This will panic and cause the entire engine to fail, errors should be handled by the engine
        log.Panicf("Unable to open file %s", filepath)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var size uint64 = 0

    for scanner.Scan() {
        size++
    }
    return size
}

func createLogFile(filepath string) {
    file, err := os.Create(filepath)
    
    if err != nil {
        // TODO: This will panic and cause the entire engine to fail, errors should be handled by the engine
        log.Panicf("Unable to create file %s : %v", filepath, err)
    }

    defer file.Close()
}
