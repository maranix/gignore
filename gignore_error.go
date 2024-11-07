package main

import (
	"bytes"
	"fmt"
)

// Current operation being performed which caused this following error.
type Operation string

// Error Kind
type Kind int8

const (
	Unknown Kind = iota
	NumberOfArguments
	UnableToConstructTemplateURL
	UnableToFetchTemplate
	TemplateNotFound
	UnableToReadResponseBody
	ResponseBodyIsEmpty
	UnableToCreateGitignoreFile
	UnableToWriteGitignoreFile
)

type Error struct {
	// Op is the operation being performed, usually translates to the name
	// of the method being invoked (Arguments, Http Get, etc).
	Op Operation

	// Kind is basically the type of Error (Number of Arguments, Request Failed,
	// etc).
	Kind Kind

	// The underlying error that caused this for debugging purposes.
	Err error
}

func (k Kind) String() string {
	switch k {
	case NumberOfArguments:
		return "Invalid Number Of Arguments"
	case UnableToConstructTemplateURL:
		return "Unable To Construct Template URL"
	case UnableToFetchTemplate:
		return "Unable To Fetch Template From URL"
	case TemplateNotFound:
		return "Template Not Found"
	case UnableToReadResponseBody:
		return "Unable To Read Response Body"
	case ResponseBodyIsEmpty:
		return "Response Body Is Empty, Aborting!"
	case UnableToCreateGitignoreFile:
		return "Unable To Create Gitignore File"
	case UnableToWriteGitignoreFile:
		return "Unable To Write Gitignore File"
	default:
		return "Unknown"
	}
}

func NewError(args ...interface{}) error {
	e := &Error{}

	if err := populateErrorFields(e, &args); err != nil {
		return err
	}

	return e
}

func (e Error) Error() string {
	b := new(bytes.Buffer)
	var err error

	writeNewLineToBuffer(b)

	if e.Op != "" {
		err = writeFieldToBuffer(b, e.Op)
	}

	if e.Kind.String() != "" {
		err = writeFieldToBuffer(b, e.Kind)
	}

	if e.Err != nil {
		err = writeFieldToBuffer(b, e.Err)
	}

	// add any err that occured during writing fields to Buffer
	writeFieldToBuffer(b, err)

	if b.Len() == 0 {
		return "Everything Ok! no errors were found."
	}

	return b.String()
}

func populateErrorFields(e *Error, args *[]interface{}) error {
	if len(*args) < 1 {
		// panic("expected to create an Error with 0 information")
	}
	for _, arg := range *args {
		switch arg := arg.(type) {
		case Operation:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case *Error:
			// Copy the Error struct to Err field if an [Error] was passed as an argument
			c := *arg
			e.Err = &c
		case error:
			e.Err = arg
		default:
			return fmt.Errorf("Unexpected type %T was provided as an argument to NewError", arg)
		}
	}

	return e
}

func writeFieldToBuffer(b *bytes.Buffer, field interface{}) error {
	writeNewLineToBuffer(b)

	switch field := field.(type) {
	case Operation:
		b.WriteString(fmt.Sprintf("\tOperation: %s", field))
	case Kind:
		b.WriteString(fmt.Sprintf("\tKind: %s", field.String()))
	case *Error:
		// We should recursively write the Err field until we there are no underlying errors.
		b.WriteString(fmt.Sprintf("\tError: %s", field.Error()))
	case error:
		b.WriteString(fmt.Sprintf("\tError: %s", field.Error()))
	default:
		return fmt.Errorf("Could not write an unknown struct field %T", field)
	}

	return nil
}

func writeStringToBuffer(b *bytes.Buffer, msg string) {
	b.WriteString(msg)
}

func writeNewLineToBuffer(b *bytes.Buffer) {
	b.WriteString("\n")
}
