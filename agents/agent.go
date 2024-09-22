package agents

type Agent interface {
    Post(message string) (interface{}, error)
    Reply(threadId string, message string) (interface{}, error)
}

