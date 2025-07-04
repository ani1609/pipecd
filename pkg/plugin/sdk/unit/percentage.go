// Copyright 2025 The PipeCD Authors.
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

package unit

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Percentage represents a percentage value that can be represented with or without a % suffix.
// It stores both the numeric value and whether it originally had a % suffix.
type Percentage struct {
	Number    int  `json:",omitempty"`
	HasSuffix bool `json:",omitempty"`
}

// String returns the string representation of the percentage.
// If HasSuffix is true, it includes the % suffix.
func (p Percentage) String() string {
	s := strconv.FormatInt(int64(p.Number), 10)
	if p.HasSuffix {
		return s + "%"
	}
	return s
}

// Int returns the numeric value of the percentage.
func (p Percentage) Int() int {
	return p.Number
}

// MarshalJSON marshals the Percentage to a JSON string representation.
func (p Percentage) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// UnmarshalJSON unmarshals a JSON string to Percentage.
// It accepts both plain numbers (e.g., "50") and percentage strings (e.g., "50%").
func (p *Percentage) UnmarshalJSON(b []byte) error {
	raw := strings.Trim(string(b), `"`)
	percentage := Percentage{
		HasSuffix: false,
	}
	if strings.HasSuffix(raw, "%") {
		percentage.HasSuffix = true
		raw = strings.TrimSuffix(raw, "%")
	}
	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid percentage: %w", err)
	}
	percentage.Number = int(value)
	*p = percentage
	return nil
}
