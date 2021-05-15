package filters

import (
	"strings"

	"github.com/jqiris/week-report/conf"
)

type KeysFilter struct {
}

func (d *KeysFilter) Filtering(list []string) (ret []string) {
	cfg := conf.GetOutputConf()
	keys := cfg.Filter
	listLen := len(list)
	for i := 0; i < listLen; i++ {
		if !d.Contains(list[i], keys) {
			ret = append(ret, list[i])
		}
	}
	return
}

func (d *KeysFilter) Contains(str string, keys []string) bool {
	for _, key := range keys {
		if strings.Contains(str, key) {
			return true
		}
	}
	return false
}

func NewKeysFilter() *KeysFilter {
	return &KeysFilter{}
}
