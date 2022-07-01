package signTemplates

import (
	"github.com/ev2-1/mt-multiserver-signs"

	"sync"
)

type TemplateID uint16

const (
	Text            TemplateID = iota // insert thing here:
	ServerPlayerCnt                   // not after this
)

// templates returns a list of templates
func Templates() (m map[TemplateID]string) {
	initTemplates()

	m = make(map[TemplateID]string)

	for k := TemplateID(0); k < ServerPlayerCnt; k++ {
		m[k] = k.String()
	}

	return
}

type Template func(pos signs.SignPos, args ...string) (error, *signs.Sign)

var templates map[TemplateID]Template
var templatesMu sync.RWMutex

func registerTemplate(id TemplateID, t Template) {
	templatesMu.Lock()
	defer templatesMu.Unlock()

	if len(templates) == 0 {
		templates = make(map[TemplateID]Template)
	}

	templates[id] = t
}

func GetTemplate(t TemplateID) Template {
	initTemplates()

	templatesMu.RLock()
	defer templatesMu.RUnlock()

	return templates[t]
}
