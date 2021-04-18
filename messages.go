package ddp

import (
	"fmt"
)

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

func MessageIs(m Message, mt MessageType) bool {
	return m.MessageType() == mt
}

var messageTypes = [...]string{"connect", "connected", "failed", "ping", "pong", "sub", "unsub", "nosub", "added", "changed", "removed", "ready", "addedBefore", "movedBefore", "method", "result", "updated"}

func (mt MessageType) String() string {
	return messageTypes[mt]
}

func (mt MessageType) MarshalJSON() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (mt *MessageType) UnmarshalJSON(b []byte) error {
	str := string(b)
	for i, t := range messageTypes {
		if str == t {
			*mt = MessageType(i)
			return nil
		}
	}
	return fmt.Errorf("unknown message type: %q", str)
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

type commonMessage struct {
	Msg MessageType `json:"msg"`
}

func newCommonMessage(mt MessageType) commonMessage {
	cm := commonMessage{
		Msg: mt,
	}
	return cm
}

func (cm commonMessage) MessageType() MessageType {
	return cm.Msg
}

type Connect struct {
	commonMessage
	Session string   `json:"session"`
	Version string   `json:"version"`
	Support []string `json:"support"`
}

func NewConnectMessage(session, version string, support []string) *Connect {
	m := new(Connect)
	m.commonMessage = newCommonMessage(MessageConnect)
	m.Session = session
	m.Version = version
	m.Support = support
	return m
}

type Connected struct {
	commonMessage
	Session string `json:"session"`
}

func NewConnectedMessage(session string) *Connected {
	m := new(Connected)
	m.commonMessage = newCommonMessage(MessageConnected)
	m.Session = session
	return m
}

type Failed struct {
	commonMessage
	Version string `json:"version"`
}

func NewFailedMessage(version string) *Failed {
	m := new(Failed)
	m.commonMessage = newCommonMessage(MessageFailed)
	m.Version = version
	return m
}

type Ping struct {
	commonMessage
	ID string `json:"id"`
}

func NewPingMessage(id string) *Ping {
	m := new(Ping)
	m.commonMessage = newCommonMessage(MessagePing)
	m.ID = id
	return m
}

type Pong struct {
	commonMessage
	ID string `json:"id"`
}

func NewPongMessage(id string) *Pong {
	m := new(Pong)
	m.commonMessage = newCommonMessage(MessagePong)
	m.ID = id
	return m
}

type Sub struct {
	commonMessage
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Params []interface{} `json:"params"`
}

func NewSubMessage(id, name string, params []interface{}) *Sub {
	m := new(Sub)
	m.commonMessage = newCommonMessage(MessageSub)
	m.ID = id
	m.Name = name
	m.Params = params
	return m
}

type UnSub struct {
	commonMessage
	ID string `json:"id"`
}

func NewUnSubMessage(id string) *UnSub {
	m := new(UnSub)
	m.commonMessage = newCommonMessage(MessageUnSub)
	m.ID = id
	return m
}

type NoSub struct {
	commonMessage
	ID    string `json:"id"`
	Error Error  `json:"error"`
}

func NewNoSubMessage(id string, err Error) *NoSub {
	m := new(NoSub)
	m.commonMessage = newCommonMessage(MessageNoSub)
	m.ID = id
	m.Error = err
	return m
}

type Added struct {
	commonMessage
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Fields     map[string]interface{} `json:"fields"`
}

func NewAddedMessage(collection, id string, fields map[string]interface{}) *Added {
	m := new(Added)
	m.commonMessage = newCommonMessage(MessageAdded)
	m.Collection = collection
	m.ID = id
	m.Fields = fields
	return m
}

type Changed struct {
	commonMessage
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Fields     map[string]interface{} `json:"fields"`
	Cleared    []string               `json:"cleared"`
}

func NewChangedMessage(collection, id string, fields map[string]interface{}, cleared []string) *Changed {
	m := new(Changed)
	m.commonMessage = newCommonMessage(MessageChanged)
	m.Collection = collection
	m.ID = id
	m.Fields = fields
	m.Cleared = cleared
	return m
}

type Removed struct {
	commonMessage
	Collection string `json:"collection"`
	ID         string `json:"id"`
}

func NewRemovedMessage(collection, id string) *Removed {
	m := new(Removed)
	m.commonMessage = newCommonMessage(MessageRemoved)
	m.Collection = collection
	m.ID = id
	return m
}

type Ready struct {
	commonMessage
	Subs []string `json:"subs"`
}

func NewReadyMessage(subs []string) *Ready {
	m := new(Ready)
	m.commonMessage = newCommonMessage(MessageReady)
	m.Subs = subs
	return m
}

type AddedBefore struct {
	commonMessage
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Fields     map[string]interface{} `json:"fields"`
	Before     *string                `json:"before"`
}

func NewAddedBeforeMessage(collection, id string, fields map[string]interface{}, before *string) *AddedBefore {
	m := new(AddedBefore)
	m.commonMessage = newCommonMessage(MessageAddedBefore)
	m.Collection = collection
	m.ID = id
	m.Fields = fields
	m.Before = before
	return m
}

type MovedBefore struct {
	commonMessage
	Collection string  `json:"collection"`
	ID         string  `json:"id"`
	Before     *string `json:"before"`
}

func NewMovedBeforeMessage(collection, id string, before *string) *MovedBefore {
	m := new(MovedBefore)
	m.commonMessage = newCommonMessage(MessageMovedBefore)
	m.Collection = collection
	m.ID = id
	m.Before = before
	return m
}

type Method struct {
	commonMessage
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	ID         string        `json:"id"`
	RandomSeed interface{}   `json:"randomSeed"`
}

func NewMethodMessage(method string, params []interface{}, id string, randomSeed interface{}) *Method {
	m := new(Method)
	m.commonMessage = newCommonMessage(MessageMethod)
	m.Method = method
	m.Params = params
	m.ID = id
	m.RandomSeed = randomSeed
	return m
}

type Result struct {
	commonMessage
	ID     string      `json:"id"`
	Error  *Error      `json:"error"`
	Result interface{} `json:"result"`
}

func NewResultMessage(id string, err *Error, result interface{}) *Result {
	m := new(Result)
	m.commonMessage = newCommonMessage(MessageResult)
	m.ID = id
	m.Error = err
	m.Result = result
	return m
}

type Updated struct {
	commonMessage
	Methods []string `json:"methods"`
}

func NewUpdatedMessage(methods []string) *Updated {
	m := new(Updated)
	m.commonMessage = newCommonMessage(MessageUpdated)
	m.Methods = methods
	return m
}
