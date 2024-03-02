package struct_kit

import (
	"bytes"
	"encoding/gob"
)

// DeepCopyByGob deep copy by gob
func DeepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}
