package conf

type Config struct {
	User    string              `json:"user"`
	Project map[string][]string `json:"project"`
	Output  OutputConf          `json:"output"`
}

type OutputConf struct {
	Title  string   `json:"title"`
	Dir    string   `json:"dir"`
	Filter []string `json:"filter"`
}
