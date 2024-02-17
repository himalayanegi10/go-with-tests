package maps

func (d *Dictionary) Search(searchWord string) string {
	return d.Dict[searchWord]
}

type Dictionary struct {
	Dict map[string]string
}

