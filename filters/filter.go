package filters

var (
	filterQueue []Filter
)

type Filter interface {
	Filtering(list []string) (ret []string)
}

func init() {
	filterQueue = make([]Filter, 0)
	//去重过滤器
	filterQueue = append(filterQueue, NewDedupFilter())
	//关键字过滤器
	filterQueue = append(filterQueue, NewKeysFilter())
}

//过滤服务
func Filtering(list []string) (ret []string) {
	if len(filterQueue) == 0 {
		return list
	}
	for _, handle := range filterQueue {
		list = handle.Filtering(list)
	}
	return list
}
