package main

type csvField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type csvFormat struct {
	ReadAll    bool       `json:"readAll"`
	Separator  string     `json:"separator"`
	LazyQuotes bool       `json:"lazyQuotes"`
	CsvFields  []csvField `json:"fields"`
}

// isSkip returns true if ReadAll is false, and is first line(index == 0)
func (f *csvFormat) isSkip(index int) bool {
	return !f.ReadAll && index == 0
}
