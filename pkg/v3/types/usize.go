// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
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

package types

import (
	"encoding/json"
)

// Deprecated: USize is a system default unsigned number, typically used in RPC to report non-consensus
// data. It is a wrapper for [[U32]] as a WASM default (as generated by Rust bindings).
// It is not to be used, since it created consensus mismatches.
type USize uint32

// UnmarshalJSON fills u with the JSON encoded byte array given by b
func (u *USize) UnmarshalJSON(b []byte) error {
	var tmp uint32
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*u = USize(tmp)
	return nil
}

// MarshalJSON returns a JSON encoded byte array of u
func (u USize) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint32(u))
}
