package videosservice

type HttpError int

var errorMessages = map[int]string{
	400: "bad request",
	404: "not found",
	502: "server down",
}

func (e HttpError) Error() string {
	if message, ok := errorMessages[int(e)]; ok {
		return message
	}

	return "unknown error"
}
