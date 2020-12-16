package data

type ES struct {
	address string
}

func NewES(address string) *ES {
	return &ES{
		address: address,
	}
}

func (es *ES) SaveQA(question string, answers []string) {
	return
}

func (es *ES) SearchQA(question string) []string {
	// search data frm es
	return []string{}
}
