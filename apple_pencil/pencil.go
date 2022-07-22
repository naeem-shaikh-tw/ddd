package apple_pencil

import "ddd/item"

type applePencil struct {
}

func NewApplePencil() item.Item {
	return applePencil{}
}
