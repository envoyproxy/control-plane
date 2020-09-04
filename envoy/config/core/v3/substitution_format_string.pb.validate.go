// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/substitution_format_string.proto

package envoy_config_core_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _substitution_format_string_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SubstitutionFormatString with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubstitutionFormatString) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OmitEmptyValues

	switch m.Format.(type) {

	case *SubstitutionFormatString_TextFormat:

		if len(m.GetTextFormat()) < 1 {
			return SubstitutionFormatStringValidationError{
				field:  "TextFormat",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *SubstitutionFormatString_JsonFormat:

		if m.GetJsonFormat() == nil {
			return SubstitutionFormatStringValidationError{
				field:  "JsonFormat",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetJsonFormat()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubstitutionFormatStringValidationError{
					field:  "JsonFormat",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return SubstitutionFormatStringValidationError{
			field:  "Format",
			reason: "value is required",
		}

	}

	return nil
}

// SubstitutionFormatStringValidationError is the validation error returned by
// SubstitutionFormatString.Validate if the designated constraints aren't met.
type SubstitutionFormatStringValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubstitutionFormatStringValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubstitutionFormatStringValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubstitutionFormatStringValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubstitutionFormatStringValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubstitutionFormatStringValidationError) ErrorName() string {
	return "SubstitutionFormatStringValidationError"
}

// Error satisfies the builtin error interface
func (e SubstitutionFormatStringValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubstitutionFormatString.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubstitutionFormatStringValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubstitutionFormatStringValidationError{}
