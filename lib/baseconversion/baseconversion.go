package baseconversion

type Converter interface {
	Encode(number int) string
	Decode(str string) int
}
