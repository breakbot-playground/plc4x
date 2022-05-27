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
package org.apache.plc4x.java.bacnetip;

import org.apache.commons.codec.binary.Hex;
import org.apache.commons.io.HexDump;
import org.apache.commons.lang3.ArrayUtils;
import org.apache.plc4x.java.bacnetip.readwrite.BVLC;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.ReadBufferByteBased;

import java.util.stream.IntStream;

import static org.apache.plc4x.java.bacnetip.Utils.tryParseBytes;
import static org.apache.plc4x.java.bacnetip.Utils.tryParseHex;
import static org.junit.jupiter.api.Assertions.assertNotNull;

public class ManualBacNetDecoder {

    public static void main(String[] args) throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*00000000*/  0x00, 0xc0, 0x4f, 0xa0, 0x6b, 0xb6, 0x00, 0x0d, 0x56, 0x29, 0xf7, 0x6c, 0x08, 0x00, 0x45, 0x00, //|..O.k...V).l..E.|
            /*00000010*/  0x00, 0x32, 0xbd, 0x5c, 0x00, 0x00, 0x80, 0x11, 0xfb, 0xf5, 0xc0, 0xa8, 0x00, 0x08, 0xc0, 0xa8, //|.2.\............|
            /*00000020*/  0x00, 0x10, 0xba, 0xc0, 0xba, 0xc0, 0x00, 0x1e, 0x38, 0x74, 0x81, 0x0a, 0x00, 0x16, 0x01, 0x04, //|........8t......|
            /*00000030*/  0x00, 0x03, 0x01, 0x06, 0xc4, 0x02, 0x80, 0x00, 0x02, 0x0e, 0x31, 0x00, 0x22, 0x05, 0xb4, 0x0f, //|..........1."...|
        };
        //tryParseBytes(rawBytesAsInts, 42, true);
        tryParseHex("bac0bac005cefe25810a05c60100300106100e310065fe05b42f2a23232323434f50595249474854424547494e232323230a202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d0a20436f70797269676874202843292032303033205374657665204b6172670a0a20546869732070726f6772616d206973206672656520736f6674776172653b20796f752063616e2072656469737472696275746520697420616e642f6f720a206d6f6469667920697420756e64657220746865207465726d73206f662074686520474e552047656e6572616c205075626c6963204c6963656e73650a206173207075626c697368656420627920746865204672656520536f66747761726520466f756e646174696f6e3b206569746865722076657273696f6e20320a206f6620746865204c6963656e73652c206f722028617420796f7572206f7074696f6e2920616e79206c617465722076657273696f6e2e0a0a20546869732070726f6772616d20697320646973747269627574656420696e2074686520686f706520746861742069742077696c6c2062652075736566756c2c0a2062757420574954484f555420414e592057415252414e54593b20776974686f7574206576656e2074686520696d706c6965642077617272616e7479206f660a204d45524348414e544142494c495459206f72204649544e45535320464f52204120504152544943554c415220505552504f53452e2020536565207468650a20474e552047656e6572616c205075626c6963204c6963656e736520666f72206d6f72652064657461696c732e0a0a20596f752073686f756c642068617665207265636569766564206120636f7079206f662074686520474e552047656e6572616c205075626c6963204c6963656e73650a20616c6f6e67207769746820746869732070726f6772616d3b206966206e6f742c20777269746520746f3a0a20546865204672656520536f66747761726520466f756e646174696f6e2c20496e632e0a2035392054656d706c6520506c616365202d205375697465203333300a20426f73746f6e2c204d41202030323131312d313330372c205553412e0a0a2041732061207370656369616c20657863657074696f6e2c206966206f746865722066696c657320696e7374616e74696174652074656d706c61746573206f720a20757365206d6163726f73206f7220696e6c696e652066756e6374696f6e732066726f6d20746869732066696c652c206f7220796f7520636f6d70696c650a20746869732066696c6520616e64206c696e6b2069742077697468206f7468657220776f726b7320746f2070726f64756365206120776f726b2062617365640a206f6e20746869732066696c652c20746869732066696c6520646f6573206e6f7420627920697473656c662063617573652074686520726573756c74696e670a20776f726b20746f20626520636f76657265642062792074686520474e552047656e6572616c205075626c6963204c6963656e73652e20486f77657665720a2074686520736f7572636520636f646520666f7220746869732066696c65206d757374207374696c6c206265206d61646520617661696c61626c6520696e0a206163636f7264616e636520776974682073656374696f6e20283329206f662074686520474e552047656e6572616c205075626c6963204c6963656e73652e0a0a205468697320657863657074696f6e20646f6573206e6f7420696e76616c696461746520616e79206f7468657220726561736f6e7320776879206120776f726b0a206261736564206f6e20746869732066696c65206d6967687420626520636f76657265642062792074686520474e552047656e6572616c205075626c69630a204c6963656e73652e0a202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d0a23232323434f50595249474854454e44232323232a2f0a2369666e6465662042",8,true);
    }

}
