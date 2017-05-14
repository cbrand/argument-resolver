package conversion

import (
	. "gopkg.in/check.v1"
	"errors"
)

type ToStringSuite struct {}

var _ = Suite(&ToStringSuite{})

func (suite *ToStringSuite) ToStringPositive(c *C, value interface{}) string {
	data, ok := ToString(value)
	c.Assert(ok, Equals, true)
	return data
}

func (suite *ToStringSuite) ToStringNegative(c *C, value interface{}) {
	data, ok := ToString(value)
	c.Assert(data, Equals, "")
	c.Assert(ok, Equals, false)
}

func (suite *ToStringSuite) TestStringUntouched(c *C) {
	c.Assert(suite.ToStringPositive(c, "test"), Equals, "test")
}

func (suite *ToStringSuite) TestIntConversion(c *C) {
	c.Assert(suite.ToStringPositive(c, 123), Equals, "123")
}

func (suite *ToStringSuite) TestFloatConversion(c *C) {
	c.Assert(suite.ToStringPositive(c, 123.1), Equals, "123.1")
}

func (suite *ToStringSuite) TestFloat32Conversion(c *C) {
	c.Assert(suite.ToStringPositive(c, float32(123.1)), Equals, "123.1")
}

func (suite *ToStringSuite) TestBoolConversion(c *C) {
	c.Assert(suite.ToStringPositive(c, true), Equals, "true")
	c.Assert(suite.ToStringPositive(c, false), Equals, "false")
}

type StringRepresentation struct {}

func (rep *StringRepresentation) String() string {
	return "test"
}

func (suite *ToStringSuite) TestStringRepresentationConversion(c *C) {
	c.Assert(suite.ToStringPositive(c, &StringRepresentation{}), Equals, "test")
}

func (suite *ToStringSuite) TestUnconvertableData(c *C) {
	suite.ToStringNegative(c, errors.New)
}
