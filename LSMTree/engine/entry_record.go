package engine 

type Record struct {
    Key   string
    Value string
}

func NewRecord(key string, value string) Record {
    return Record{Key: key, Value: value}
}
