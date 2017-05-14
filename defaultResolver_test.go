package argument_resolver

import (
	. "gopkg.in/check.v1"
)

type DefaultResolverSuite struct {
	data map[string]interface{}
	resolver DefaultResolver
}

var _ = Suite(&DefaultResolverSuite{})

func (suite *DefaultResolverSuite) SetUpTest(c *C) {
	suite.data = map[string]interface{}{
		"string": "string",
		"float": 0.2,
		"int": 1,
		"bool": true,
	}
	suite.resolver = NewArgumentResolver(suite.data)
}

func (suite *DefaultResolverSuite) TestGetOr(c *C) {
	data := suite.resolver.GetOr("string", "should-not-be-returned")
	c.Assert(data, Equals, "string")
}

func (suite *DefaultResolverSuite) TestGeOrtNotFound(c *C) {
	data := suite.resolver.GetOr("string-not-found", "should-be-returned")
	c.Assert(data, Equals, "should-be-returned")
}

func (suite *DefaultResolverSuite) TestStringOr(c *C) {
	data := suite.resolver.StringOr("string", "not-returned")
	c.Assert(data, Equals, "string")
}

func (suite *DefaultResolverSuite) TestStringNotFound(c *C) {
	data := suite.resolver.StringOr("not-existing", "should-be-returned")
	c.Assert(data, Equals, "should-be-returned")
}

func (suite *DefaultResolverSuite) TestIntOr(c *C) {
	data := suite.resolver.IntOr("int", -1)
	c.Assert(data, Equals, 1)
}

func (suite *DefaultResolverSuite) TestIntOrNotFound(c *C) {
	data := suite.resolver.IntOr("int-not-found", -1)
	c.Assert(data, Equals, -1)
}

func (suite *DefaultResolverSuite) TestBoolOr(c *C) {
	data := suite.resolver.BoolOr("bool", false)
	c.Assert(data, Equals, true)
}

func (suite *DefaultResolverSuite) TestBoolNotFound(c *C) {
	data := suite.resolver.BoolOr("bool-not-found", true)
	c.Assert(data, Equals, true)
}
