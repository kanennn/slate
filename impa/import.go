package impa

type Data struct {
	Css []byte
	Markdown []byte
}

func Import() Data {
	return Data{
		Css: importCss(),
		Markdown: importMarkdown(),
	}
}