// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
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

package cmac

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"testing"
)

func TestShiftLeft(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type ShiftLeftTest struct{}

func init() { RegisterTestSuite(&ShiftLeftTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *ShiftLeftTest) NilBuffer() {
	f := func() { shiftLeft(nil) }
	ExpectThat(f, Panics(HasSubstr("empty")))
}

func (t *ShiftLeftTest) EmptyBuffer() {
	f := func() { shiftLeft([]byte{}) }
	ExpectThat(f, Panics(HasSubstr("empty")))
}

func (t *ShiftLeftTest) OneByteBuffers() {
	var input []byte
	var expected []byte

	input = []byte{fromBinary("00000000")}
	expected = []byte{fromBinary("00000000")}
	ExpectThat(shiftLeft(input), DeepEquals(expected))

	input = []byte{fromBinary("10000000")}
	expected = []byte{fromBinary("00000000")}
	ExpectThat(shiftLeft(input), DeepEquals(expected))

	input = []byte{fromBinary("00000001")}
	expected = []byte{fromBinary("00000010")}
	ExpectThat(shiftLeft(input), DeepEquals(expected))

	input = []byte{fromBinary("11000001")}
	expected = []byte{fromBinary("10000010")}
	ExpectThat(shiftLeft(input), DeepEquals(expected))
}

func (t *ShiftLeftTest) MultiByteBuffers() {
	ExpectEq("TODO", "")
}
