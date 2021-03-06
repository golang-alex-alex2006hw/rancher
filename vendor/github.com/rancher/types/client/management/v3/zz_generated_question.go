package client

const (
	QuestionType              = "question"
	QuestionFieldDefault      = "default"
	QuestionFieldDescription  = "description"
	QuestionFieldGroup        = "group"
	QuestionFieldInvalidChars = "invalidChars"
	QuestionFieldLabel        = "label"
	QuestionFieldMax          = "max"
	QuestionFieldMaxLength    = "maxLength"
	QuestionFieldMin          = "min"
	QuestionFieldMinLength    = "minLength"
	QuestionFieldOptions      = "options"
	QuestionFieldRequired     = "required"
	QuestionFieldType         = "type"
	QuestionFieldValidChars   = "validChars"
	QuestionFieldVariable     = "variable"
)

type Question struct {
	Default      string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description  string   `json:"description,omitempty" yaml:"description,omitempty"`
	Group        string   `json:"group,omitempty" yaml:"group,omitempty"`
	InvalidChars string   `json:"invalidChars,omitempty" yaml:"invalidChars,omitempty"`
	Label        string   `json:"label,omitempty" yaml:"label,omitempty"`
	Max          *int64   `json:"max,omitempty" yaml:"max,omitempty"`
	MaxLength    *int64   `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Min          *int64   `json:"min,omitempty" yaml:"min,omitempty"`
	MinLength    *int64   `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Options      []string `json:"options,omitempty" yaml:"options,omitempty"`
	Required     bool     `json:"required,omitempty" yaml:"required,omitempty"`
	Type         string   `json:"type,omitempty" yaml:"type,omitempty"`
	ValidChars   string   `json:"validChars,omitempty" yaml:"validChars,omitempty"`
	Variable     string   `json:"variable,omitempty" yaml:"variable,omitempty"`
}
