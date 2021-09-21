/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type DisconnectResponse struct {
	CommunicationChannelId uint8
	Status                 Status
	Parent                 *KnxNetIpMessage
}

// The corresponding interface
type IDisconnectResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DisconnectResponse) MsgType() uint16 {
	return 0x020A
}

func (m *DisconnectResponse) InitializeParent(parent *KnxNetIpMessage) {
}

func NewDisconnectResponse(communicationChannelId uint8, status Status) *KnxNetIpMessage {
	child := &DisconnectResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		Parent:                 NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastDisconnectResponse(structType interface{}) *DisconnectResponse {
	castFunc := func(typ interface{}) *DisconnectResponse {
		if casted, ok := typ.(DisconnectResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*DisconnectResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastDisconnectResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastDisconnectResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DisconnectResponse) GetTypeName() string {
	return "DisconnectResponse"
}

func (m *DisconnectResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DisconnectResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *DisconnectResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DisconnectResponseParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("DisconnectResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, pullErr
	}
	status, _statusErr := StatusParse(readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DisconnectResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DisconnectResponse{
		CommunicationChannelId: communicationChannelId,
		Status:                 status,
		Parent:                 &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *DisconnectResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DisconnectResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (communicationChannelId)
		communicationChannelId := uint8(m.CommunicationChannelId)
		_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
		if _communicationChannelIdErr != nil {
			return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
		}

		// Simple Field (status)
		if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
			return pushErr
		}
		_statusErr := m.Status.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("status"); popErr != nil {
			return popErr
		}
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("DisconnectResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *DisconnectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
