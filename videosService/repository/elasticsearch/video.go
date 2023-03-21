package elasticsearch

type Video struct {
	Id           string
	Title        string
	Description  string
	ChannelTitle string
}

func (esVideo Video) GetId() string {
	return esVideo.Id
}
