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
	"errors"
	"slices"
	"strconv"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Merge is a utility to resolve and combine errors resulting from
// evaluation. If ok is false, execution of validation should stop (either due
// to failFast or the result is not a ValidationError).
//
//nolint:errorlint
func Merge(dst, src error, failFast bool) (ok bool, err error) {
	if src == nil {
		return true, dst
	}

	srcValErrs, ok := src.(*ValidationError)
	if !ok {
		return false, src
	}

	if dst == nil {
		return !(failFast && len(srcValErrs.Violations) > 0), src
	}

	dstValErrs, ok := dst.(*ValidationError)
	if !ok {
		// what should we do here?
		return false, dst
	}

	dstValErrs.Violations = append(dstValErrs.Violations, srcValErrs.Violations...)
	return !(failFast && len(dstValErrs.Violations) > 0), dst
}

// FieldPathElement returns a buf.validate.FieldPathElement that corresponds to
// a provided FieldDescriptor. If the provided FieldDescriptor is nil, nil is
// returned.
func FieldPathElement(field protoreflect.FieldDescriptor) *validate.FieldPathElement {
	if field == nil {
		return nil
	}
	return &validate.FieldPathElement{
		FieldNumber: proto.Int32(int32(field.Number())),
		FieldName:   proto.String(field.TextName()),
		FieldType:   descriptorpb.FieldDescriptorProto_Type(field.Kind()).Enum(),
	}
}

// FieldPath returns a single-element buf.validate.FieldPath corresponding to
// the provided FieldDescriptor, or nil if the provided FieldDescriptor is nil.
func FieldPath(field protoreflect.FieldDescriptor) *validate.FieldPath {
	if field == nil {
		return nil
	}
	return &validate.FieldPath{
		Elements: []*validate.FieldPathElement{
			FieldPathElement(field),
		},
	}
}

// UpdatePaths modifies the field and rule paths of an error, appending an
// element to the end of each field path (if provided) and prepending a list of
// elements to the beginning of each rule path (if provided.)
//
// Note that this function is ordinarily used to append field paths in reverse
// order, as the stack bubbles up through the evaluators. Then, at the end, the
// path is reversed. Rule paths are generally static, so this optimization isn't
// applied for rule paths.
func UpdatePaths(err error, fieldSuffix *validate.FieldPathElement, rulePrefix []*validate.FieldPathElement) {
	if fieldSuffix == nil && len(rulePrefix) == 0 {
		return
	}
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			if fieldSuffix != nil {
				if violation.Proto.GetField() == nil {
					violation.Proto.Field = &validate.FieldPath{}
				}
				violation.Proto.Field.Elements = append(violation.Proto.Field.Elements, fieldSuffix)
			}
			if len(rulePrefix) != 0 {
				violation.Proto.Rule.Elements = append(
					append([]*validate.FieldPathElement{}, rulePrefix...),
					violation.Proto.GetRule().GetElements()...,
				)
			}
		}
	}
}

// FinalizePaths reverses all field paths in the error and populates the
// deprecated string-based field path.
func FinalizePaths(err error) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			if violation.Proto.GetField() != nil {
				slices.Reverse(violation.Proto.GetField().GetElements())
				//nolint:staticcheck // Intentional use of deprecated field
				violation.Proto.FieldPath = proto.String(FieldPathString(violation.Proto.GetField().GetElements()))
			}
		}
	}
}

// FieldPathString takes a FieldPath and encodes it to a string-based dotted
// field path.
func FieldPathString(path []*validate.FieldPathElement) string {
	var result strings.Builder
	for i, element := range path {
		if i > 0 {
			result.WriteByte('.')
		}
		result.WriteString(element.GetFieldName())
		subscript := element.GetSubscript()
		if subscript == nil {
			continue
		}
		result.WriteByte('[')
		switch value := subscript.(type) {
		case *validate.FieldPathElement_Index:
			result.WriteString(strconv.FormatUint(value.Index, 10))
		case *validate.FieldPathElement_BoolKey:
			result.WriteString(strconv.FormatBool(value.BoolKey))
		case *validate.FieldPathElement_IntKey:
			result.WriteString(strconv.FormatInt(value.IntKey, 10))
		case *validate.FieldPathElement_UintKey:
			result.WriteString(strconv.FormatUint(value.UintKey, 10))
		case *validate.FieldPathElement_StringKey:
			result.WriteString(strconv.Quote(value.StringKey))
		}
		result.WriteByte(']')
	}
	return result.String()
}

// MarkForKey marks the provided error as being for a map key, by setting the
// `for_key` flag on each violation within the validation error.
func MarkForKey(err error) {
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		for _, violation := range valErr.Violations {
			violation.Proto.ForKey = proto.Bool(true)
		}
	}
}