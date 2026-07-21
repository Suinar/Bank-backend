package fixture

// EmptyBody returns a valid empty JSON request body.
func EmptyBody() struct{} {
	return struct{}{}
}

// StringPointer returns a pointer suitable for optional string fields in tests.
func StringPointer(value string) *string {
	return &value
}

// RunePointer returns a pointer suitable for optional rune fields in tests.
func RunePointer(value rune) *rune {
	return &value
}

// Int8Pointer returns a pointer suitable for optional int8 fields in tests.
func Int8Pointer(value int8) *int8 {
	return &value
}
