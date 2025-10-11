package cloud

type DbCloud struct {
	url string
}

func NewDbCloud(url string) *DbCloud {
	return &DbCloud{
		url: url,
	}
}

func (db *DbCloud) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *DbCloud) Write(content []byte) {
	
}