package commitlog

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
        log.Printf("Log file exists, calculating size \n")
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

func (cl *CommitLog) Append(newRecord engine.Record) {
    file, err := os.OpenFile(cl.Filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        log.Panicf("Unable to open the file %s :  %v", cl.Filepath, err); 
    }
    defer file.Close()

    commitlogEntry := fmt.Sprintf("%s::%s\n", newRecord.Key, newRecord.Value)
    if _, err := file.WriteString(commitlogEntry); err != nil {
        log.Panicf("Unable to append to file %s : %v", cl.Filepath, err)
    }
    cl.Size++
}

func createRecordFromFile(line string) engine.Record {
    data := strings.Split(line, "::")
    return engine.NewRecord(data[0], data[1])
}

func (cl *CommitLog) ReadLog() []engine.Record {
    file, err := os.Open(cl.Filepath)
    if err != nil {
        log.Panicf("Unable to open the file %s :  %v", cl.Filepath, err); 
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    var commitlogRecords []engine.Record

    for fileScanner.Scan() {
        commitlogRecord := createRecordFromFile(fileScanner.Text())
        commitlogRecords = append(commitlogRecords, commitlogRecord)
    }

    return commitlogRecords
}

func (cl *CommitLog) GetSize() uint64 {
    return cl.Size
}

func (cl *CommitLog) Clear() {
    if err := os.Remove(cl.Filepath); err != nil {
        log.Panicf("Unable to open the file %s :  %v", cl.Filepath, err); 
    }
    createLogFile(cl.Filepath)
    cl.Size = 0
}
