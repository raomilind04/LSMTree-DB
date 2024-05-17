package memtable

import (
    "github.com/emirpasic/gods/maps/treemap"
    "db/LSMTree/engine" 
)

type Memtable struct {
    data *treemap.Map
}

func NewMemtable() Memtable {
    mt := Memtable{}
    mt.data = treemap.NewWithStringComparator()

    return mt
}

func NewMemtableWithRecords(records []engine.Record) Memtable {
    mt := Memtable{}
    mt.data = treemap.NewWithStringComparator()

    for _,record := range records {
        (mt.data).Put(record.Key, record.Value)
    }

    return mt
}

func (mt Memtable) Put(record engine.Record) {
    (mt.data).Put(record.Key, record.Value)
}

func (mt Memtable) Get(key string) {
    (mt.data).Get(key)
}

func (mt Memtable) GetSize() int {
    return (mt.data).Size()
}

func (mt Memtable) Clear() {
    (mt.data).Clear()
}

func (mt *Memtable) GetMap() *treemap.Map{
    return mt.data
}
