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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "alarm_message_object_query_type.h"

// Code generated by code-generation. DO NOT EDIT.


// Constant values.
static const uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_QUERY_TYPE_VARIABLE_SPEC_const = 0x12;
uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_QUERY_TYPE_VARIABLE_SPEC() {
  return PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_QUERY_TYPE_VARIABLE_SPEC_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_alarm_message_object_query_type_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_alarm_message_object_query_type** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  uint16_t curPos;
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_alarm_message_object_query_type));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (lengthDataset)
  uint8_t lengthDataset = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &lengthDataset);
  if(_res != OK) {
    return _res;
  }
  (*_message)->length_dataset = lengthDataset;

  // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
  {
    uint16_t _reserved = 0;
    _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &_reserved);
    if(_res != OK) {
      return _res;
    }
    if(_reserved != 0x0000) {
      printf("Expected constant value '%d' but got '%d' for reserved field.", 0x0000, _reserved);
    }
  }

  // Const Field (variableSpec)
  uint8_t variableSpec = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &variableSpec);
  if(_res != OK) {
    return _res;
  }
  if(variableSpec != PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_QUERY_TYPE_VARIABLE_SPEC()) {
    return PARSE_ERROR;
    // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_QUERY_TYPE_VARIABLE_SPEC + " but got " + variableSpec);
  }

  // Simple Field (eventState)
  plc4c_s7_read_write_state* eventState;
  _res = plc4c_s7_read_write_state_parse(readBuffer, (void*) &eventState);
  if(_res != OK) {
    return _res;
  }
  (*_message)->event_state = eventState;

  // Simple Field (ackStateGoing)
  plc4c_s7_read_write_state* ackStateGoing;
  _res = plc4c_s7_read_write_state_parse(readBuffer, (void*) &ackStateGoing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->ack_state_going = ackStateGoing;

  // Simple Field (ackStateComing)
  plc4c_s7_read_write_state* ackStateComing;
  _res = plc4c_s7_read_write_state_parse(readBuffer, (void*) &ackStateComing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->ack_state_coming = ackStateComing;

  // Simple Field (timeComing)
  plc4c_s7_read_write_date_and_time* timeComing;
  _res = plc4c_s7_read_write_date_and_time_parse(readBuffer, (void*) &timeComing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->time_coming = timeComing;

  // Simple Field (valueComing)
  plc4c_s7_read_write_associated_value_type* valueComing;
  _res = plc4c_s7_read_write_associated_value_type_parse(readBuffer, (void*) &valueComing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->value_coming = valueComing;

  // Simple Field (timeGoing)
  plc4c_s7_read_write_date_and_time* timeGoing;
  _res = plc4c_s7_read_write_date_and_time_parse(readBuffer, (void*) &timeGoing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->time_going = timeGoing;

  // Simple Field (valueGoing)
  plc4c_s7_read_write_associated_value_type* valueGoing;
  _res = plc4c_s7_read_write_associated_value_type_parse(readBuffer, (void*) &valueGoing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->value_going = valueGoing;

  return OK;
}

plc4c_return_code plc4c_s7_read_write_alarm_message_object_query_type_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_alarm_message_object_query_type* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (lengthDataset)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, _message->length_dataset);
  if(_res != OK) {
    return _res;
  }

  // Reserved Field
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, 0x0000);
  if(_res != OK) {
    return _res;
  }

  // Const Field (variableSpec)
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_QUERY_TYPE_VARIABLE_SPEC());

  // Simple Field (eventState)
  _res = plc4c_s7_read_write_state_serialize(writeBuffer, _message->event_state);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (ackStateGoing)
  _res = plc4c_s7_read_write_state_serialize(writeBuffer, _message->ack_state_going);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (ackStateComing)
  _res = plc4c_s7_read_write_state_serialize(writeBuffer, _message->ack_state_coming);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (timeComing)
  _res = plc4c_s7_read_write_date_and_time_serialize(writeBuffer, _message->time_coming);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (valueComing)
  _res = plc4c_s7_read_write_associated_value_type_serialize(writeBuffer, _message->value_coming);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (timeGoing)
  _res = plc4c_s7_read_write_date_and_time_serialize(writeBuffer, _message->time_going);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (valueGoing)
  _res = plc4c_s7_read_write_associated_value_type_serialize(writeBuffer, _message->value_going);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_s7_read_write_alarm_message_object_query_type_length_in_bytes(plc4c_s7_read_write_alarm_message_object_query_type* _message) {
  return plc4c_s7_read_write_alarm_message_object_query_type_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_alarm_message_object_query_type_length_in_bits(plc4c_s7_read_write_alarm_message_object_query_type* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (lengthDataset)
  lengthInBits += 8;

  // Reserved Field (reserved)
  lengthInBits += 16;

  // Const Field (variableSpec)
  lengthInBits += 8;

  // Simple field (eventState)
  lengthInBits += plc4c_s7_read_write_state_length_in_bits(_message->event_state);

  // Simple field (ackStateGoing)
  lengthInBits += plc4c_s7_read_write_state_length_in_bits(_message->ack_state_going);

  // Simple field (ackStateComing)
  lengthInBits += plc4c_s7_read_write_state_length_in_bits(_message->ack_state_coming);

  // Simple field (timeComing)
  lengthInBits += plc4c_s7_read_write_date_and_time_length_in_bits(_message->time_coming);

  // Simple field (valueComing)
  lengthInBits += plc4c_s7_read_write_associated_value_type_length_in_bits(_message->value_coming);

  // Simple field (timeGoing)
  lengthInBits += plc4c_s7_read_write_date_and_time_length_in_bits(_message->time_going);

  // Simple field (valueGoing)
  lengthInBits += plc4c_s7_read_write_associated_value_type_length_in_bits(_message->value_going);

  return lengthInBits;
}

