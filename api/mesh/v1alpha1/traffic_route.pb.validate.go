// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mesh/v1alpha1/traffic_route.proto

package v1alpha1

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
var _traffic_route_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TrafficRoute with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TrafficRoute) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSources()) < 1 {
		return TrafficRouteValidationError{
			field:  "Sources",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TrafficRouteValidationError{
					field:  fmt.Sprintf("Sources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetDestinations()) < 1 {
		return TrafficRouteValidationError{
			field:  "Destinations",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetDestinations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TrafficRouteValidationError{
					field:  fmt.Sprintf("Destinations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetConf()) < 1 {
		return TrafficRouteValidationError{
			field:  "Conf",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetConf() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TrafficRouteValidationError{
					field:  fmt.Sprintf("Conf[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TrafficRouteValidationError is the validation error returned by
// TrafficRoute.Validate if the designated constraints aren't met.
type TrafficRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrafficRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrafficRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrafficRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrafficRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrafficRouteValidationError) ErrorName() string { return "TrafficRouteValidationError" }

// Error satisfies the builtin error interface
func (e TrafficRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrafficRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrafficRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrafficRouteValidationError{}

// Validate checks the field values on TrafficRoute_WeightedDestination with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *TrafficRoute_WeightedDestination) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetWeight() < 0 {
		return TrafficRoute_WeightedDestinationValidationError{
			field:  "Weight",
			reason: "value must be greater than or equal to 0",
		}
	}

	if len(m.GetDestination()) < 1 {
		return TrafficRoute_WeightedDestinationValidationError{
			field:  "Destination",
			reason: "value must contain at least 1 pair(s)",
		}
	}

	for key, val := range m.GetDestination() {
		_ = val

		if utf8.RuneCountInString(key) < 1 {
			return TrafficRoute_WeightedDestinationValidationError{
				field:  fmt.Sprintf("Destination[%v]", key),
				reason: "value length must be at least 1 runes",
			}
		}

		if utf8.RuneCountInString(val) < 1 {
			return TrafficRoute_WeightedDestinationValidationError{
				field:  fmt.Sprintf("Destination[%v]", key),
				reason: "value length must be at least 1 runes",
			}
		}

	}

	return nil
}

// TrafficRoute_WeightedDestinationValidationError is the validation error
// returned by TrafficRoute_WeightedDestination.Validate if the designated
// constraints aren't met.
type TrafficRoute_WeightedDestinationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrafficRoute_WeightedDestinationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrafficRoute_WeightedDestinationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrafficRoute_WeightedDestinationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrafficRoute_WeightedDestinationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrafficRoute_WeightedDestinationValidationError) ErrorName() string {
	return "TrafficRoute_WeightedDestinationValidationError"
}

// Error satisfies the builtin error interface
func (e TrafficRoute_WeightedDestinationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrafficRoute_WeightedDestination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrafficRoute_WeightedDestinationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrafficRoute_WeightedDestinationValidationError{}
