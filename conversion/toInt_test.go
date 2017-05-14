package conversion

import (
	. "gopkg.in/check.v1"
)

type ToIntSuite struct {}

var _ = Suite(&ToIntSuite{})

func (suite *ToIntSuite) ToIntPositive(c *C, value interface{}) int {
	data, ok := ToInt(value)
	c.Assert(ok, Equals, true)
	return data
}

func (suite *ToIntSuite) ToIntNegative(c *C, value interface{}) {
	data, ok := ToInt(value)
	c.Assert(data, Equals, 0)
	c.Assert(ok, Equals, false)
}

func (suite *ToIntSuite) TestIntUntouched(c *C) {
	c.Assert(suite.ToIntPositive(c, 1), Equals, 1)
}

func (suite *ToIntSuite) TestIntConversion(c *C) {
	c.Assert(suite.ToIntPositive(c, "123"), Equals, 123)
}

func (suite *ToIntSuite) TestIntConversionFloatString(c *C) {
	c.Assert(suite.ToIntPositive(c, "123.3"), Equals, 123)
}

func (suite *ToIntSuite) TestFloatConversion(c *C) {
	c.Assert(suite.ToIntPositive(c, 123.1), Equals, 123)
}

func (suite *ToIntSuite) TestFloat32Conversion(c *C) {
	c.Assert(suite.ToIntPositive(c, float32(123.1)), Equals, 123)
}

func (suite *ToIntSuite) TestStringNotExtractable(c *C) {
	suite.ToIntNegative(c, "test")
}

func (suite *ToIntSuite) TestBoolNotExtractable(c *C) {
	suite.ToIntNegative(c, true)
}
