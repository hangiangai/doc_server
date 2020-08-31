
package doc

func match(note *note, key string, value []string){
	switch key {
	case "title","t":
		note.Title = value[0]
	case "url","u":
		note.Url = value[0]
	case "header","h":
		note.Header = value[0]
	case "method","m":
		note.Method = value[0]
	case "param","p":
		note.Params = append(note.Params, value)
	} 
}