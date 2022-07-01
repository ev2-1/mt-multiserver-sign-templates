package signTemplates

import (
	"github.com/ev2-1/mt-multiserver-signs"

	"errors"
	"sync"
)

var initTemplatesMu sync.Once

func initTemplates() {
	initTemplatesMu.Do(func() {	
		registerTemplate(Text, func(pos signs.SignPos, a ...string) (error, *signs.Sign) {
			if len(a) != 1 {
				return errors.New("Not enough arguments to Template Text"), nil
			}

			return nil, &signs.Sign{
				Pos: &pos,

				Text:  "%s",
				Color: "black",
				Dyn: []signs.DynContent{
					&signs.Text{
						Text: a[0],
					},
				},
			}
		})

		registerTemplate(ServerPlayerCnt, func(pos signs.SignPos, a ...string) (error, *signs.Sign) {
			if len(a) != 1 {
				return errors.New("Not enough arguments to Template ServerPlayerCnt (Arguments: Server)"), nil
			}

			return nil, &signs.Sign{
				Pos: &pos,

				Text:  "X--------------X\n|%s|\n|%s|\nX--------------X",
				Color: "black",
				Dyn: []signs.DynContent{
					&signs.Center{
						Filler: ' ',
						Length: 12,

						Content: &signs.Text{
							Text: a[0],
						},
					},
					&signs.Center{
						Filler: ' ',
						Length: 12,

						Content: &signs.PlayerCnt{
							Srv: a[0],
						},
					},
				},

				OnClick: &signs.Hop{
					Srv: a[0],
				},
			}
		})
	})
}
