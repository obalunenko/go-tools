// Copyright 2023-2024 Buf Technologies, Inc.
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

package errors

import (
	"fmt"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Violation represents a single instance where a validation rule was not met.
// It provides information about the field that caused the violation, the
// specific unfulfilled constraint, and a human-readable error message.
type Violation struct {
	// Proto contains the violation's proto.Message form.
	Proto *validate.Violation

	// FieldValue contains the value of the specific field that failed
	// validation. If there was no value, this will contain an invalid value.
	FieldValue protoreflect.Value

	// FieldDescriptor contains the field descriptor corresponding to the
	// field that failed validation.
	FieldDescriptor protoreflect.FieldDescriptor

	// RuleValue contains the value of the rule that specified the failed
	// constraint. Not all constraints have a value; only standard and
	// predefined constraints have rule values. In violations caused by other
	// kinds of constraints, like custom contraints, this will contain an
	// invalid value.
	RuleValue protoreflect.Value

	// RuleDescriptor contains the field descriptor corresponding to the
	// rule that failed validation.
	RuleDescriptor protoreflect.FieldDescriptor
}

// A ValidationError is returned if one or more constraint violations were
// detected.
type ValidationError struct {
	Violations []*Violation
}

func (err *ValidationError) Error() string {
	bldr := &strings.Builder{}
	bldr.WriteString("validation error:")
	for _, violation := range err.Violations {
		bldr.WriteString("\n - ")
		if fieldPath := FieldPathString(violation.Proto.GetField().GetElements()); fieldPath != "" {
			bldr.WriteString(fieldPath)
			bldr.WriteString(": ")
		}
		_, _ = fmt.Fprintf(bldr, "%s [%s]",
			violation.Proto.GetMessage(),
			violation.Proto.GetConstraintId())
	}
	return bldr.String()
}

// ToProto converts this error into its proto.Message form.
func (err *ValidationError) ToProto() *validate.Violations {
	violations := &validate.Violations{
		Violations: make([]*validate.Violation, len(err.Violations)),
	}
	for i, violation := range err.Violations {
		violations.Violations[i] = violation.Proto
	}
	return violations
}