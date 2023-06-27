// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: package/v1/topic_service.proto

package packagev1

import context "context"
import fmt "fmt"
import http "net/http"
import io "io"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// ======================
// TopicService Interface
// ======================

type TopicService interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*CreateTopicResponse, error)
}

// ============================
// TopicService Protobuf Client
// ============================

type topicServiceProtobufClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTopicServiceProtobufClient creates a Protobuf client that implements the TopicService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewTopicServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TopicService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "package.v1", "TopicService")
	urls := [1]string{
		serviceURL + "CreateTopic",
	}

	return &topicServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *topicServiceProtobufClient) CreateTopic(ctx context.Context, in *CreateTopicRequest) (*CreateTopicResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TopicService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateTopic")
	caller := c.callCreateTopic
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateTopicRequest) (*CreateTopicResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTopicRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTopicRequest) when calling interceptor")
					}
					return c.callCreateTopic(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTopicResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTopicResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *topicServiceProtobufClient) callCreateTopic(ctx context.Context, in *CreateTopicRequest) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ========================
// TopicService JSON Client
// ========================

type topicServiceJSONClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTopicServiceJSONClient creates a JSON client that implements the TopicService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewTopicServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TopicService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "package.v1", "TopicService")
	urls := [1]string{
		serviceURL + "CreateTopic",
	}

	return &topicServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *topicServiceJSONClient) CreateTopic(ctx context.Context, in *CreateTopicRequest) (*CreateTopicResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TopicService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateTopic")
	caller := c.callCreateTopic
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateTopicRequest) (*CreateTopicResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTopicRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTopicRequest) when calling interceptor")
					}
					return c.callCreateTopic(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTopicResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTopicResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *topicServiceJSONClient) callCreateTopic(ctx context.Context, in *CreateTopicRequest) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===========================
// TopicService Server Handler
// ===========================

type topicServiceServer struct {
	TopicService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewTopicServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewTopicServiceServer(svc TopicService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &topicServiceServer{
		TopicService:     svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *topicServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *topicServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// TopicServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const TopicServicePathPrefix = "/twirp/package.v1.TopicService/"

func (s *topicServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TopicService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "package.v1.TopicService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "CreateTopic":
		s.serveCreateTopic(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *topicServiceServer) serveCreateTopic(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateTopicJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateTopicProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *topicServiceServer) serveCreateTopicJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateTopic")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(CreateTopicRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TopicService.CreateTopic
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateTopicRequest) (*CreateTopicResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTopicRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTopicRequest) when calling interceptor")
					}
					return s.TopicService.CreateTopic(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTopicResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTopicResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateTopicResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateTopicResponse and nil error while calling CreateTopic. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *topicServiceServer) serveCreateTopicProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateTopic")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(CreateTopicRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TopicService.CreateTopic
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateTopicRequest) (*CreateTopicResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTopicRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTopicRequest) when calling interceptor")
					}
					return s.TopicService.CreateTopic(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTopicResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTopicResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateTopicResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateTopicResponse and nil error while calling CreateTopic. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *topicServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor3, 0
}

func (s *topicServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *topicServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "package.v1", "TopicService")
}

var twirpFileDescriptor3 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x48, 0x4c, 0xce,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0xc9, 0x2f, 0xc8, 0x4c, 0x8e, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xca, 0xeb, 0x95,
	0x19, 0x4a, 0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x45,
	0x52, 0x62, 0xe8, 0x86, 0x40, 0xc4, 0x95, 0x1c, 0xb8, 0x84, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52,
	0x43, 0x40, 0x82, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x5a, 0x5c, 0xdc, 0x10, 0x9b,
	0x4a, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x38, 0x7f, 0x39, 0xb1,
	0x15, 0xb1, 0x08, 0x30, 0x4b, 0xc8, 0x05, 0x71, 0x81, 0x65, 0x43, 0x40, 0x92, 0x4a, 0x76, 0x5c,
	0xc2, 0x28, 0x26, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xa9, 0x73, 0xb1, 0x82, 0x15, 0x81,
	0x35, 0x73, 0x1b, 0x09, 0xea, 0x21, 0x5c, 0xa9, 0x07, 0x51, 0x09, 0x91, 0x37, 0x8a, 0xe3, 0xe2,
	0x01, 0xf3, 0x83, 0x21, 0x9e, 0x12, 0xf2, 0xe3, 0xe2, 0x46, 0x32, 0x4f, 0x48, 0x0e, 0x59, 0x23,
	0xa6, 0x53, 0xa5, 0xe4, 0x71, 0xca, 0x43, 0x1c, 0xe2, 0x24, 0x16, 0x25, 0x82, 0xf0, 0xbb, 0x35,
	0x94, 0x59, 0x66, 0x98, 0xc4, 0x06, 0x0e, 0x00, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0x36, 0x4d, 0xfa, 0x5f, 0x01, 0x00, 0x00,
}
