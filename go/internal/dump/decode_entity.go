//ff:func feature=dump type=decoder control=sequence
//ff:what DecodeEntity — one dump line's JSON object as an Entity; an error for anything else or a missing ID
package dump

import (
	"encoding/json"
	"errors"
)

// DecodeEntity decodes raw, one entity JSON object of the dump (TrimLine),
// into an Entity. A value that is not such an object, or has no ID, is an
// error; unknown fields are ignored (the raw line keeps them).
func DecodeEntity(raw []byte) (Entity, error) {
	var e Entity
	if err := json.Unmarshal(raw, &e); err != nil {
		return Entity{}, err
	}
	if e.ID == "" {
		return Entity{}, errors.New("dump: entity without id")
	}
	return e, nil
}
