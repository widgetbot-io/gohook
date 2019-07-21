package utils

import "errors"

var (
	ErrEventNotSpecifiedToParse = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod        = errors.New("invalid HTTP Method")
	ErrEventNotFound            = errors.New("event not defined to be parsed")
	ErrParsingPayload           = errors.New("error parsing payload")
	ErrParsingSystemPayload     = errors.New("error parsing system payload")
)
