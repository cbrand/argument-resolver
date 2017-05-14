package conversion

import (
	. "gopkg.in/check.v1"
)

type ToBoolSuite struct {}

var _ = Suite(&ToBoolSuite{})

func (suite *ToBoolSuite) ToBoolPositive(c *C, value interface{}) bool {
	data, ok := ToBool(value)
	c.Assert(ok, Equals, true)
	return data
}

func (suite *ToBoolSuite) ToBoolNegative(c *C, value interface{}) {
	data, ok := ToBool(value)
	c.Assert(data, Equals, false)
	c.Assert(ok, Equals, false)
}

func (suite *ToBoolSuite) TestBoolUntouched(c *C) {
	c.Assert(suite.ToBoolPositive(c, false), Equals, false)
	c.Assert(suite.ToBoolPositive(c, true), Equals, true)
}

func (suite *ToBoolSuite) TestIntConversion(c *C) {
	c.Assert(suite.ToBoolPositive(c, 123), Equals, true)
	c.Assert(suite.ToBoolPositive(c, 0), Equals, false)
}

func (suite *ToBoolSuite) TestStringConversion(c *C) {
	c.Assert(suite.ToBoolPositive(c, "true"), Equals, true)
}

func (suite *ToBoolSuite) TestFloatConversion(c *C) {
	c.Assert(suite.ToBoolPositive(c, 123.1), Equals, true)
}

func (suite *ToBoolSuite) TestFloat32Conversion(c *C) {
	c.Assert(suite.ToBoolPositive(c, float32(123.1)), Equals, true)
}

func (suite *ToBoolSuite) TestStringNotExtractable(c *C) {
	suite.ToBoolNegative(c, "test")
}

func (suite *ToBoolSuite) TestNotExtractable(c *C) {
	suite.ToBoolNegative(c, 'a')
}
