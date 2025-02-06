package runtime

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,gte=0,lte=130"`
}

func TestExtractValidationErrorMessage(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name      string
		input     error
		fieldName *string
		want      string
	}{
		{
			name:      "nil error",
			input:     nil,
			fieldName: nil,
			want:      "",
		},
		{
			name:      "non-validation error",
			input:     errors.New("regular error"),
			fieldName: nil,
			want:      "regular error",
		},
		{
			name:      "single validation error - required field",
			input:     validate.Struct(TestStruct{}),
			fieldName: nil,
			want:      "Field 'Name' failed validation with tag 'required'. Field 'Email' failed validation with tag 'required'. Field 'Age' failed validation with tag 'required'. ",
		},
		{
			name:      "invalid email validation error",
			input:     validate.Struct(TestStruct{Name: "John", Email: "invalid-email", Age: 25}),
			fieldName: nil,
			want:      "Field 'Email' failed validation with tag 'email'. ",
		},
		{
			name:      "validation error with custom field name",
			input:     validate.Struct(TestStruct{Name: "John", Email: "invalid-email", Age: 25}),
			fieldName: strPtr("CustomField"),
			want:      "Field 'CustomField' failed validation with tag 'email'. ",
		},
		{
			name:      "age out of range validation",
			input:     validate.Struct(TestStruct{Name: "John", Email: "john@example.com", Age: 150}),
			fieldName: nil,
			want:      "Field 'Age' failed validation with tag 'lte'. ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractValidationErrorMessage(tt.input, tt.fieldName)
			if got != tt.want {
				t.Errorf("ExtractValidationErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

// strPtr returns a pointer to the given string
func strPtr(s string) *string {
	return &s
}
