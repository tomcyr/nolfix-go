// /home/tomcyr/code/nolfix/id.go
package nolfix

import "github.com/google/uuid"

type IdGenerator interface {
	NextID() string
}

type UUIDGenerator struct{}

func (UUIDGenerator) NextID() string { return uuid.New().String() }
