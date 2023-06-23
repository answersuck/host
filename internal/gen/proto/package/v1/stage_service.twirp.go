// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: package/v1/stage_service.proto

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
// StageService Interface
// ======================

type StageService interface {
	// CreateStage create new stage and adds it to package.
	CreateStage(context.Context, *CreateStageRequest) (*CreateStageResponse, error)

	// ListStages returns list of package stages.
	ListStages(context.Context, *ListStagesRequest) (*ListStagesResponse, error)

	// ListStageQuestions returns list of stage questions with their topics and extra data such as cost, interval etc.
	ListStageQuestions(context.Context, *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error)
}

// ============================
// StageService Protobuf Client
// ============================

type stageServiceProtobufClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewStageServiceProtobufClient creates a Protobuf client that implements the StageService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewStageServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) StageService {
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
	serviceURL += baseServicePath(pathPrefix, "package.v1", "StageService")
	urls := [3]string{
		serviceURL + "CreateStage",
		serviceURL + "ListStages",
		serviceURL + "ListStageQuestions",
	}

	return &stageServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *stageServiceProtobufClient) CreateStage(ctx context.Context, in *CreateStageRequest) (*CreateStageResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateStage")
	caller := c.callCreateStage
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateStageRequest) (*CreateStageResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateStageRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateStageRequest) when calling interceptor")
					}
					return c.callCreateStage(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateStageResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateStageResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *stageServiceProtobufClient) callCreateStage(ctx context.Context, in *CreateStageRequest) (*CreateStageResponse, error) {
	out := new(CreateStageResponse)
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

func (c *stageServiceProtobufClient) ListStages(ctx context.Context, in *ListStagesRequest) (*ListStagesResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
	ctx = ctxsetters.WithMethodName(ctx, "ListStages")
	caller := c.callListStages
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListStagesRequest) (*ListStagesResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStagesRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStagesRequest) when calling interceptor")
					}
					return c.callListStages(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStagesResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStagesResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *stageServiceProtobufClient) callListStages(ctx context.Context, in *ListStagesRequest) (*ListStagesResponse, error) {
	out := new(ListStagesResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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

func (c *stageServiceProtobufClient) ListStageQuestions(ctx context.Context, in *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
	ctx = ctxsetters.WithMethodName(ctx, "ListStageQuestions")
	caller := c.callListStageQuestions
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStageQuestionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStageQuestionsRequest) when calling interceptor")
					}
					return c.callListStageQuestions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStageQuestionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStageQuestionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *stageServiceProtobufClient) callListStageQuestions(ctx context.Context, in *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
	out := new(ListStageQuestionsResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
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
// StageService JSON Client
// ========================

type stageServiceJSONClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewStageServiceJSONClient creates a JSON client that implements the StageService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewStageServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) StageService {
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
	serviceURL += baseServicePath(pathPrefix, "package.v1", "StageService")
	urls := [3]string{
		serviceURL + "CreateStage",
		serviceURL + "ListStages",
		serviceURL + "ListStageQuestions",
	}

	return &stageServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *stageServiceJSONClient) CreateStage(ctx context.Context, in *CreateStageRequest) (*CreateStageResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateStage")
	caller := c.callCreateStage
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateStageRequest) (*CreateStageResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateStageRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateStageRequest) when calling interceptor")
					}
					return c.callCreateStage(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateStageResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateStageResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *stageServiceJSONClient) callCreateStage(ctx context.Context, in *CreateStageRequest) (*CreateStageResponse, error) {
	out := new(CreateStageResponse)
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

func (c *stageServiceJSONClient) ListStages(ctx context.Context, in *ListStagesRequest) (*ListStagesResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
	ctx = ctxsetters.WithMethodName(ctx, "ListStages")
	caller := c.callListStages
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListStagesRequest) (*ListStagesResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStagesRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStagesRequest) when calling interceptor")
					}
					return c.callListStages(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStagesResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStagesResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *stageServiceJSONClient) callListStages(ctx context.Context, in *ListStagesRequest) (*ListStagesResponse, error) {
	out := new(ListStagesResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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

func (c *stageServiceJSONClient) ListStageQuestions(ctx context.Context, in *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
	ctx = ctxsetters.WithMethodName(ctx, "ListStageQuestions")
	caller := c.callListStageQuestions
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStageQuestionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStageQuestionsRequest) when calling interceptor")
					}
					return c.callListStageQuestions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStageQuestionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStageQuestionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *stageServiceJSONClient) callListStageQuestions(ctx context.Context, in *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
	out := new(ListStageQuestionsResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
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
// StageService Server Handler
// ===========================

type stageServiceServer struct {
	StageService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewStageServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewStageServiceServer(svc StageService, opts ...interface{}) TwirpServer {
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

	return &stageServiceServer{
		StageService:     svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *stageServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *stageServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
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

// StageServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const StageServicePathPrefix = "/twirp/package.v1.StageService/"

func (s *stageServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "package.v1")
	ctx = ctxsetters.WithServiceName(ctx, "StageService")
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
	if pkgService != "package.v1.StageService" {
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
	case "CreateStage":
		s.serveCreateStage(ctx, resp, req)
		return
	case "ListStages":
		s.serveListStages(ctx, resp, req)
		return
	case "ListStageQuestions":
		s.serveListStageQuestions(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *stageServiceServer) serveCreateStage(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateStageJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateStageProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *stageServiceServer) serveCreateStageJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateStage")
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
	reqContent := new(CreateStageRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.StageService.CreateStage
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateStageRequest) (*CreateStageResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateStageRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateStageRequest) when calling interceptor")
					}
					return s.StageService.CreateStage(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateStageResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateStageResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateStageResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateStageResponse and nil error while calling CreateStage. nil responses are not supported"))
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

func (s *stageServiceServer) serveCreateStageProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateStage")
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
	reqContent := new(CreateStageRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.StageService.CreateStage
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateStageRequest) (*CreateStageResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateStageRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateStageRequest) when calling interceptor")
					}
					return s.StageService.CreateStage(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateStageResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateStageResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateStageResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateStageResponse and nil error while calling CreateStage. nil responses are not supported"))
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

func (s *stageServiceServer) serveListStages(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListStagesJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListStagesProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *stageServiceServer) serveListStagesJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListStages")
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
	reqContent := new(ListStagesRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.StageService.ListStages
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListStagesRequest) (*ListStagesResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStagesRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStagesRequest) when calling interceptor")
					}
					return s.StageService.ListStages(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStagesResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStagesResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListStagesResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListStagesResponse and nil error while calling ListStages. nil responses are not supported"))
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

func (s *stageServiceServer) serveListStagesProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListStages")
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
	reqContent := new(ListStagesRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.StageService.ListStages
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListStagesRequest) (*ListStagesResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStagesRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStagesRequest) when calling interceptor")
					}
					return s.StageService.ListStages(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStagesResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStagesResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListStagesResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListStagesResponse and nil error while calling ListStages. nil responses are not supported"))
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

func (s *stageServiceServer) serveListStageQuestions(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListStageQuestionsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListStageQuestionsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *stageServiceServer) serveListStageQuestionsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListStageQuestions")
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
	reqContent := new(ListStageQuestionsRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.StageService.ListStageQuestions
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStageQuestionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStageQuestionsRequest) when calling interceptor")
					}
					return s.StageService.ListStageQuestions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStageQuestionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStageQuestionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListStageQuestionsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListStageQuestionsResponse and nil error while calling ListStageQuestions. nil responses are not supported"))
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

func (s *stageServiceServer) serveListStageQuestionsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListStageQuestions")
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
	reqContent := new(ListStageQuestionsRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.StageService.ListStageQuestions
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListStageQuestionsRequest) (*ListStageQuestionsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListStageQuestionsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListStageQuestionsRequest) when calling interceptor")
					}
					return s.StageService.ListStageQuestions(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListStageQuestionsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListStageQuestionsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListStageQuestionsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListStageQuestionsResponse and nil error while calling ListStageQuestions. nil responses are not supported"))
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

func (s *stageServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *stageServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *stageServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "package.v1", "StageService")
}

var twirpFileDescriptor1 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0xc4, 0x89, 0xc7, 0xa1, 0xa4, 0x53, 0x54, 0x5c, 0x4b, 0x49, 0x23, 0x0b, 0x90,
	0x11, 0xe0, 0x28, 0xe1, 0x08, 0x27, 0xc3, 0x25, 0x02, 0x15, 0x70, 0xe1, 0xc2, 0x25, 0x32, 0xf1,
	0x52, 0xad, 0xa0, 0x71, 0xea, 0x35, 0xbe, 0xf2, 0x17, 0x50, 0x7f, 0x12, 0x3f, 0x2b, 0x27, 0x94,
	0xd9, 0x75, 0xb2, 0xe9, 0x17, 0xb9, 0xad, 0x67, 0xde, 0xcc, 0x7b, 0x33, 0x6f, 0xbd, 0xd0, 0x5f,
	0xa4, 0xb3, 0x1f, 0xe9, 0x19, 0x1b, 0x56, 0xa3, 0xa1, 0x28, 0xd3, 0x33, 0x36, 0x15, 0xac, 0xa8,
	0xf8, 0x8c, 0x45, 0x8b, 0x22, 0x2f, 0x73, 0x04, 0x95, 0x8f, 0xaa, 0x91, 0xef, 0x69, 0xd8, 0x3a,
	0x4c, 0x28, 0xff, 0x50, 0xcb, 0x94, 0xf9, 0x82, 0xcf, 0x54, 0xfc, 0x61, 0x95, 0xfe, 0xe4, 0x59,
	0x5a, 0xb2, 0x61, 0x7d, 0x90, 0x89, 0xe0, 0xaf, 0x09, 0xf8, 0xa6, 0x60, 0x69, 0xc9, 0x4e, 0x57,
	0xa4, 0x09, 0xbb, 0xf8, 0xc5, 0x44, 0x89, 0x3d, 0xa8, 0xf9, 0xa6, 0x3c, 0xf3, 0x8c, 0x81, 0x11,
	0x36, 0x13, 0x47, 0x45, 0x26, 0x19, 0x86, 0x00, 0x52, 0xe3, 0x3c, 0x3d, 0x67, 0x9e, 0x39, 0x30,
	0x42, 0x27, 0x76, 0x96, 0xb1, 0x5d, 0x34, 0xba, 0x96, 0xd7, 0x4f, 0x1c, 0x4a, 0x9e, 0xa4, 0xe7,
	0x0c, 0x8f, 0xc1, 0x95, 0xc8, 0xbc, 0xc8, 0x58, 0xe1, 0x59, 0xd4, 0x49, 0x16, 0x7f, 0x58, 0x45,
	0x30, 0x85, 0xfb, 0x44, 0xc9, 0xf3, 0xf9, 0x94, 0x14, 0x0b, 0xaf, 0x31, 0xb0, 0x42, 0x77, 0xfc,
	0x3c, 0xda, 0x4c, 0x1c, 0x5d, 0x97, 0x18, 0x7d, 0x52, 0x55, 0x9f, 0x57, 0x45, 0x31, 0x2c, 0xe3,
	0xd6, 0xa5, 0xd1, 0x68, 0x5b, 0x5d, 0x48, 0xf6, 0x2e, 0xf4, 0x94, 0xf0, 0xa7, 0x70, 0x6f, 0x0b,
	0x8c, 0x3d, 0x68, 0x90, 0x70, 0xe3, 0xaa, 0x70, 0x0a, 0xe3, 0x10, 0x3a, 0x6b, 0x49, 0x3c, 0x13,
	0x9e, 0x39, 0xb0, 0xc2, 0x66, 0xdc, 0x59, 0xc6, 0xce, 0xa5, 0x61, 0xb7, 0x8d, 0x2e, 0x78, 0x46,
	0xe2, 0xd6, 0x88, 0x49, 0x26, 0x82, 0xb7, 0x70, 0xb0, 0x25, 0x50, 0x2c, 0xf2, 0xb9, 0x60, 0xf8,
	0x02, 0x5a, 0x6a, 0x04, 0x62, 0x72, 0xc7, 0x07, 0xfa, 0x48, 0x1f, 0xe5, 0x31, 0xa9, 0x31, 0xc1,
	0x17, 0x38, 0x7a, 0xcf, 0x45, 0x49, 0x3d, 0x6a, 0xbd, 0x62, 0x47, 0x43, 0x8e, 0xa0, 0x2d, 0xd7,
	0xcc, 0x33, 0xb2, 0xa3, 0x99, 0xb4, 0xe8, 0x7b, 0x92, 0x05, 0xdf, 0xc1, 0xbf, 0xa9, 0xad, 0xd2,
	0xd8, 0xdb, 0x72, 0x92, 0x16, 0xa2, 0xdb, 0xf7, 0x14, 0x6c, 0x65, 0x8a, 0x49, 0xa6, 0xec, 0xeb,
	0x13, 0xd0, 0x32, 0x13, 0x05, 0x08, 0xc6, 0xb0, 0xbf, 0xe6, 0xd9, 0x51, 0x76, 0xf0, 0x1b, 0x50,
	0xaf, 0x51, 0x9a, 0x5e, 0x83, 0x4d, 0x0a, 0x84, 0x67, 0x10, 0xe9, 0x23, 0x9d, 0xf4, 0x3a, 0x3e,
	0x92, 0x5b, 0x57, 0x35, 0xfe, 0x33, 0x68, 0x52, 0x00, 0xf7, 0xc0, 0x5c, 0x73, 0x9a, 0x3c, 0x43,
	0x54, 0xae, 0xd3, 0x75, 0x95, 0x56, 0x8f, 0xff, 0x98, 0xd0, 0x21, 0xf4, 0xa9, 0xfc, 0xd9, 0xf0,
	0x04, 0x5c, 0xcd, 0x4a, 0xec, 0xdf, 0x7d, 0x09, 0xfd, 0xe3, 0x5b, 0xf3, 0x6a, 0x96, 0x77, 0x00,
	0x1b, 0xc5, 0xd8, 0xbb, 0x6d, 0x12, 0xd9, 0xad, 0x7f, 0xf7, 0xa0, 0x38, 0xd3, 0xd6, 0xb5, 0xb6,
	0x12, 0x1f, 0xdf, 0x58, 0x75, 0xf5, 0x06, 0xf9, 0x4f, 0xfe, 0x07, 0x93, 0x24, 0xf1, 0xe1, 0xd7,
	0x07, 0x9b, 0x47, 0xe4, 0x95, 0x3a, 0x56, 0xa3, 0x6f, 0x36, 0x3d, 0x18, 0x2f, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0x86, 0x15, 0xdc, 0xa9, 0x04, 0x00, 0x00,
}
