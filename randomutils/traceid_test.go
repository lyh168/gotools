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

package randomutils

import (
	"testing"
)

func TestRandom64(t *testing.T) {
	var (
		spanID  ID
		gen     = NewRandom64()
		traceID = gen.TraceID()
	)

	if traceID.Empty() {
		t.Errorf("Expected valid TraceID, got: %+v", traceID)
	}

	if want, have := uint64(0), traceID.High; want != have {
		t.Errorf("Expected TraceID.High to be 0, got %d", have)
	}

	spanID = gen.SpanID(traceID)

	if want, have := ID(traceID.Low), spanID; want != have {
		t.Errorf("Expected root span to have span ID %d, got %d", want, have)
	}

	spanID = gen.SpanID(TraceID{})

	if spanID == 0 {
		t.Errorf("Expected child span to have a valid span ID, got 0")
	}
}

func TestRandom128(t *testing.T) {
	var (
		spanID  ID
		gen     = NewRandom128()
		traceID = gen.TraceID()
	)

	if traceID.Empty() {
		t.Errorf("Expected valid TraceID, got: %+v", traceID)
	}

	if traceID.Low == 0 {
		t.Error("Expected TraceID.Low to have value, got 0")
	}

	if traceID.High == 0 {
		t.Error("Expected TraceID.High to have value, got 0")
	}

	spanID = gen.SpanID(traceID)

	if want, have := ID(traceID.Low), spanID; want != have {
		t.Errorf("Expected root span to have span ID %d, got %d", want, have)
	}

	spanID = gen.SpanID(TraceID{})

	if spanID == 0 {
		t.Errorf("Expected child span to have a valid span ID, got 0")
	}
}

func TestRandomTimeStamped(t *testing.T) {
	var (
		spanID  ID
		gen     = NewRandomTimestamped()
		traceID = gen.TraceID()
	)

	if traceID.Empty() {
		t.Errorf("Expected valid TraceID, got: %+v", traceID)
	}

	if traceID.Low == 0 {
		t.Error("Expected TraceID.Low to have value, got 0")
	}

	if traceID.High == 0 {
		t.Error("Expected TraceID.High to have value, got 0")
	}

	spanID = gen.SpanID(traceID)

	if want, have := ID(traceID.Low), spanID; want != have {
		t.Errorf("Expected root span to have span ID %d, got %d", want, have)
	}

	spanID = gen.SpanID(TraceID{})

	if spanID == 0 {
		t.Errorf("Expected child span to have a valid span ID, got 0")
	}

	// test chronological order
	var ids []TraceID

	for i := 0; i < 1000; i++ {
		ids = append(ids, gen.TraceID())
	}

	var latestTS uint64
	for idx, traceID := range ids {
		if new, old := traceID.High>>32, latestTS; new < old {
			t.Errorf("[%d] expected a higher timestamp part in traceid but got: old: %d new: %d", idx, old, new)
		}
		latestTS = traceID.High >> 32
	}

}
