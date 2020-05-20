package main

type Field interface {
	GetType() string
}

type BaseField struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

func (s *BaseField) GetType() string {
	return s.Type
}

type ObjectField struct {
	BaseField
	Properties map[string]Field `json:"properties"`
}

type ArrayField struct {
	BaseField
	Items Field `json:"items"`
}
