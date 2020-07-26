package input

import (
	"influ-dojo/api/domain/client"
)

type Favorite struct {
	Bot client.Bot `json:"-"`
}

func (f *Favorite) Favorite() error {
	return f.Bot.Favorite()
}
