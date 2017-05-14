package conversion

import (
	. "gopkg.in/check.v1"
)

type ToFloatSuite struct {}

var _ = Suite(&ToFloatSuite{})

func (suite *ToFloatSuite) ToFloatPositive(c *C, value interface{}) float64 {
	data, ok := ToFloat(value)
	c.Assert(ok, Equals, true)
	return data
}

func (suite *ToFloatSuite) ToFloatNegative(c *C, value interface{}) {
	data, ok := ToFloat(value)
	c.Assert(data, Equals, 0.0)
	c.Assert(ok, Equals, false)
}

func (suite *ToFloatSuite) TestFloatUntouched(c *C) {
	c.Assert(suite.ToFloatPositive(c, 1.1), Equals, 1.1)
}

func (suite *ToFloatSuite) TestIntConversion(c *C) {
	c.Assert(suite.ToFloatPositive(c, 123), Equals, 123.0)
}

func (suite *ToFloatSuite) TestFloatConversionFloatString(c *C) {
	c.Assert(suite.ToFloatPositive(c, "123.3"), Equals, 123.3)
}

func (suite *ToFloatSuite) TestFloatConversion(c *C) {
	c.Assert(suite.ToFloatPositive(c, 123.1), Equals, 123.1)
}

func (suite *ToFloatSuite) TestFloat32Conversion(c *C) {
	c.Assert(suite.ToFloatPositive(c, float32(123.1)), Equals, float64(float32(123.1)))
}

func (suite *ToFloatSuite) TestStringNotExtractable(c *C) {
	suite.ToFloatNegative(c, "test")
}

func (suite *ToFloatSuite) TestBoolNotExtractable(c *C) {
	suite.ToFloatNegative(c, true)
}
