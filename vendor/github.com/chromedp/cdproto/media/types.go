package media

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// PlayerID players will get an ID that is unique within the agent context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerId
type PlayerID string

// String returns the PlayerID as string value.
func (t PlayerID) String() string {
	return string(t)
}

// Timestamp [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-Timestamp
type Timestamp float64

// Float64 returns the Timestamp as float64 value.
func (t Timestamp) Float64() float64 {
	return float64(t)
}

// PlayerProperty player Property type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerProperty
type PlayerProperty struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

// PlayerEventType break out events into different types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerEventType
type PlayerEventType string

// String returns the PlayerEventType as string value.
func (t PlayerEventType) String() string {
	return string(t)
}

// PlayerEventType values.
const (
	PlayerEventTypePlaybackEvent PlayerEventType = "playbackEvent"
	PlayerEventTypeSystemEvent   PlayerEventType = "systemEvent"
	PlayerEventTypeMessageEvent  PlayerEventType = "messageEvent"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PlayerEventType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PlayerEventType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PlayerEventType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PlayerEventType(in.String()) {
	case PlayerEventTypePlaybackEvent:
		*t = PlayerEventTypePlaybackEvent
	case PlayerEventTypeSystemEvent:
		*t = PlayerEventTypeSystemEvent
	case PlayerEventTypeMessageEvent:
		*t = PlayerEventTypeMessageEvent

	default:
		in.AddError(errors.New("unknown PlayerEventType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PlayerEventType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PlayerEvent [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerEvent
type PlayerEvent struct {
	Type      PlayerEventType `json:"type"`
	Timestamp Timestamp       `json:"timestamp"` // Events are timestamped relative to the start of the player creation not relative to the start of playback.
	Name      string          `json:"name"`
	Value     string          `json:"value"`
}
