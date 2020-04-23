// Copyright 2019 The OpenZipkin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package idgenerator contains several Span and Trace ID generators which can be
used by the Zipkin tracer. Additional third party generators can be plugged in
if they adhere to the IDGenerator interface.
*/
package randomUtils

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	ErrValidTraceIDRequired = errors.New("valid traceId required")
)

// ID type
type ID uint64

// String outputs the 64-bit ID as hex string.
func (i ID) String() string {
	return fmt.Sprintf("%016x", uint64(i))
}

// MarshalJSON serializes an ID type (SpanID, ParentSpanID) to HEX.
func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", i.String())), nil
}

// UnmarshalJSON deserializes an ID type (SpanID, ParentSpanID) from HEX.
func (i *ID) UnmarshalJSON(b []byte) (err error) {
	var id uint64
	if len(b) < 3 {
		return nil
	}
	id, err = strconv.ParseUint(string(b[1:len(b)-1]), 16, 64)
	*i = ID(id)
	return err
}

// TraceID is a 128 bit number internally stored as 2x uint64 (high & low).
// In case of 64 bit traceIDs, the value can be found in Low.
type TraceID struct {
	High uint64
	Low  uint64
}

// Empty returns if TraceID has zero value.
func (t TraceID) Empty() bool {
	return t.Low == 0 && t.High == 0
}

// String outputs the 128-bit traceID as hex string.
func (t TraceID) String() string {
	if t.High == 0 {
		return fmt.Sprintf("%016x", t.Low)
	}
	return fmt.Sprintf("%016x%016x", t.High, t.Low)
}

// TraceIDFromHex returns the TraceID from a hex string.
func TraceIDFromHex(h string) (t TraceID, err error) {
	if len(h) > 16 {
		if t.High, err = strconv.ParseUint(h[0:len(h)-16], 16, 64); err != nil {
			return
		}
		t.Low, err = strconv.ParseUint(h[len(h)-16:], 16, 64)
		return
	}
	t.Low, err = strconv.ParseUint(h, 16, 64)
	return
}

// MarshalJSON custom JSON serializer to export the TraceID in the required
// zero padded hex representation.
func (t TraceID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.String())), nil
}

// UnmarshalJSON custom JSON deserializer to retrieve the traceID from the hex
// encoded representation.
func (t *TraceID) UnmarshalJSON(traceID []byte) error {
	if len(traceID) < 3 {
		return ErrValidTraceIDRequired
	}
	// A valid JSON string is encoded wrapped in double quotes. We need to trim
	// these before converting the hex payload.
	tID, err := TraceIDFromHex(string(traceID[1 : len(traceID)-1]))
	if err != nil {
		return err
	}
	*t = tID
	return nil
}

var (
	seededIDGen = rand.New(rand.NewSource(time.Now().UnixNano()))
	// NewSource returns a new pseudo-random Source seeded with the given value.
	// Unlike the default Source used by top-level functions, this source is not
	// safe for concurrent use by multiple goroutines. Hence the need for a mutex.
	seededIDLock sync.Mutex
)

// IDGenerator interface can be used to provide the Zipkin Tracer with custom
// implementations to generate Span and Trace IDs.
type IDGenerator interface {
	SpanID(traceID TraceID) ID // Generates a new Span ID
	TraceID() TraceID          // Generates a new Trace ID
}

// NewRandom64 returns an ID Generator which can generate 64 bit trace and span
// id's
func NewRandom64() IDGenerator {
	return &randomID64{}
}

// NewRandom128 returns an ID Generator which can generate 128 bit trace and 64
// bit span id's
func NewRandom128() IDGenerator {
	return &randomID128{}
}

// NewRandomTimestamped generates 128 bit time sortable traceid's and 64 bit
// spanid's.
func NewRandomTimestamped() IDGenerator {
	return &randomTimestamped{}
}

// randomID64 can generate 64 bit traceid's and 64 bit spanid's.
type randomID64 struct{}

func (r *randomID64) TraceID() (id TraceID) {
	seededIDLock.Lock()
	id = TraceID{
		Low: uint64(seededIDGen.Int63()),
	}
	seededIDLock.Unlock()
	return
}

func (r *randomID64) SpanID(traceID TraceID) (id ID) {
	if !traceID.Empty() {
		return ID(traceID.Low)
	}
	seededIDLock.Lock()
	id = ID(seededIDGen.Int63())
	seededIDLock.Unlock()
	return
}

// randomID128 can generate 128 bit traceid's and 64 bit spanid's.
type randomID128 struct{}

func (r *randomID128) TraceID() (id TraceID) {
	seededIDLock.Lock()
	id = TraceID{
		High: uint64(seededIDGen.Int63()),
		Low:  uint64(seededIDGen.Int63()),
	}
	seededIDLock.Unlock()
	return
}

func (r *randomID128) SpanID(traceID TraceID) (id ID) {
	if !traceID.Empty() {
		return ID(traceID.Low)
	}
	seededIDLock.Lock()
	id = ID(seededIDGen.Int63())
	seededIDLock.Unlock()
	return
}

// randomTimestamped can generate 128 bit time sortable traceid's compatible
// with AWS X-Ray and 64 bit spanid's.
type randomTimestamped struct{}

func (t *randomTimestamped) TraceID() (id TraceID) {
	seededIDLock.Lock()
	id = TraceID{
		High: uint64(time.Now().Unix()<<32) + uint64(seededIDGen.Int31()),
		Low:  uint64(seededIDGen.Int63()),
	}
	seededIDLock.Unlock()
	return
}

func (t *randomTimestamped) SpanID(traceID TraceID) (id ID) {
	if !traceID.Empty() {
		return ID(traceID.Low)
	}
	seededIDLock.Lock()
	id = ID(seededIDGen.Int63())
	seededIDLock.Unlock()
	return
}
