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
from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
import math


@dataclass
class ModbusPDUDiagnosticRequest(PlcMessage, ModbusPDU):
    subFunction: c_uint16
    data: c_uint16

    # Accessors for discriminator values.
    def getErrorFlag(self) -> c_bool:
        return c_bool(False)

    def getFunctionFlag(self) -> c_uint8:
        return c_uint8(0x08)

    def getResponse(self) -> c_bool:
        return c_bool(False)

    def __post_init__(self):
        super().__init__()

    def getSubFunction(self) -> c_uint16:
        return self.subFunction

    def getData(self) -> c_uint16:
        return self.data

    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUDiagnosticRequest")

        # Simple Field (subFunction)
        writeSimpleField("subFunction", subFunction, writeUnsignedInt(writeBuffer, 16))

        # Simple Field (data)
        writeSimpleField("data", data, writeUnsignedInt(writeBuffer, 16))

        writeBuffer.popContext("ModbusPDUDiagnosticRequest")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUDiagnosticRequest = self

        # Simple field (subFunction)
        lengthInBits += 16

        # Simple field (data)
        lengthInBits += 16

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, response: c_bool
    ) -> ModbusPDUDiagnosticRequestBuilder:
        readBuffer.pullContext("ModbusPDUDiagnosticRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        subFunction: c_uint16 = readSimpleField(
            "subFunction", readUnsignedInt(readBuffer, 16)
        )

        data: c_uint16 = readSimpleField("data", readUnsignedInt(readBuffer, 16))

        readBuffer.closeContext("ModbusPDUDiagnosticRequest")
        # Create the instance
        return ModbusPDUDiagnosticRequestBuilder(subFunction, data)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUDiagnosticRequest):
            return False

        that: ModbusPDUDiagnosticRequest = ModbusPDUDiagnosticRequest(o)
        return (
            (self.getSubFunction() == that.getSubFunction())
            and (self.getData() == that.getData())
            and super().equals(that)
            and True
        )

    def hashCode(self) -> int:
        return hash(super().hashCode(), self.getSubFunction(), self.getData())

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


@dataclass
class ModbusPDUDiagnosticRequestBuilder(ModbusPDUModbusPDUBuilder):
    subFunction: c_uint16
    data: c_uint16

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUDiagnosticRequest:
        modbusPDUDiagnosticRequest: ModbusPDUDiagnosticRequest = (
            ModbusPDUDiagnosticRequest(self.subFunction, self.data)
        )
        return modbusPDUDiagnosticRequest
