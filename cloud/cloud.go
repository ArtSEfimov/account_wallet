package cloud

type CloudDB struct {
	url string
}

func NewCloudDB(url string) *CloudDB {
	return &CloudDB{url: url}
}

func (c *CloudDB) Read() ([]byte, error) {
	return []byte{}, nil
}

func (c *CloudDB) Write(data []byte) error {
	return nil
}
