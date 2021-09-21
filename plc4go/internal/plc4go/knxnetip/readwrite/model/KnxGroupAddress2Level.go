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
type KnxGroupAddress2Level struct {
	MainGroup uint8
	SubGroup  uint16
	Parent    *KnxGroupAddress
}

// The corresponding interface
type IKnxGroupAddress2Level interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxGroupAddress2Level) NumLevels() uint8 {
	return 2
}

func (m *KnxGroupAddress2Level) InitializeParent(parent *KnxGroupAddress) {
}

func NewKnxGroupAddress2Level(mainGroup uint8, subGroup uint16) *KnxGroupAddress {
	child := &KnxGroupAddress2Level{
		MainGroup: mainGroup,
		SubGroup:  subGroup,
		Parent:    NewKnxGroupAddress(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastKnxGroupAddress2Level(structType interface{}) *KnxGroupAddress2Level {
	castFunc := func(typ interface{}) *KnxGroupAddress2Level {
		if casted, ok := typ.(KnxGroupAddress2Level); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxGroupAddress2Level); ok {
			return casted
		}
		if casted, ok := typ.(KnxGroupAddress); ok {
			return CastKnxGroupAddress2Level(casted.Child)
		}
		if casted, ok := typ.(*KnxGroupAddress); ok {
			return CastKnxGroupAddress2Level(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxGroupAddress2Level) GetTypeName() string {
	return "KnxGroupAddress2Level"
}

func (m *KnxGroupAddress2Level) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxGroupAddress2Level) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (mainGroup)
	lengthInBits += 5

	// Simple field (subGroup)
	lengthInBits += 11

	return lengthInBits
}

func (m *KnxGroupAddress2Level) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxGroupAddress2LevelParse(readBuffer utils.ReadBuffer) (*KnxGroupAddress, error) {
	if pullErr := readBuffer.PullContext("KnxGroupAddress2Level"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (mainGroup)
	mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 5)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field")
	}

	// Simple Field (subGroup)
	subGroup, _subGroupErr := readBuffer.ReadUint16("subGroup", 11)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field")
	}

	if closeErr := readBuffer.CloseContext("KnxGroupAddress2Level"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxGroupAddress2Level{
		MainGroup: mainGroup,
		SubGroup:  subGroup,
		Parent:    &KnxGroupAddress{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *KnxGroupAddress2Level) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddress2Level"); pushErr != nil {
			return pushErr
		}

		// Simple Field (mainGroup)
		mainGroup := uint8(m.MainGroup)
		_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 5, (mainGroup))
		if _mainGroupErr != nil {
			return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
		}

		// Simple Field (subGroup)
		subGroup := uint16(m.SubGroup)
		_subGroupErr := writeBuffer.WriteUint16("subGroup", 11, (subGroup))
		if _subGroupErr != nil {
			return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
		}

		if popErr := writeBuffer.PopContext("KnxGroupAddress2Level"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxGroupAddress2Level) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
