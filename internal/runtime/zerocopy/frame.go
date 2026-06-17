package zerocopy

type Frame struct {
	Header []byte
	Body   []byte
}

func New(
	header []byte,
	body []byte,
) *Frame {
	return &Frame{
		Header: header,
		Body:   body,
	}
}
