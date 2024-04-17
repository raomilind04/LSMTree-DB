package engine 

type dbEngineInterface interface {
    put(key string, value string) (string, bool)
    get(key string) (string, bool)
    delete(key string) bool
}

