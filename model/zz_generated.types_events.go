// Copyright 2020 The Serverless Workflow Specification Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package model

import "fmt"
import "encoding/json"
import "reflect"

// CloudEvent correlation definition
type CorrelationDef struct {
	// CloudEvent Extension Context Attribute name
	ContextAttributeName string `json:"contextAttributeName"`

	// CloudEvent Extension Context Attribute value
	ContextAttributeValue *string `json:"contextAttributeValue,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CorrelationDef) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["contextAttributeName"]; !ok || v == nil {
		return fmt.Errorf("field contextAttributeName: required")
	}
	type Plain CorrelationDef
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CorrelationDef(plain)
	return nil
}

type EventdefKind string

// Serverless Workflow specification - events schema
type Events map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventdefKind) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventdefKind {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventdefKind, v)
	}
	*j = EventdefKind(v)
	return nil
}

type Eventdef struct {
	// CloudEvent correlation definitions
	Correlation []CorrelationDef `json:"correlation,omitempty"`

	// Defines the CloudEvent as either 'consumed' or 'produced' by the workflow.
	// Default is 'consumed'
	Kind EventdefKind `json:"kind,omitempty"`

	// Metadata information
	Metadata Metadata_1 `json:"metadata,omitempty"`

	// Unique event name
	Name *string `json:"name,omitempty"`

	// List of unique event names that share the rest of the properties (source, type,
	// kind, default)
	Names []string `json:"names,omitempty"`

	// CloudEvent source
	Source *string `json:"source,omitempty"`

	// CloudEvent type
	Type *string `json:"type,omitempty"`
}

const EventdefKindConsumed EventdefKind = "consumed"
const EventdefKindProduced EventdefKind = "produced"

// UnmarshalJSON implements json.Unmarshaler.
func (j *Eventdef) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain Eventdef
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["kind"]; !ok || v == nil {
		plain.Kind = "consumed"
	}
	*j = Eventdef(plain)
	return nil
}

var enumValues_EventdefKind = []interface{}{
	"consumed",
	"produced",
}
