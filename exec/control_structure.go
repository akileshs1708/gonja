package exec

import (
	"github.com/akileshs1708/gonja/v2/nodes"
)

type ControlStructure interface {
	nodes.ControlStructure
	Execute(*Renderer, *nodes.ControlStructureBlock) error
}
