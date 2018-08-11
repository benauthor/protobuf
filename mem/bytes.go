// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package mem

// Bytes represents a byte slice created from a BytesPool.
//
// Only create these from BytesPools.
type Bytes struct {
	segListPool *segListPool
	value       []byte
	valueLen    int
}

// NewUnmanagedBytes creates a new Bytes unmanaged by any Pool.
//
// Users should generally not call this directly.
func NewUnmanagedBytes(valueLen int) *Bytes {
	return newBytes(nil, valueLen)
}

// Value gets the value.
func (b *Bytes) Value() []byte {
	return b.value[:b.valueLen]
}

// Len gets the length.
//
// This is equivalent to len(b.Value()).
func (b *Bytes) Len() int {
	return b.valueLen
}

// Truncate truncates the value to the given length.
func (b *Bytes) Truncate(valueLen int) {
	b.valueLen = valueLen
}

// MemsetZero sets the data inside the value to zero.
//
// Optimized per https://golang.org/cl/137880043
func (b *Bytes) MemsetZero() {
	for i := range b.value {
		b.value[i] = 0
	}
}

// Recycle recycles the Bytes.
//
// If this is nil, this is a no-op.
func (b *Bytes) Recycle() {
	if b == nil {
		return
	}
	if b.segListPool == nil {
		return
	}
	b.segListPool.put(b)
}

func newBytes(segListPool *segListPool, valueLen int) *Bytes {
	return &Bytes{
		segListPool: segListPool,
		value:       make([]byte, valueLen),
		valueLen:    valueLen,
	}
}