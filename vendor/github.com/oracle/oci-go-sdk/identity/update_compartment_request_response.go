// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateCompartmentRequest wrapper for the UpdateCompartment operation
type UpdateCompartmentRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"path" name:"compartmentId"`

	// Request object for updating a compartment.
	UpdateCompartmentDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateCompartmentRequest) String() string {
	return common.PointerString(request)
}

// UpdateCompartmentResponse wrapper for the UpdateCompartment operation
type UpdateCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Compartment instance
	Compartment `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateCompartmentResponse) String() string {
	return common.PointerString(response)
}
