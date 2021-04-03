package ddp

type Message interface {
	MessageType() MessageType
}

type MessageType int

const (
	MessageConnect MessageType = iota
	MessageConnected
	MessageFailed
	MessagePing
	MessagePong
	MessageSub
	MessageUnSub
	MessageNoSub
	MessageAdded
	MessageChanged
	MessageRemoved
	MessageReady
	MessageAddedBefore
	MessageMovedBefore
	MessageMethod
	MessageResult
	MessageUpdated
)

var messageNames = [...]string{"connect", "connected", "failed", "ping", "pong", "sub", "unSub", "noSub", "added", "changed", "removed", "ready", "addedBefore", "movedBefore", "method", "result", "updated"}

func (mt MessageType) String() string {
	return messageNames[mt]
}

type Error struct {
	Err              string      `json:"error"`
	Reason           string      `json:"reason"`
	Message          string      `json:"message"`
	ErrorType        string      `json:"errorType"`
	OffendingMessage interface{} `json:"offendingMessage"`
}

func (e Error) Error() string {
	// todo: return more?
	return e.Err
}

type Connect struct {
	Session string   `json:"session"`
	Version string   `json:"version"`
	Support []string `json:"support"`
}

func (Connect) MessageType() MessageType { return MessageConnect }

type Connected struct {
	Session string `json:"session"`
}

func (Connected) MessageType() MessageType { return MessageConnected }

type Failed struct {
	Version string `json:"version"`
}

func (Failed) MessageType() MessageType { return MessageFailed }

type Ping struct {
	ID string `json:"id"`
}

func (Ping) MessageType() MessageType { return MessagePing }

type Pong struct {
	ID string `json:"id"`
}

func (Pong) MessageType() MessageType { return MessagePong }

type Sub struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Params []interface{} `json:"params"`
}

func (Sub) MessageType() MessageType { return MessageSub }

type UnSub struct {
	ID string `json:"id"`
}

func (UnSub) MessageType() MessageType { return MessageUnSub }

type NoSub struct {
	ID    string `json:"id"`
	Error Error  `json:"error"`
}

func (NoSub) MessageType() MessageType { return MessageNoSub }

type Added struct {
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Fields     map[string]interface{} `json:"fields"`
}

func (Added) MessageType() MessageType { return MessageAdded }

type Changed struct {
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Fields     map[string]interface{} `json:"fields"`
	Cleared    []string               `json:"cleared"`
}

func (Changed) MessageType() MessageType { return MessageChanged }

type Removed struct {
	Collection string `json:"collection"`
	ID         string `json:"id"`
}

func (Removed) MessageType() MessageType { return MessageRemoved }

type Ready struct {
	Subs []string `json:"subs"`
}

func (Ready) MessageType() MessageType { return MessageReady }

type AddedBefore struct {
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Fields     map[string]interface{} `json:"fields"`
	Before     *string                `json:"before"`
}

func (AddedBefore) MessageType() MessageType { return MessageAddedBefore }

type MovedBefore struct {
	Collection string  `json:"collection"`
	ID         string  `json:"id"`
	Before     *string `json:"before"`
}

func (MovedBefore) MessageType() MessageType { return MessageMovedBefore }

type Method struct {
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	ID         string        `json:"id"`
	RandomSeed interface{}   `json:"randomSeed"`
}

func (Method) MessageType() MessageType { return MessageMethod }

type Result struct {
	ID     string      `json:"id"`
	Error  *Error      `json:"error"`
	Result interface{} `json:"result"`
}

func (Result) MessageType() MessageType { return MessageResult }

type Updated struct {
	Methods []string `json:"methods"`
}

func (Updated) MessageType() MessageType { return MessageUpdated }
