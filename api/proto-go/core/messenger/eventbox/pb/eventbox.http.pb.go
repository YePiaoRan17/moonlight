// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: eventbox.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// EventBoxServiceHandler is the server API for EventBoxService service.
type EventBoxServiceHandler interface {
	// POST /api/dice/eventbox/message/create
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	// GET /api/dice/eventbox/register
	PrefixGet(context.Context, *PrefixGetRequest) (*PrefixGetResponse, error)
	// PUT /api/dice/eventbox/register
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// DELETE /api/dice/eventbox/register
	Del(context.Context, *DelRequest) (*DelResponse, error)
	// GET /api/dice/eventbox/version
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	// GET /api/dice/eventbox/actions/get-smtp-info
	GetSMTPInfo(context.Context, *GetSMTPInfoRequest) (*GetSMTPInfoResponse, error)
	// GET /api/dice/eventbox/webhooks
	ListHooks(context.Context, *ListHooksRequest) (*ListHooksResponse, error)
	// GET /api/dice/eventbox/webhooks/{id}
	InspectHook(context.Context, *InspectHookRequest) (*InspectHookResponse, error)
	// POST /api/dice/eventbox/webhooks
	CreateHook(context.Context, *CreateHookRequest) (*CreateHookResponse, error)
	// PUT /api/dice/eventbox/webhooks/{id}
	EditHook(context.Context, *EditHookRequest) (*EditHookResponse, error)
	// POST /api/dice/eventbox/webhooks/{id}/actions/ping
	PingHook(context.Context, *PingHookRequest) (*PingHookResponse, error)
	// DELETE /api/dice/eventbox/webhooks/{id}
	DeleteHook(context.Context, *DeleteHookRequest) (*DeleteHookResponse, error)
	// GET /api/dice/eventbox/webhook_events
	ListHookEvents(context.Context, *ListHookEventsRequest) (*ListHookEventsResponse, error)
	// GET /api/dice/eventbox/stat
	Stat(context.Context, *StatRequest) (*StatResponse, error)
}

// RegisterEventBoxServiceHandler register EventBoxServiceHandler to http.Router.
func RegisterEventBoxServiceHandler(r http.Router, srv EventBoxServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_CreateMessage := func(method, path string, fn func(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateMessageRequest))
		}
		var CreateMessage_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateMessage_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "CreateMessage", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateMessage_info)
				}
				r = r.WithContext(ctx)
				var in CreateMessageRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_PrefixGet := func(method, path string, fn func(context.Context, *PrefixGetRequest) (*PrefixGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PrefixGetRequest))
		}
		var PrefixGet_info transport.ServiceInfo
		if h.Interceptor != nil {
			PrefixGet_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "PrefixGet", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PrefixGet_info)
				}
				r = r.WithContext(ctx)
				var in PrefixGetRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Put := func(method, path string, fn func(context.Context, *PutRequest) (*PutResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PutRequest))
		}
		var Put_info transport.ServiceInfo
		if h.Interceptor != nil {
			Put_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "Put", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Put_info)
				}
				r = r.WithContext(ctx)
				var in PutRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Del := func(method, path string, fn func(context.Context, *DelRequest) (*DelResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DelRequest))
		}
		var Del_info transport.ServiceInfo
		if h.Interceptor != nil {
			Del_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "Del", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Del_info)
				}
				r = r.WithContext(ctx)
				var in DelRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetVersion := func(method, path string, fn func(context.Context, *GetVersionRequest) (*GetVersionResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetVersionRequest))
		}
		var GetVersion_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetVersion_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "GetVersion", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetVersion_info)
				}
				r = r.WithContext(ctx)
				var in GetVersionRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetSMTPInfo := func(method, path string, fn func(context.Context, *GetSMTPInfoRequest) (*GetSMTPInfoResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetSMTPInfoRequest))
		}
		var GetSMTPInfo_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetSMTPInfo_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "GetSMTPInfo", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetSMTPInfo_info)
				}
				r = r.WithContext(ctx)
				var in GetSMTPInfoRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListHooks := func(method, path string, fn func(context.Context, *ListHooksRequest) (*ListHooksResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListHooksRequest))
		}
		var ListHooks_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListHooks_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "ListHooks", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListHooks_info)
				}
				r = r.WithContext(ctx)
				var in ListHooksRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_InspectHook := func(method, path string, fn func(context.Context, *InspectHookRequest) (*InspectHookResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*InspectHookRequest))
		}
		var InspectHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			InspectHook_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "InspectHook", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, InspectHook_info)
				}
				r = r.WithContext(ctx)
				var in InspectHookRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateHook := func(method, path string, fn func(context.Context, *CreateHookRequest) (*CreateHookResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateHookRequest))
		}
		var CreateHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateHook_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "CreateHook", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateHook_info)
				}
				r = r.WithContext(ctx)
				var in CreateHookRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_EditHook := func(method, path string, fn func(context.Context, *EditHookRequest) (*EditHookResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*EditHookRequest))
		}
		var EditHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			EditHook_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "EditHook", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, EditHook_info)
				}
				r = r.WithContext(ctx)
				var in EditHookRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_PingHook := func(method, path string, fn func(context.Context, *PingHookRequest) (*PingHookResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*PingHookRequest))
		}
		var PingHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			PingHook_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "PingHook", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PingHook_info)
				}
				r = r.WithContext(ctx)
				var in PingHookRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteHook := func(method, path string, fn func(context.Context, *DeleteHookRequest) (*DeleteHookResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteHookRequest))
		}
		var DeleteHook_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteHook_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "DeleteHook", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteHook_info)
				}
				r = r.WithContext(ctx)
				var in DeleteHookRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListHookEvents := func(method, path string, fn func(context.Context, *ListHookEventsRequest) (*ListHookEventsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListHookEventsRequest))
		}
		var ListHookEvents_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListHookEvents_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "ListHookEvents", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListHookEvents_info)
				}
				r = r.WithContext(ctx)
				var in ListHookEventsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Stat := func(method, path string, fn func(context.Context, *StatRequest) (*StatResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*StatRequest))
		}
		var Stat_info transport.ServiceInfo
		if h.Interceptor != nil {
			Stat_info = transport.NewServiceInfo("erda.core.messenger.eventbox.EventBoxService", "Stat", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Stat_info)
				}
				r = r.WithContext(ctx)
				var in StatRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateMessage("POST", "/api/dice/eventbox/message/create", srv.CreateMessage)
	add_PrefixGet("GET", "/api/dice/eventbox/register", srv.PrefixGet)
	add_Put("PUT", "/api/dice/eventbox/register", srv.Put)
	add_Del("DELETE", "/api/dice/eventbox/register", srv.Del)
	add_GetVersion("GET", "/api/dice/eventbox/version", srv.GetVersion)
	add_GetSMTPInfo("GET", "/api/dice/eventbox/actions/get-smtp-info", srv.GetSMTPInfo)
	add_ListHooks("GET", "/api/dice/eventbox/webhooks", srv.ListHooks)
	add_InspectHook("GET", "/api/dice/eventbox/webhooks/{id}", srv.InspectHook)
	add_CreateHook("POST", "/api/dice/eventbox/webhooks", srv.CreateHook)
	add_EditHook("PUT", "/api/dice/eventbox/webhooks/{id}", srv.EditHook)
	add_PingHook("POST", "/api/dice/eventbox/webhooks/{id}/actions/ping", srv.PingHook)
	add_DeleteHook("DELETE", "/api/dice/eventbox/webhooks/{id}", srv.DeleteHook)
	add_ListHookEvents("GET", "/api/dice/eventbox/webhook_events", srv.ListHookEvents)
	add_Stat("GET", "/api/dice/eventbox/stat", srv.Stat)
}