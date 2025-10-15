package reverse

import (
	"crypto/rand"
	"io"

	"github.com/luoluodaduan/xray-core/common/dice"
)

func (c *Control) FillInRandom() {
	randomLength := dice.Roll(64)
	randomLength++
	c.Random = make([]byte, randomLength)
	io.ReadFull(rand.Reader, c.Random)
}
