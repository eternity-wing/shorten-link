package baseconversion

type Converter interface {
	Encode(number int) (string, error)
	Decode(str string) (int, error)
}
