package object

import (
  "fmt"
)

type ObjectType string

type Object interface {
  Type() ObjectType
  Inspect() string
}

const (
  INTEGER_OBJ = "INTEGER"
)

type Integer struct {
  Value int64
}

func (i *Integer) Inspect() strin {
  return fmt.Sprintf("%d", i.Value)
}
