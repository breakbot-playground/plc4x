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

from ctypes import c_uint16
from plc4py.api.messages.PlcMessage import PlcMessage
import math


@dataclass
class Dummy(PlcMessage):
    dummy: c_uint16

    def __post_init__(self):
        super().__init__()

    def getDummy(self) -> c_uint16:
        return self.dummy

    def serialize(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("Dummy")

        # Simple Field (dummy)
        writeSimpleField(
            "dummy",
            dummy,
            writeUnsignedInt(writeBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        writeBuffer.popContext("Dummy")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = 0
        _value: Dummy = self

        # Simple field (dummy)
        lengthInBits += 16

        return lengthInBits

    def staticParse(readBuffer: ReadBuffer, args) -> Dummy:
        positionAware: PositionAware = readBuffer
        return staticParse(readBuffer)

    @staticmethod
    def staticParseContext(readBuffer: ReadBuffer) -> Dummy:
        readBuffer.pullContext("Dummy")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        dummy: c_uint16 = readSimpleField(
            "dummy",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        readBuffer.closeContext("Dummy")
        # Create the instance
        _dummy: Dummy = Dummy(dummy)
        return _dummy

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, Dummy):
            return False

        that: Dummy = Dummy(o)
        return (self.getDummy() == that.getDummy()) and True

    def hashCode(self) -> int:
        return hash(self.getDummy())

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"
