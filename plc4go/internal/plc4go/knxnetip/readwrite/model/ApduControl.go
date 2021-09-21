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
type ApduControl struct {
	Child IApduControlChild
}

// The corresponding interface
type IApduControl interface {
	ControlType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IApduControlParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IApduControl, serializeChildFunction func() error) error
	GetTypeName() string
}

type IApduControlChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ApduControl)
	GetTypeName() string
	IApduControl
}

func NewApduControl() *ApduControl {
	return &ApduControl{}
}

func CastApduControl(structType interface{}) *ApduControl {
	castFunc := func(typ interface{}) *ApduControl {
		if casted, ok := typ.(ApduControl); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControl); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControl) GetTypeName() string {
	return "ApduControl"
}

func (m *ApduControl) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControl) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ApduControl) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (controlType)
	lengthInBits += 2

	return lengthInBits
}

func (m *ApduControl) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlParse(readBuffer utils.ReadBuffer) (*ApduControl, error) {
	if pullErr := readBuffer.PullContext("ApduControl"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType, _controlTypeErr := readBuffer.ReadUint8("controlType", 2)
	if _controlTypeErr != nil {
		return nil, errors.Wrap(_controlTypeErr, "Error parsing 'controlType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ApduControl
	var typeSwitchError error
	switch {
	case controlType == 0x0: // ApduControlConnect
		_parent, typeSwitchError = ApduControlConnectParse(readBuffer)
	case controlType == 0x1: // ApduControlDisconnect
		_parent, typeSwitchError = ApduControlDisconnectParse(readBuffer)
	case controlType == 0x2: // ApduControlAck
		_parent, typeSwitchError = ApduControlAckParse(readBuffer)
	case controlType == 0x3: // ApduControlNack
		_parent, typeSwitchError = ApduControlNackParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ApduControl"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ApduControl) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ApduControl) SerializeParent(writeBuffer utils.WriteBuffer, child IApduControl, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ApduControl"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType := uint8(child.ControlType())
	_controlTypeErr := writeBuffer.WriteUint8("controlType", 2, (controlType))

	if _controlTypeErr != nil {
		return errors.Wrap(_controlTypeErr, "Error serializing 'controlType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduControl"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ApduControl) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
