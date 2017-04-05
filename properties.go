package siren

// Properties are custom attributes for entities.
type Properties map[string]interface{}

// Merge will combine the input properties with the target, overwriting any conflicting keys.
func (p *Properties) Merge(input Properties) {
	for key, value := range input {
		(*p)[key] = value
	}
}
