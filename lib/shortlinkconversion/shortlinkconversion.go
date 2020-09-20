package shortlinkconversion

import "os"

type linkConverter interface {
	Encode(number int) string
	Decode(str string) int
}

type shortLinkConvertor struct {
	convertor linkConverter
}

func InitConvertor(c linkConverter) *shortLinkConvertor {
	return &shortLinkConvertor{
		convertor: c,
	}
}

func (s *shortLinkConvertor) setConvertor(c linkConverter) {
	s.convertor = c
}

func (s *shortLinkConvertor) GetShorten(num int) string {
	return os.Getenv("DOMAIN") + "/" + s.convertor.Encode(num)
}

func (s *shortLinkConvertor) Encode(number int) string {
	return s.convertor.Encode(number)
}

func (s *shortLinkConvertor) Decode(str string) int {
	return s.convertor.Decode(str)
}
