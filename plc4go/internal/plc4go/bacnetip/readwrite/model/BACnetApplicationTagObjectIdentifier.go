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
type BACnetApplicationTagObjectIdentifier struct {
	*BACnetApplicationTag
	ObjectType       BACnetObjectType
	ProprietaryValue uint16
	InstanceNumber   uint32
	IsProprietary    bool
}

// The corresponding interface
type IBACnetApplicationTagObjectIdentifier interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagObjectIdentifier) TagNumber() uint8 {
	return 0xC
}

func (m *BACnetApplicationTagObjectIdentifier) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) {
	m.Header = header
}

func NewBACnetApplicationTagObjectIdentifier(objectType BACnetObjectType, proprietaryValue uint16, instanceNumber uint32, isProprietary bool, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) *BACnetApplicationTag {
	child := &BACnetApplicationTagObjectIdentifier{
		ObjectType:           objectType,
		ProprietaryValue:     proprietaryValue,
		InstanceNumber:       instanceNumber,
		IsProprietary:        isProprietary,
		BACnetApplicationTag: NewBACnetApplicationTag(header, tagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagObjectIdentifier(structType interface{}) *BACnetApplicationTagObjectIdentifier {
	castFunc := func(typ interface{}) *BACnetApplicationTagObjectIdentifier {
		if casted, ok := typ.(BACnetApplicationTagObjectIdentifier); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagObjectIdentifier); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagObjectIdentifier(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagObjectIdentifier(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagObjectIdentifier) GetTypeName() string {
	return "BACnetApplicationTagObjectIdentifier"
}

func (m *BACnetApplicationTagObjectIdentifier) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagObjectIdentifier) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Manual Field (objectType)
	lengthInBits += uint16(int32(10))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	// Simple field (instanceNumber)
	lengthInBits += 22

	return lengthInBits
}

func (m *BACnetApplicationTagObjectIdentifier) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagObjectIdentifierParse(readBuffer utils.ReadBuffer) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}

	// Manual Field (objectType)
	objectType, _objectTypeErr := ReadObjectType(readBuffer)
	if _objectTypeErr != nil {
		return nil, errors.Wrap(_objectTypeErr, "Error parsing 'objectType' field")
	}

	// Manual Field (proprietaryValue)
	proprietaryValue, _proprietaryValueErr := ReadProprietaryObjectType(readBuffer, objectType)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}

	// Virtual field
	_isProprietary := bool((objectType) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)

	// Simple Field (instanceNumber)
	_instanceNumber, _instanceNumberErr := readBuffer.ReadUint32("instanceNumber", 22)
	if _instanceNumberErr != nil {
		return nil, errors.Wrap(_instanceNumberErr, "Error parsing 'instanceNumber' field")
	}
	instanceNumber := _instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagObjectIdentifier{
		ObjectType:           objectType,
		ProprietaryValue:     proprietaryValue,
		InstanceNumber:       instanceNumber,
		IsProprietary:        isProprietary,
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagObjectIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagObjectIdentifier"); pushErr != nil {
			return pushErr
		}

		// Manual Field (objectType)
		_objectTypeErr := WriteObjectType(writeBuffer, m.ObjectType)
		if _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}

		// Manual Field (proprietaryValue)
		_proprietaryValueErr := WriteProprietaryObjectType(writeBuffer, m.ObjectType, m.ProprietaryValue)
		if _proprietaryValueErr != nil {
			return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
		}
		// Virtual field
		if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.IsProprietary); _isProprietaryErr != nil {
			return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
		}

		// Simple Field (instanceNumber)
		instanceNumber := uint32(m.InstanceNumber)
		_instanceNumberErr := writeBuffer.WriteUint32("instanceNumber", 22, (instanceNumber))
		if _instanceNumberErr != nil {
			return errors.Wrap(_instanceNumberErr, "Error serializing 'instanceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagObjectIdentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
