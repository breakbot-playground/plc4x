#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License") you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from abc import ABC, abstractmethod
from dataclasses import dataclass

from ctypes import c_bool
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusPDUReadFileRecordRequestItem import (
    ModbusPDUReadFileRecordRequestItem,
)
from typing import List
import math


@dataclass
class ModbusPDUReadFileRecordRequest(PlcMessage, ModbusPDU):
    items: List[ModbusPDUReadFileRecordRequestItem]

    # Accessors for discriminator values.
    def getErrorFlag(self) -> c_bool:
        return c_bool(False)

    def getFunctionFlag(self) -> c_uint8:
        return c_uint8(0x14)

    def getResponse(self) -> c_bool:
        return c_bool(False)

    def __post_init__(self):
        super().__init__()

    def getItems(self) -> List[ModbusPDUReadFileRecordRequestItem]:
        return self.items

    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUReadFileRecordRequest")

        # Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        byteCount: c_uint8 = c_uint8((ARRAY_SIZE_IN_BYTES(self.getItems())))
        writeImplicitField("byteCount", byteCount, writeUnsignedShort(writeBuffer, 8))

        # Array Field (items)
        writeComplexTypeArrayField("items", items, writeBuffer)

        writeBuffer.popContext("ModbusPDUReadFileRecordRequest")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUReadFileRecordRequest = self

        # Implicit Field (byteCount)
        lengthInBits += 8

        # Array field
        if self.items is not None:
            for element in items:
                lengthInBits += element.getLengthInBits()

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, response: c_bool
    ) -> ModbusPDUReadFileRecordRequestBuilder:
        readBuffer.pullContext("ModbusPDUReadFileRecordRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        byteCount: c_uint8 = readImplicitField(
            "byteCount", readUnsignedShort(readBuffer, 8)
        )

        items: List[ModbusPDUReadFileRecordRequestItem] = readLengthArrayField(
            "items",
            DataReaderComplexDefault(
                ModbusPDUReadFileRecordRequestItem.staticParse(readBuffer), readBuffer
            ),
            byteCount,
        )

        readBuffer.closeContext("ModbusPDUReadFileRecordRequest")
        # Create the instance
        return ModbusPDUReadFileRecordRequestBuilder(items)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadFileRecordRequest):
            return False

        that: ModbusPDUReadFileRecordRequest = ModbusPDUReadFileRecordRequest(o)
        return (self.getItems() == that.getItems()) and super().equals(that) and True

    def hashCode(self) -> int:
        return hash(super().hashCode(), self.getItems())

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


@dataclass
class ModbusPDUReadFileRecordRequestBuilder(ModbusPDUModbusPDUBuilder):
    items: List[ModbusPDUReadFileRecordRequestItem]

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUReadFileRecordRequest:
        modbusPDUReadFileRecordRequest: ModbusPDUReadFileRecordRequest = (
            ModbusPDUReadFileRecordRequest(self.items)
        )
        return modbusPDUReadFileRecordRequest
