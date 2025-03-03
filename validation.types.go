package runtime

import "reflect"

// Copy of the "Func" type from the go-playground/validator package to decouple the package from the go-playground/validator specific version
type ValidationFunc func(fl ValidationFieldLevel) bool

// Copy of the "FieldLevel" interface from the go-playground/validator package to decouple the package from the go-playground/validator specific version
type ValidationFieldLevel interface {

	// Top returns the top level struct, if any
	Top() reflect.Value

	// Parent returns the current fields parent struct, if any or
	// the comparison value if called 'VarWithValue'
	Parent() reflect.Value

	// Field returns current field for validation
	Field() reflect.Value

	// FieldName returns the field's name with the tag
	// name taking precedence over the fields actual name.
	FieldName() string

	// StructFieldName returns the struct field's name
	StructFieldName() string

	// Param returns param for validation against current field
	Param() string

	// GetTag returns the current validations tag name
	GetTag() string

	// ExtractType gets the actual underlying type of field value.
	// It will dive into pointers, customTypes and return you the
	// underlying value and it's kind.
	ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool)

	// GetStructFieldOK traverses the parent struct to retrieve a specific field denoted by the provided namespace
	// in the param and returns the field, field kind and whether is was successful in retrieving
	// the field at all.
	//
	// NOTE: when not successful ok will be false, this can happen when a nested struct is nil and so the field
	// could not be retrieved because it didn't exist.
	//
	// Deprecated: Use GetStructFieldOK2() instead which also return if the value is nullable.
	GetStructFieldOK() (reflect.Value, reflect.Kind, bool)

	// GetStructFieldOKAdvanced is the same as GetStructFieldOK except that it accepts the parent struct to start looking for
	// the field and namespace allowing more extensibility for validators.
	//
	// Deprecated: Use GetStructFieldOKAdvanced2() instead which also return if the value is nullable.
	GetStructFieldOKAdvanced(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool)

	// GetStructFieldOK2 traverses the parent struct to retrieve a specific field denoted by the provided namespace
	// in the param and returns the field, field kind, if it's a nullable type and whether is was successful in retrieving
	// the field at all.
	//
	// NOTE: when not successful ok will be false, this can happen when a nested struct is nil and so the field
	// could not be retrieved because it didn't exist.
	GetStructFieldOK2() (reflect.Value, reflect.Kind, bool, bool)

	// GetStructFieldOKAdvanced2 is the same as GetStructFieldOK except that it accepts the parent struct to start looking for
	// the field and namespace allowing more extensibility for validators.
	GetStructFieldOKAdvanced2(val reflect.Value, namespace string) (reflect.Value, reflect.Kind, bool, bool)
}
