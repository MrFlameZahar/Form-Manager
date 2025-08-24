package model

import (
	"time"
)

type Form struct {
	FormID      uint
	Title       string
	Description string
	Creator     User
	CreatedAt   time.Time
	FormFields  []FormField
}

type FormField struct {
	FormFieldID  uint
	DisplayOrder int
	Label        string
	FieldType    string
	Required     bool
}

type FieldType struct {
	FieldTypeID uint
	Type        string
	Description string
	Options     string
}
