package doc

func match(note *note, key string, value []string) {
	switch key {
	case "title":
		note.Title = value[0]
	case "url":
		note.Url = value[0]
	case "header":
		note.Header = value[0]
	case "method":
		note.Method = value[0]
	case "param":
		note.Params = append(note.Params, value)
	}
}
