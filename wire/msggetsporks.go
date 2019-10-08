// Copyright (c) 2013-2015 The btcsuite developers
// Added by Aviator
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
)

// MsgGetSporks implements the Message interface and represents a bitcoin
// getsporks message.  The spork message tells the receiving peer the status
// of the spork defined by the SporkID field. Upon receiving a spork message,
// the client must verify the signature before accepting the spork message as valid.
//
// This message has no payload.
type MsgGetSporks struct{}

// BchDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetSporks) BchDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return nil
}

// BchEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgGetSporks) BchEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgGetSporks) Command() string {
	return CmdGetSporks
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgGetSporks) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgGetSporks returns a new bitcoin getsporks message that conforms to the
// Message interface.  See MsgGetSporks for details.
func NewMsgGetSporks() *MsgGetSporks {
	return &MsgGetSporks{}
}
