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
type BVLCOriginalBroadcastNPDU struct {
	Npdu   *NPDU
	Parent *BVLC
}

// The corresponding interface
type IBVLCOriginalBroadcastNPDU interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCOriginalBroadcastNPDU) BvlcFunction() uint8 {
	return 0x0B
}

func (m *BVLCOriginalBroadcastNPDU) InitializeParent(parent *BVLC) {
}

func NewBVLCOriginalBroadcastNPDU(npdu *NPDU) *BVLC {
	child := &BVLCOriginalBroadcastNPDU{
		Npdu:   npdu,
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCOriginalBroadcastNPDU(structType interface{}) *BVLCOriginalBroadcastNPDU {
	castFunc := func(typ interface{}) *BVLCOriginalBroadcastNPDU {
		if casted, ok := typ.(BVLCOriginalBroadcastNPDU); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCOriginalBroadcastNPDU); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCOriginalBroadcastNPDU(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCOriginalBroadcastNPDU(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCOriginalBroadcastNPDU) GetTypeName() string {
	return "BVLCOriginalBroadcastNPDU"
}

func (m *BVLCOriginalBroadcastNPDU) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCOriginalBroadcastNPDU) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (npdu)
	lengthInBits += m.Npdu.LengthInBits()

	return lengthInBits
}

func (m *BVLCOriginalBroadcastNPDU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCOriginalBroadcastNPDUParse(readBuffer utils.ReadBuffer, bvlcLength uint16) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCOriginalBroadcastNPDU"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, pullErr
	}
	npdu, _npduErr := NPDUParse(readBuffer, uint16(bvlcLength)-uint16(uint16(4)))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCOriginalBroadcastNPDU"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCOriginalBroadcastNPDU{
		Npdu:   npdu,
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCOriginalBroadcastNPDU) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCOriginalBroadcastNPDU"); pushErr != nil {
			return pushErr
		}

		// Simple Field (npdu)
		if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
			return pushErr
		}
		_npduErr := m.Npdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
			return popErr
		}
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCOriginalBroadcastNPDU"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCOriginalBroadcastNPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
