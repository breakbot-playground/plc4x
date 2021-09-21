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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetServiceAckAtomicWriteFile struct {
	Parent *BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckAtomicWriteFile interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckAtomicWriteFile) ServiceChoice() uint8 {
	return 0x07
}

func (m *BACnetServiceAckAtomicWriteFile) InitializeParent(parent *BACnetServiceAck) {
}

func NewBACnetServiceAckAtomicWriteFile() *BACnetServiceAck {
	child := &BACnetServiceAckAtomicWriteFile{
		Parent: NewBACnetServiceAck(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetServiceAckAtomicWriteFile(structType interface{}) *BACnetServiceAckAtomicWriteFile {
	castFunc := func(typ interface{}) *BACnetServiceAckAtomicWriteFile {
		if casted, ok := typ.(BACnetServiceAckAtomicWriteFile); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckAtomicWriteFile); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckAtomicWriteFile(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckAtomicWriteFile(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckAtomicWriteFile) GetTypeName() string {
	return "BACnetServiceAckAtomicWriteFile"
}

func (m *BACnetServiceAckAtomicWriteFile) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAckAtomicWriteFile) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckAtomicWriteFile) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckAtomicWriteFileParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicWriteFile"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicWriteFile"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckAtomicWriteFile{
		Parent: &BACnetServiceAck{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetServiceAckAtomicWriteFile) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicWriteFile"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicWriteFile"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
