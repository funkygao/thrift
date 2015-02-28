/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package thrift

import (
	"bytes"
	"testing"
)

var binaryProtoF = NewTBinaryProtocolFactoryDefault()
var compactProtoF = NewTCompactProtocolFactory()

const bufsize = 1024

var buf = bytes.NewBuffer(make([]byte, 0, bufsize))

var tfv = []TTransportFactory{
	NewTMemoryBufferTransportFactory(bufsize),
	NewStreamTransportFactory(buf, buf, true),
	NewTFramedTransportFactory(NewTMemoryBufferTransportFactory(bufsize)),
}

var listBytesOverhead = 1 + 4 // thetype(byte) * thelen(int32)

func readWriteXXBytes(bytes int) int64 {
	return int64(listBytesOverhead+bytes) * 2 // rw
}

func BenchmarkBinaryBool_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBool(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(BOOL_VALUES)))
}

func BenchmarkBinaryBool_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBool(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(BOOL_VALUES)))
}

func BenchmarkBinaryByte_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteByte(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(BYTE_VALUES)))
}

func BenchmarkBinaryByte_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteByte(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(BYTE_VALUES)))
}

func BenchmarkBinaryI16_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI16(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(INT16_VALUES) * 2))
}

func BenchmarkBinaryI16_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI16(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(INT16_VALUES) * 2))
}

func BenchmarkBinaryI32_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI32(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(INT32_VALUES) * 4))
}

func BenchmarkBinaryI32_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI32(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(INT32_VALUES) * 4))
}

func BenchmarkBinaryI64_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI64(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(INT64_VALUES) * 8))
}

func BenchmarkBinaryI64_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI64(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(INT64_VALUES) * 8))
}

func BenchmarkBinaryDouble_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteDouble(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(DOUBLE_VALUES) * 16))
}

func BenchmarkBinaryDouble_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteDouble(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(DOUBLE_VALUES) * 16))
}

func BenchmarkBinaryString_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteString(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(STRING_VALUES)))
}

func BenchmarkBinaryString_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteString(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(STRING_VALUES)))
}

func BenchmarkBinaryBinary_0(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[0].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBinary(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(protocol_bdata)))
}

func BenchmarkBinaryBinary_1(b *testing.B) {
	b.ReportAllocs()
	trans := tfv[1].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBinary(b, p, trans)
	}
	b.SetBytes(readWriteXXBytes(len(protocol_bdata)))
}

func BenchmarkBinaryBool_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBool(b, p, trans)
	}
}

func BenchmarkBinaryByte_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteByte(b, p, trans)
	}
}

func BenchmarkBinaryI16_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI16(b, p, trans)
	}
}

func BenchmarkBinaryI32_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI32(b, p, trans)
	}
}
func BenchmarkBinaryI64_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI64(b, p, trans)
	}
}
func BenchmarkBinaryDouble_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteDouble(b, p, trans)
	}
}
func BenchmarkBinaryString_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteString(b, p, trans)
	}
}
func BenchmarkBinaryBinary_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := binaryProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBinary(b, p, trans)
	}
}

func BenchmarkCompactBool_0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBool(b, p, trans)
	}
}

func BenchmarkCompactByte_0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteByte(b, p, trans)
	}
}

func BenchmarkCompactI16_0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI16(b, p, trans)
	}
}

func BenchmarkCompactI32_0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI32(b, p, trans)
	}
}
func BenchmarkCompactI64_0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI64(b, p, trans)
	}
}
func BenchmarkCompactDouble0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteDouble(b, p, trans)
	}
}
func BenchmarkCompactString0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteString(b, p, trans)
	}
}
func BenchmarkCompactBinary0(b *testing.B) {
	trans := tfv[0].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBinary(b, p, trans)
	}
}

func BenchmarkCompactBool_1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBool(b, p, trans)
	}
}

func BenchmarkCompactByte_1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteByte(b, p, trans)
	}
}

func BenchmarkCompactI16_1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI16(b, p, trans)
	}
}

func BenchmarkCompactI32_1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI32(b, p, trans)
	}
}
func BenchmarkCompactI64_1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI64(b, p, trans)
	}
}
func BenchmarkCompactDouble1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteDouble(b, p, trans)
	}
}
func BenchmarkCompactString1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteString(b, p, trans)
	}
}
func BenchmarkCompactBinary1(b *testing.B) {
	trans := tfv[1].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBinary(b, p, trans)
	}
}

func BenchmarkCompactBool_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBool(b, p, trans)
	}
}

func BenchmarkCompactByte_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteByte(b, p, trans)
	}
}

func BenchmarkCompactI16_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI16(b, p, trans)
	}
}

func BenchmarkCompactI32_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI32(b, p, trans)
	}
}
func BenchmarkCompactI64_2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteI64(b, p, trans)
	}
}
func BenchmarkCompactDouble2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteDouble(b, p, trans)
	}
}
func BenchmarkCompactString2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteString(b, p, trans)
	}
}
func BenchmarkCompactBinary2(b *testing.B) {
	trans := tfv[2].GetTransport(nil)
	p := compactProtoF.GetProtocol(trans)
	for i := 0; i < b.N; i++ {
		ReadWriteBinary(b, p, trans)
	}
}
