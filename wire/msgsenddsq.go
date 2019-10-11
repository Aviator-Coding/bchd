// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
)

// MsgSendDsq implements the Message interface and represents a bitcoin
// getaddr message.  It is used to request a list of known active peers on the
// network from a peer to help identify potential nodes.  The list is returned
// via one or more addr messages (MsgAddr).
//
// This message has no payload.
type MsgSendDsq struct{}

// BchDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgSendDsq) BchDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	return nil
}

// BchEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgSendDsq) BchEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgSendDsq) Command() string {
	return CmdSendDsq
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgSendDsq) MaxPayloadLength(pver uint32) uint32 {
	return 1
}

// NewMsgSendDsq returns a new bitcoin getaddr message that conforms to the
// Message interface.  See MsgSendDsq for details.
func NewMsgSendDsq() *MsgSendDsq {
	return &MsgSendDsq{}
}
