package shortlinkconversion

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

type Base62MockedConvertor struct {
	mock.Mock
}

func (m *Base62MockedConvertor) GetBaseNumber() int {
	return 62
}

func (m *Base62MockedConvertor) Encode(num int) string {
	args := m.Called(num)
	return args.String(0)
}
func (m *Base62MockedConvertor) Decode(str string) int {
	args := m.Called(str)
	return args.Int(0)
}

type stringTestCase struct {
	str   string
	value int
}

func TestShortLinkConvertor_Encode(t *testing.T) {
	var tc = stringTestCase{str: "100", value: 3844}
	c := new(Base62MockedConvertor)
	c.On("Encode", tc.value).Return(tc.str)

	var sl = InitConvertor(c)
	assert.Equal(t, sl.Encode(tc.value), tc.str)
	c.AssertExpectations(t)
}

func TestShortLinkConvertor_Decode(t *testing.T) {
	var tc = stringTestCase{str: "100", value: 3844}
	c := new(Base62MockedConvertor)
	c.On("Decode", tc.str).Return(tc.value)

	var sl = InitConvertor(c)
	assert.Equal(t, sl.Decode(tc.str), tc.value)
	c.AssertExpectations(t)
}

func TestShortLinkConvertor_GetShorten(t *testing.T) {
	var tc = stringTestCase{str: "100", value: 3844}
	c := new(Base62MockedConvertor)
	c.On("Encode", tc.value).Return(tc.str)

	var sl = InitConvertor(c)
	var expShorten = os.Getenv("DOMAIN") + "/" + tc.str
	assert.Equal(t, sl.GetShorten(tc.value), expShorten)
	c.AssertExpectations(t)
}
