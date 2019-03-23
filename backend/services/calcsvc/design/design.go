package calcsvcdesign

import . "goa.design/goa/dsl"

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("calcserver", func() {
		Host("development", func() {
			URI("http://localhost:8082")
			URI("grpc://localhost:8083")
		})
	})
})

// Service describes a service
var _ = Service("calcsvc", func() {

	Description("The Calculator Service performs operations on numbers")

	// HTTP transport properties.
	HTTP(func() {
		Path("/calcsvc")
	})

	// Method describes a service method (endpoint)
	Method("add", func() {

		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
		})

		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)

		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/v1/add/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})

	Method("multiply", func() {

		Payload(func() {
			Attribute("a", Int, "Multiplicand")
			Attribute("b", Int, "Multiplier")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			POST("/v1/multiply")
			Response(StatusOK)
		})
	})
})

var _ = Service("openapi", func() {
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})
