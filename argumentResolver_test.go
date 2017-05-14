package resolver

import (
	. "gopkg.in/check.v1"
)

type ArgumentResolverSuite struct {
	data map[string]interface{}
	resolver DefaultResolver
}

var _ = Suite(&ArgumentResolverSuite{})

func (suite *ArgumentResolverSuite) SetUpTest(c *C) {
	suite.data = map[string]interface{}{
		"string": "string",
		"float": float64(0.2),
		"float32": float32(0.4),
		"int": 1,
		"bool": true,
	}
	suite.resolver = NewArgumentResolver(suite.data)
}

func (suite *ArgumentResolverSuite) TestGet(c *C) {
	_, ok := suite.resolver.Get("string")
	c.Assert(ok, Equals, true)
}

func (suite *ArgumentResolverSuite) TestGetNotFound(c *C) {
	_, ok := suite.resolver.Get("string-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestString(c *C) {
	data, ok := suite.resolver.String("string")
	c.Assert(data, Equals, "string")
	c.Assert(ok, Equals, true)
}

func (suite *ArgumentResolverSuite) TestStringNotFound(c *C) {
	_, ok := suite.resolver.String("not-existing")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestStringNoInferrence(c *C) {
	_, ok := suite.resolver.String("float")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestInt(c *C) {
	data, ok := suite.resolver.Int("int")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, 1)
}

func (suite *ArgumentResolverSuite) TestIntNotFound(c *C) {
	_, ok := suite.resolver.Int("int-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestIntNoInferrence(c *C) {
	_, ok := suite.resolver.Int("string")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestGetFloat(c *C) {
	data, ok := suite.resolver.Float("float")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, 0.2)
}

func (suite *ArgumentResolverSuite) TestFloatNotFound(c *C) {
	_, ok := suite.resolver.Float("float-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestFloatNoInferrence(c *C) {
	_, ok := suite.resolver.Float("string")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestGetFloat32(c *C) {
	data, ok := suite.resolver.Float("float32")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, float64(float32(0.4)))
}

func (suite *ArgumentResolverSuite) TestBool(c *C) {
	data, ok := suite.resolver.Bool("bool")
	c.Assert(ok, Equals, true)
	c.Assert(data, Equals, true)
}

func (suite *ArgumentResolverSuite) TestBoolNotFound(c *C) {
	_, ok := suite.resolver.Bool("bool-not-found")
	c.Assert(ok, Equals, false)
}

func (suite *ArgumentResolverSuite) TestBoolInferrence(c *C) {
	_, ok := suite.resolver.Bool("string")
	c.Assert(ok, Equals, false)
}


