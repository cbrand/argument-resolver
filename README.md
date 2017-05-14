# Argument Resolver #

Small Go helper library to convert maps without a specified type (e.g. `map[string]interface{}` to ints, string, 
float64 or bool entries.) The primary use case for the creation of this mini library was to parse Arguments from the
[graphql-go/graphql](https://github.com/graphql-go/graphqly) library.

## Usage ##

The library provides two different ways to parse data. Depending on the use case.

* `ArgumentResolver` - Checks directly the data for the given interface doesn't try to do any conversions. This can be
used by the `NewArgumentResolver` function.
* `ConversionResolver` - Resolves data requests with limited conversion capability. E.g. a string `"123"` is converted
to the number `123` when using the `Int(name string)` function. This can be used via the `NewConversionResolver` function.

## Example Usage ##

```go
import resolver "github.com/cbrand/argument-resolver"

data := resolver.NewConversionResolver(map[string]interface{}{
        "teststring": "123",
        "testnumber": 123,
        "testboolean": "false",
})

str, ok := data.String("teststring")
// str is "123", ok is true

// str is "", ok is false
str, ok = data.String("does-not-exist")

str, ok = data.String("testnumber")
// str is "123", ok is true

number, ok := data.Int("testnumber")
// number is 123, ok is true

floatNumber, ok := data.Float("testnumber")
// number is 123.0, ok is true

boolValue, ok := data.Bool("testboolean")
// boolValue is false, ok is true

str = data.StringOr("teststring", "as-an-alternative")
// str is "123"

str = data.StringOr("does-not-exist", "as-an-alternative")
// str is "as-an-alternative"
```

Generally these functions exist on a resolver:
* Basic request functions. Returns the requested type. All functions return as a second argument a boolean value 
to specify if an extraction succeeded or not. The following functions exist:
    * `String(name string) (string, bool)`
    * `Int(name string) (int, bool)`
    * `Float(name string) (float64, bool)` 
    * `Bool(name string) (bool, bool)`
* Alternative request functions. Returns the value behind the passed name or key. If an extraction is not possible, 
the alternative second argument is returned. Useful for default configuration options.
    * `StringOr(name string, alternative string) string`
    * `IntOr(name string, alternative int) int`
    * `FloatOr(name string, alternative float64) float64`
    * `BoolOr(name string, alternative bool) bool`

## License ##

This utility library is published under the [MIT license](MIT_LICENSE.md) 
