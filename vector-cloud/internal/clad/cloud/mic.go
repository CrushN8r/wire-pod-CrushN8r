// Autogenerated Go message buffer code.
// Source: clad/cloud/mic.clad
// Full command line: victor-clad/tools/message-buffers/emitters/Go_emitter.py -C src -o generated/cladgo/src clad/cloud/mic.clad

package cloud

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/digital-dream-labs/vector-cloud/internal/clad"
)

// ENUM StreamType
type StreamType uint8

const (
	StreamType_Normal StreamType = iota
	StreamType_Blackjack
	StreamType_KnowledgeGraph
)

// ENUM ErrorType
type ErrorType uint8

const (
	ErrorType_Server ErrorType = iota
	ErrorType_Timeout
	ErrorType_Json
	ErrorType_InvalidConfig
	ErrorType_Connecting
	ErrorType_NewStream
	ErrorType_Token
	ErrorType_TLS
	ErrorType_Connectivity
)

// ENUM ConnectionCode
type ConnectionCode uint8

const (
	ConnectionCode_Available ConnectionCode = iota
	ConnectionCode_Connectivity
	ConnectionCode_Tls
	ConnectionCode_Auth
	ConnectionCode_Bandwidth
)

// STRUCTURE StreamOpen
type StreamOpen struct {
	Session string
}

func (s *StreamOpen) Size() uint32 {
	var result uint32
	result += 1                      // Session length (uint_8)
	result += uint32(len(s.Session)) // uint_8 array
	return result
}

func (s *StreamOpen) Unpack(buf *bytes.Buffer) error {
	var SessionLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &SessionLen); err != nil {
		return err
	}
	s.Session = string(buf.Next(int(SessionLen)))
	if len(s.Session) != int(SessionLen) {
		return errors.New("string byte mismatch")
	}
	return nil
}

func (s *StreamOpen) Pack(buf *bytes.Buffer) error {
	if len(s.Session) > 255 {
		return errors.New("max_length overflow in field Session")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(s.Session))); err != nil {
		return err
	}
	if _, err := buf.WriteString(s.Session); err != nil {
		return err
	}
	return nil
}

func (s *StreamOpen) String() string {
	return fmt.Sprint("Session: {", s.Session, "}")
}

// STRUCTURE Hotword
type Hotword struct {
	Mode      StreamType
	Locale    string
	Timezone  string
	NoLogging bool
}

func (h *Hotword) Size() uint32 {
	var result uint32
	result += 1                       // Mode StreamType
	result += 1                       // Locale length (uint_8)
	result += uint32(len(h.Locale))   // uint_8 array
	result += 1                       // Timezone length (uint_8)
	result += uint32(len(h.Timezone)) // uint_8 array
	result += 1                       // NoLogging bool
	return result
}

func (h *Hotword) Unpack(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.LittleEndian, &h.Mode); err != nil {
		return err
	}
	var LocaleLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &LocaleLen); err != nil {
		return err
	}
	h.Locale = string(buf.Next(int(LocaleLen)))
	if len(h.Locale) != int(LocaleLen) {
		return errors.New("string byte mismatch")
	}
	var TimezoneLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &TimezoneLen); err != nil {
		return err
	}
	h.Timezone = string(buf.Next(int(TimezoneLen)))
	if len(h.Timezone) != int(TimezoneLen) {
		return errors.New("string byte mismatch")
	}
	if err := binary.Read(buf, binary.LittleEndian, &h.NoLogging); err != nil {
		return err
	}
	return nil
}

func (h *Hotword) Pack(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.LittleEndian, h.Mode); err != nil {
		return err
	}
	if len(h.Locale) > 255 {
		return errors.New("max_length overflow in field Locale")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(h.Locale))); err != nil {
		return err
	}
	if _, err := buf.WriteString(h.Locale); err != nil {
		return err
	}
	if len(h.Timezone) > 255 {
		return errors.New("max_length overflow in field Timezone")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(h.Timezone))); err != nil {
		return err
	}
	if _, err := buf.WriteString(h.Timezone); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, h.NoLogging); err != nil {
		return err
	}
	return nil
}

func (h *Hotword) String() string {
	return fmt.Sprint("Mode: {", h.Mode, "} ",
		"Locale: {", h.Locale, "} ",
		"Timezone: {", h.Timezone, "} ",
		"NoLogging: {", h.NoLogging, "}")
}

// STRUCTURE Filename
type Filename struct {
	File string
}

func (f *Filename) Size() uint32 {
	var result uint32
	result += 1                   // File length (uint_8)
	result += uint32(len(f.File)) // uint_8 array
	return result
}

func (f *Filename) Unpack(buf *bytes.Buffer) error {
	var FileLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &FileLen); err != nil {
		return err
	}
	f.File = string(buf.Next(int(FileLen)))
	if len(f.File) != int(FileLen) {
		return errors.New("string byte mismatch")
	}
	return nil
}

func (f *Filename) Pack(buf *bytes.Buffer) error {
	if len(f.File) > 255 {
		return errors.New("max_length overflow in field File")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(f.File))); err != nil {
		return err
	}
	if _, err := buf.WriteString(f.File); err != nil {
		return err
	}
	return nil
}

func (f *Filename) String() string {
	return fmt.Sprint("File: {", f.File, "}")
}

// STRUCTURE AudioData
type AudioData struct {
	Data []int16
}

func (a *AudioData) Size() uint32 {
	var result uint32
	result += 2                       // Data length (uint_16)
	result += uint32(len(a.Data)) * 2 // int_16 array
	return result
}

func (a *AudioData) Unpack(buf *bytes.Buffer) error {
	var DataLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &DataLen); err != nil {
		return err
	}
	a.Data = make([]int16, DataLen)
	if err := binary.Read(buf, binary.LittleEndian, &a.Data); err != nil {
		return err
	}
	return nil
}

func (a *AudioData) Pack(buf *bytes.Buffer) error {
	if len(a.Data) > 65535 {
		return errors.New("max_length overflow in field Data")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(a.Data))); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, a.Data); err != nil {
		return err
	}
	return nil
}

func (a *AudioData) String() string {
	return fmt.Sprint("Data: {", a.Data, "}")
}

// STRUCTURE IntentResult
type IntentResult struct {
	Intent     string
	Parameters string
	Metadata   string
}

func (i *IntentResult) Size() uint32 {
	var result uint32
	result += 1                         // Intent length (uint_8)
	result += uint32(len(i.Intent))     // uint_8 array
	result += 2                         // Parameters length (uint_16)
	result += uint32(len(i.Parameters)) // uint_8 array
	result += 2                         // Metadata length (uint_16)
	result += uint32(len(i.Metadata))   // uint_8 array
	return result
}

func (i *IntentResult) Unpack(buf *bytes.Buffer) error {
	var IntentLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &IntentLen); err != nil {
		return err
	}
	i.Intent = string(buf.Next(int(IntentLen)))
	if len(i.Intent) != int(IntentLen) {
		return errors.New("string byte mismatch")
	}
	var ParametersLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &ParametersLen); err != nil {
		return err
	}
	i.Parameters = string(buf.Next(int(ParametersLen)))
	if len(i.Parameters) != int(ParametersLen) {
		return errors.New("string byte mismatch")
	}
	var MetadataLen uint16
	if err := binary.Read(buf, binary.LittleEndian, &MetadataLen); err != nil {
		return err
	}
	i.Metadata = string(buf.Next(int(MetadataLen)))
	if len(i.Metadata) != int(MetadataLen) {
		return errors.New("string byte mismatch")
	}
	return nil
}

func (i *IntentResult) Pack(buf *bytes.Buffer) error {
	if len(i.Intent) > 255 {
		return errors.New("max_length overflow in field Intent")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(i.Intent))); err != nil {
		return err
	}
	if _, err := buf.WriteString(i.Intent); err != nil {
		return err
	}
	if len(i.Parameters) > 65535 {
		return errors.New("max_length overflow in field Parameters")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(i.Parameters))); err != nil {
		return err
	}
	if _, err := buf.WriteString(i.Parameters); err != nil {
		return err
	}
	if len(i.Metadata) > 65535 {
		return errors.New("max_length overflow in field Metadata")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(i.Metadata))); err != nil {
		return err
	}
	if _, err := buf.WriteString(i.Metadata); err != nil {
		return err
	}
	return nil
}

func (i *IntentResult) String() string {
	return fmt.Sprint("Intent: {", i.Intent, "} ",
		"Parameters: {", i.Parameters, "} ",
		"Metadata: {", i.Metadata, "}")
}

// STRUCTURE IntentError
type IntentError struct {
	Error ErrorType
	Extra string
}

func (i *IntentError) Size() uint32 {
	var result uint32
	result += 1                    // Error ErrorType
	result += 1                    // Extra length (uint_8)
	result += uint32(len(i.Extra)) // uint_8 array
	return result
}

func (i *IntentError) Unpack(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.LittleEndian, &i.Error); err != nil {
		return err
	}
	var ExtraLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &ExtraLen); err != nil {
		return err
	}
	i.Extra = string(buf.Next(int(ExtraLen)))
	if len(i.Extra) != int(ExtraLen) {
		return errors.New("string byte mismatch")
	}
	return nil
}

func (i *IntentError) Pack(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.LittleEndian, i.Error); err != nil {
		return err
	}
	if len(i.Extra) > 255 {
		return errors.New("max_length overflow in field Extra")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(i.Extra))); err != nil {
		return err
	}
	if _, err := buf.WriteString(i.Extra); err != nil {
		return err
	}
	return nil
}

func (i *IntentError) String() string {
	return fmt.Sprint("Error: {", i.Error, "} ",
		"Extra: {", i.Extra, "}")
}

// STRUCTURE ConnectionResult
type ConnectionResult struct {
	Code            ConnectionCode
	Status          string
	NumPackets      uint8
	ExpectedPackets uint8
}

func (c *ConnectionResult) Size() uint32 {
	var result uint32
	result += 1                     // Code ConnectionCode
	result += 1                     // Status length (uint_8)
	result += uint32(len(c.Status)) // uint_8 array
	result += 1                     // NumPackets uint_8
	result += 1                     // ExpectedPackets uint_8
	return result
}

func (c *ConnectionResult) Unpack(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.LittleEndian, &c.Code); err != nil {
		return err
	}
	var StatusLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &StatusLen); err != nil {
		return err
	}
	c.Status = string(buf.Next(int(StatusLen)))
	if len(c.Status) != int(StatusLen) {
		return errors.New("string byte mismatch")
	}
	if err := binary.Read(buf, binary.LittleEndian, &c.NumPackets); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &c.ExpectedPackets); err != nil {
		return err
	}
	return nil
}

func (c *ConnectionResult) Pack(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.LittleEndian, c.Code); err != nil {
		return err
	}
	if len(c.Status) > 255 {
		return errors.New("max_length overflow in field Status")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint8(len(c.Status))); err != nil {
		return err
	}
	if _, err := buf.WriteString(c.Status); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, c.NumPackets); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, c.ExpectedPackets); err != nil {
		return err
	}
	return nil
}

func (c *ConnectionResult) String() string {
	return fmt.Sprint("Code: {", c.Code, "} ",
		"Status: {", c.Status, "} ",
		"NumPackets: {", c.NumPackets, "} ",
		"ExpectedPackets: {", c.ExpectedPackets, "}")
}

// UNION Message
type MessageTag uint8

const (
	MessageTag_Hotword          MessageTag = iota // 0
	MessageTag_Audio                              // 1
	MessageTag_AudioDone                          // 2
	MessageTag_ConnectionCheck                    // 3
	MessageTag_StopSignal                         // 4
	MessageTag_TestStarted                        // 5
	MessageTag_StreamTimeout                      // 6
	MessageTag_ConnectionResult                   // 7
	MessageTag_DebugFile                          // 8
	MessageTag_Result                             // 9
	MessageTag_Error                              // 10
	MessageTag_StreamOpen                         // 11
	MessageTag_INVALID          MessageTag = 255
)

type Message struct {
	tag   *MessageTag
	value clad.Struct
}

func (m *Message) Tag() MessageTag {
	if m.tag == nil {
		return MessageTag_INVALID
	}
	return *m.tag
}

func (m *Message) Size() uint32 {
	if m.tag == nil || *m.tag == MessageTag_INVALID {
		return 1
	}
	return 1 + m.value.Size()
}

func (m *Message) Pack(buf *bytes.Buffer) error {
	tag := MessageTag_INVALID
	if m.tag != nil {
		tag = *m.tag
	}
	if err := binary.Write(buf, binary.LittleEndian, tag); err != nil {
		return err
	}
	if tag == MessageTag_INVALID {
		return nil
	}
	return m.value.Pack(buf)
}

func (m *Message) unpackStruct(tag MessageTag, buf *bytes.Buffer) (clad.Struct, error) {
	switch tag {
	case MessageTag_Hotword:
		var ret Hotword
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_Audio:
		var ret AudioData
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_AudioDone:
		var ret Void
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_ConnectionCheck:
		var ret Void
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_StopSignal:
		var ret Void
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_TestStarted:
		var ret Void
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_StreamTimeout:
		var ret Void
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_ConnectionResult:
		var ret ConnectionResult
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_DebugFile:
		var ret Filename
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_Result:
		var ret IntentResult
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_Error:
		var ret IntentError
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	case MessageTag_StreamOpen:
		var ret StreamOpen
		if err := ret.Unpack(buf); err != nil {
			return nil, err
		}
		return &ret, nil
	default:
		return nil, errors.New("invalid tag to unpackStruct")
	}
}

func (m *Message) Unpack(buf *bytes.Buffer) error {
	tag := MessageTag_INVALID
	if err := binary.Read(buf, binary.LittleEndian, &tag); err != nil {
		return err
	}
	m.tag = &tag
	if tag == MessageTag_INVALID {
		m.value = nil
		return nil
	}
	val, err := m.unpackStruct(tag, buf)
	if err != nil {
		*m.tag = MessageTag_INVALID
		return err
	}
	m.value = val
	return nil
}

func (t MessageTag) String() string {
	switch t {
	case MessageTag_Hotword:
		return "Hotword"
	case MessageTag_Audio:
		return "Audio"
	case MessageTag_AudioDone:
		return "AudioDone"
	case MessageTag_ConnectionCheck:
		return "ConnectionCheck"
	case MessageTag_StopSignal:
		return "StopSignal"
	case MessageTag_TestStarted:
		return "TestStarted"
	case MessageTag_StreamTimeout:
		return "StreamTimeout"
	case MessageTag_ConnectionResult:
		return "ConnectionResult"
	case MessageTag_DebugFile:
		return "DebugFile"
	case MessageTag_Result:
		return "Result"
	case MessageTag_Error:
		return "Error"
	case MessageTag_StreamOpen:
		return "StreamOpen"
	default:
		return "INVALID"
	}
}

func (m *Message) String() string {
	if m.tag == nil {
		return "nil"
	}
	if *m.tag == MessageTag_INVALID {
		return "INVALID"
	}
	return fmt.Sprintf("%s: {%s}", *m.tag, m.value)
}

func (m *Message) GetHotword() *Hotword {
	if m.tag == nil || *m.tag != MessageTag_Hotword {
		return nil
	}
	return m.value.(*Hotword)
}

func (m *Message) SetHotword(value *Hotword) {
	newTag := MessageTag_Hotword
	m.tag = &newTag
	m.value = value
}

func NewMessageWithHotword(value *Hotword) *Message {
	var ret Message
	ret.SetHotword(value)
	return &ret
}

func (m *Message) GetAudio() *AudioData {
	if m.tag == nil || *m.tag != MessageTag_Audio {
		return nil
	}
	return m.value.(*AudioData)
}

func (m *Message) SetAudio(value *AudioData) {
	newTag := MessageTag_Audio
	m.tag = &newTag
	m.value = value
}

func NewMessageWithAudio(value *AudioData) *Message {
	var ret Message
	ret.SetAudio(value)
	return &ret
}

func (m *Message) GetAudioDone() *Void {
	if m.tag == nil || *m.tag != MessageTag_AudioDone {
		return nil
	}
	return m.value.(*Void)
}

func (m *Message) SetAudioDone(value *Void) {
	newTag := MessageTag_AudioDone
	m.tag = &newTag
	m.value = value
}

func NewMessageWithAudioDone(value *Void) *Message {
	var ret Message
	ret.SetAudioDone(value)
	return &ret
}

func (m *Message) GetConnectionCheck() *Void {
	if m.tag == nil || *m.tag != MessageTag_ConnectionCheck {
		return nil
	}
	return m.value.(*Void)
}

func (m *Message) SetConnectionCheck(value *Void) {
	newTag := MessageTag_ConnectionCheck
	m.tag = &newTag
	m.value = value
}

func NewMessageWithConnectionCheck(value *Void) *Message {
	var ret Message
	ret.SetConnectionCheck(value)
	return &ret
}

func (m *Message) GetStopSignal() *Void {
	if m.tag == nil || *m.tag != MessageTag_StopSignal {
		return nil
	}
	return m.value.(*Void)
}

func (m *Message) SetStopSignal(value *Void) {
	newTag := MessageTag_StopSignal
	m.tag = &newTag
	m.value = value
}

func NewMessageWithStopSignal(value *Void) *Message {
	var ret Message
	ret.SetStopSignal(value)
	return &ret
}

func (m *Message) GetTestStarted() *Void {
	if m.tag == nil || *m.tag != MessageTag_TestStarted {
		return nil
	}
	return m.value.(*Void)
}

func (m *Message) SetTestStarted(value *Void) {
	newTag := MessageTag_TestStarted
	m.tag = &newTag
	m.value = value
}

func NewMessageWithTestStarted(value *Void) *Message {
	var ret Message
	ret.SetTestStarted(value)
	return &ret
}

func (m *Message) GetStreamTimeout() *Void {
	if m.tag == nil || *m.tag != MessageTag_StreamTimeout {
		return nil
	}
	return m.value.(*Void)
}

func (m *Message) SetStreamTimeout(value *Void) {
	newTag := MessageTag_StreamTimeout
	m.tag = &newTag
	m.value = value
}

func NewMessageWithStreamTimeout(value *Void) *Message {
	var ret Message
	ret.SetStreamTimeout(value)
	return &ret
}

func (m *Message) GetConnectionResult() *ConnectionResult {
	if m.tag == nil || *m.tag != MessageTag_ConnectionResult {
		return nil
	}
	return m.value.(*ConnectionResult)
}

func (m *Message) SetConnectionResult(value *ConnectionResult) {
	newTag := MessageTag_ConnectionResult
	m.tag = &newTag
	m.value = value
}

func NewMessageWithConnectionResult(value *ConnectionResult) *Message {
	var ret Message
	ret.SetConnectionResult(value)
	return &ret
}

func (m *Message) GetDebugFile() *Filename {
	if m.tag == nil || *m.tag != MessageTag_DebugFile {
		return nil
	}
	return m.value.(*Filename)
}

func (m *Message) SetDebugFile(value *Filename) {
	newTag := MessageTag_DebugFile
	m.tag = &newTag
	m.value = value
}

func NewMessageWithDebugFile(value *Filename) *Message {
	var ret Message
	ret.SetDebugFile(value)
	return &ret
}

func (m *Message) GetResult() *IntentResult {
	if m.tag == nil || *m.tag != MessageTag_Result {
		return nil
	}
	return m.value.(*IntentResult)
}

func (m *Message) SetResult(value *IntentResult) {
	newTag := MessageTag_Result
	m.tag = &newTag
	m.value = value
}

func NewMessageWithResult(value *IntentResult) *Message {
	var ret Message
	ret.SetResult(value)
	return &ret
}

func (m *Message) GetError() *IntentError {
	if m.tag == nil || *m.tag != MessageTag_Error {
		return nil
	}
	return m.value.(*IntentError)
}

func (m *Message) SetError(value *IntentError) {
	newTag := MessageTag_Error
	m.tag = &newTag
	m.value = value
}

func NewMessageWithError(value *IntentError) *Message {
	var ret Message
	ret.SetError(value)
	return &ret
}

func (m *Message) GetStreamOpen() *StreamOpen {
	if m.tag == nil || *m.tag != MessageTag_StreamOpen {
		return nil
	}
	return m.value.(*StreamOpen)
}

func (m *Message) SetStreamOpen(value *StreamOpen) {
	newTag := MessageTag_StreamOpen
	m.tag = &newTag
	m.value = value
}

func NewMessageWithStreamOpen(value *StreamOpen) *Message {
	var ret Message
	ret.SetStreamOpen(value)
	return &ret
}