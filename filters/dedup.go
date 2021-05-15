package filters

import "sort"

type DedupFilter struct {
}

func (d *DedupFilter) Filtering(list []string) (ret []string) {
	sort.Strings(list)
	listLen := len(list)
	for i := 0; i < listLen; i++ {
		if (i > 0 && list[i-1] == list[i]) || len(list[i]) == 0 {
			continue
		}
		ret = append(ret, list[i])
	}
	return
}

func NewDedupFilter() *DedupFilter {
	return &DedupFilter{}
}
