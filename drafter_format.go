package main

type DrafterFormat int

const (
	DRAFTER_SERIALIZE_YAML DrafterFormat = iota
	DRAFTER_SERIALIZE_JSON
)

func (df DrafterFormat) String() string {
	switch df {
	case DRAFTER_SERIALIZE_YAML:
		return "drafter_serialize_yaml"
	case DRAFTER_SERIALIZE_JSON:
		return "drafter_serialize_json"
	default:
		return "unknown"
	}
}