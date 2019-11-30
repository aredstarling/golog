package golog

// Attributes to specify the different logging levels
type Attributes map[string]interface{}

// Merge multiple attributes together
func (a Attributes) Merge(attributes Attributes) Attributes {
	for k, v := range attributes {
		a[k] = v
	}

	return a
}
