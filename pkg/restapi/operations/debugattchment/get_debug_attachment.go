// Code generated by go-swagger; DO NOT EDIT.

package debugattchment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDebugAttachmentHandlerFunc turns a function with the right signature into a get debug attachment handler
type GetDebugAttachmentHandlerFunc func(GetDebugAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDebugAttachmentHandlerFunc) Handle(params GetDebugAttachmentParams) middleware.Responder {
	return fn(params)
}

// GetDebugAttachmentHandler interface for that can handle valid get debug attachment params
type GetDebugAttachmentHandler interface {
	Handle(GetDebugAttachmentParams) middleware.Responder
}

// NewGetDebugAttachment creates a new http.Handler for the get debug attachment operation
func NewGetDebugAttachment(ctx *middleware.Context, handler GetDebugAttachmentHandler) *GetDebugAttachment {
	return &GetDebugAttachment{Context: ctx, Handler: handler}
}

/*GetDebugAttachment swagger:route GET /debugattachment/{debugAttachmentId} debugattchment getDebugAttachment

Return a debug attachment

Return a debug attachment

*/
type GetDebugAttachment struct {
	Context *middleware.Context
	Handler GetDebugAttachmentHandler
}

func (o *GetDebugAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDebugAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
