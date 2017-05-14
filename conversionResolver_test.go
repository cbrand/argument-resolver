package resolver

import (
	. "gopkg.in/check.v1"
)

type ConversionResolverSuite struct {
	data map[string]interface{}
	resolver DefaultResolver
}

var _ = Suite(&ConversionResolverSuite{})

func (suite *ConversionResolverSuite) SetUpTest(c *C) {
	suite.data = map[string]interface{}{
		"string": 0.2,
		"float": "0.2",
		"int": "1",
		"bool": "true",
	}
	suite.resolver = NewConversionResolver(suite.data)
}

func (suite *ConversionResolverSuite) TestGet(c *C) {
	_, ok := suite.resolver.Get("string")
	c.Assert(ok, Equals, true)
}

func (suite *ConversionResolverSuite) TestGetNotFound(c *C) {
	_, ok := suite.resolver.Get("string-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ConversionResolverSuite) TestString(c *C) {
	data, ok := suite.resolver.String("string")
	c.Assert(data, Equals, "0.2")
	c.Assert(ok, Equals, true)
}

func (suite *ConversionResolverSuite) TestStringNotFound(c *C) {
	_, ok := suite.resolver.String("not-existing")
	c.Assert(ok, Equals, false)
}

func (suite *ConversionResolverSuite) TestInt(c *C) {
	data, ok := suite.resolver.Int("int")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, 1)
}

func (suite *ConversionResolverSuite) TestIntNotFound(c *C) {
	_, ok := suite.resolver.Int("int-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ConversionResolverSuite) TestGetFloat(c *C) {
	data, ok := suite.resolver.Float("float")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, 0.2)
}

func (suite *ConversionResolverSuite) TestFloatNotFound(c *C) {
	_, ok := suite.resolver.Float("float-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ConversionResolverSuite) TestBool(c *C) {
	data, ok := suite.resolver.Bool("bool")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, true)
}

func (suite *ConversionResolverSuite) TestBoolNotFound(c *C) {
	_, ok := suite.resolver.Bool("bool-not-found")
	c.Assert(ok, Equals, false)
}
