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
type S7VarRequestParameterItem struct {
	Child IS7VarRequestParameterItemChild
}

// The corresponding interface
type IS7VarRequestParameterItem interface {
	ItemType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7VarRequestParameterItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7VarRequestParameterItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7VarRequestParameterItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7VarRequestParameterItem)
	GetTypeName() string
	IS7VarRequestParameterItem
}

func NewS7VarRequestParameterItem() *S7VarRequestParameterItem {
	return &S7VarRequestParameterItem{}
}

func CastS7VarRequestParameterItem(structType interface{}) *S7VarRequestParameterItem {
	castFunc := func(typ interface{}) *S7VarRequestParameterItem {
		if casted, ok := typ.(S7VarRequestParameterItem); ok {
			return &casted
		}
		if casted, ok := typ.(*S7VarRequestParameterItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7VarRequestParameterItem) GetTypeName() string {
	return "S7VarRequestParameterItem"
}

func (m *S7VarRequestParameterItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7VarRequestParameterItem) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *S7VarRequestParameterItem) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7VarRequestParameterItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7VarRequestParameterItemParse(readBuffer utils.ReadBuffer) (*S7VarRequestParameterItem, error) {
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItem"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := readBuffer.ReadUint8("itemType", 8)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7VarRequestParameterItem
	var typeSwitchError error
	switch {
	case itemType == 0x12: // S7VarRequestParameterItemAddress
		_parent, typeSwitchError = S7VarRequestParameterItemAddressParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *S7VarRequestParameterItem) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7VarRequestParameterItem) SerializeParent(writeBuffer utils.WriteBuffer, child IS7VarRequestParameterItem, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7VarRequestParameterItem"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint8(child.ItemType())
	_itemTypeErr := writeBuffer.WriteUint8("itemType", 8, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7VarRequestParameterItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7VarRequestParameterItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
