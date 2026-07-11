package core_http_types

import (
	"encoding/json"

	"github.com/PhitonBedrosovich/golang-todoapp/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	// если вместо номера телефона выставили null
	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	// если выставили не null
	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return nil
	}

	n.Value = &value

	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set: n.Set,
	}
}

/*
-----------------------
JSON: {}

Nullable:
	-- Value: *nil
	-- Set: false
-----------------------
JSON: {
	"phone_nimber": "+79998887766"
}

Nullable:
	-- Value: *"+79998887766"
	-- Set: true
-----------------------
JSON: {
	"phone_nimber": null
}

Nullable:
	-- Value: *nil
	-- Set: true
-----------------------
*/