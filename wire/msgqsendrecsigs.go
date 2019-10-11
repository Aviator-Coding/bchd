// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
)

// MsgQsendRecSigs message is used to notify a peer to send plain LLMQ recovered signatures 
// (inventory type MSG_QUORUM_RECOVERED_SIG). Otherwise the peer would only announce/send the 
// higher level messages produced when a recovered signature is found (e.g. InstantSend islock 
// messages or ChainLock clsig messages).
// This message has no payload.
type MsgQsendRecSigs struct{}

// BchDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgQsendRecSigs) BchDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return nil
}

// BchEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgQsendRecSigs) BchEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgQsendRecSigs) Command() string {
	return CmdQsendRecSigs
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgQsendRecSigs) MaxPayloadLength(pver uint32) uint32 {
	return 1
}

// NewMsgQsendRecSigs returns a new bitcoin getaddr message that conforms to the
// Message interface.  See MsgQsendRecSigs for details.
func NewMsgQsendRecSigs() *MsgQsendRecSigs {
	return &MsgQsendRecSigs{}
}
