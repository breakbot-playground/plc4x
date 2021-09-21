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
type ModbusPDUWriteSingleCoilRequest struct {
	Address uint16
	Value   uint16
	Parent  *ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteSingleCoilRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteSingleCoilRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUWriteSingleCoilRequest) FunctionFlag() uint8 {
	return 0x05
}

func (m *ModbusPDUWriteSingleCoilRequest) Response() bool {
	return false
}

func (m *ModbusPDUWriteSingleCoilRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUWriteSingleCoilRequest(address uint16, value uint16) *ModbusPDU {
	child := &ModbusPDUWriteSingleCoilRequest{
		Address: address,
		Value:   value,
		Parent:  NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUWriteSingleCoilRequest(structType interface{}) *ModbusPDUWriteSingleCoilRequest {
	castFunc := func(typ interface{}) *ModbusPDUWriteSingleCoilRequest {
		if casted, ok := typ.(ModbusPDUWriteSingleCoilRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteSingleCoilRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteSingleCoilRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteSingleCoilRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteSingleCoilRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleCoilRequest"
}

func (m *ModbusPDUWriteSingleCoilRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUWriteSingleCoilRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUWriteSingleCoilRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteSingleCoilRequestParse(readBuffer utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleCoilRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (address)
	address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}

	// Simple Field (value)
	value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleCoilRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteSingleCoilRequest{
		Address: address,
		Value:   value,
		Parent:  &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUWriteSingleCoilRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleCoilRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (value)
		value := uint16(m.Value)
		_valueErr := writeBuffer.WriteUint16("value", 16, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleCoilRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteSingleCoilRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
