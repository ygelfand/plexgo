package plexgo

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/LukeHagar/plexgo/internal/config"
	"github.com/LukeHagar/plexgo/internal/hooks"
	"github.com/LukeHagar/plexgo/internal/utils"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"github.com/LukeHagar/plexgo/models/sdkerrors"
	"github.com/LukeHagar/plexgo/retry"
)

type HomeUsers struct {
	rootSDK          *PlexAPI
	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

func newHomeUsers(rootSDK *PlexAPI, sdkConfig config.SDKConfiguration, hooks *hooks.Hooks) *HomeUsers {
	return &HomeUsers{
		rootSDK:          rootSDK,
		sdkConfiguration: sdkConfig,
		hooks:            hooks,
	}
}

// GetHomeUsers - Get list of all users in the home
func (s *HomeUsers) GetHomeUsers(ctx context.Context, opts ...operations.Option) (*operations.GetHomeUsersResponse, error) {
	globals := operations.GetHomeUsersGlobals{
		Accepts:          s.sdkConfiguration.Globals.Accepts,
		ClientIdentifier: s.sdkConfiguration.Globals.ClientIdentifier,
		Product:          s.sdkConfiguration.Globals.Product,
		Version:          s.sdkConfiguration.Globals.Version,
		Platform:         s.sdkConfiguration.Globals.Platform,
		PlatformVersion:  s.sdkConfiguration.Globals.PlatformVersion,
		Device:           s.sdkConfiguration.Globals.Device,
		Model:            s.sdkConfiguration.Globals.Model,
		DeviceVendor:     s.sdkConfiguration.Globals.DeviceVendor,
		DeviceName:       s.sdkConfiguration.Globals.DeviceName,
		Marketplace:      s.sdkConfiguration.Globals.Marketplace,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(operations.GetHomeUsersServerList[0], map[string]string{})
	if o.ServerURL != nil {
		baseURL = *o.ServerURL
	}

	opURL, err := url.JoinPath(baseURL, "/home/users")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "get-home-users",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, operations.GetHomeUsersRequest{}, globals)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{"429", "500", "502", "503", "504"},
		}, func() (*http.Response, error) {
			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		}
	}

	res := &operations.GetHomeUsersResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	if httpRes.StatusCode == 200 {
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		var out operations.GetHomeUsersResponseBody
		if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
			return nil, err
		}
		res.Object = &out
	} else {
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// SwitchUser - Switch to a different user in the home
func (s *HomeUsers) SwitchUser(ctx context.Context, request operations.SwitchUserRequest, opts ...operations.Option) (*operations.SwitchUserResponse, error) {
	globals := operations.SwitchUserGlobals{
		Accepts:          s.sdkConfiguration.Globals.Accepts,
		ClientIdentifier: s.sdkConfiguration.Globals.ClientIdentifier,
		Product:          s.sdkConfiguration.Globals.Product,
		Version:          s.sdkConfiguration.Globals.Version,
		Platform:         s.sdkConfiguration.Globals.Platform,
		PlatformVersion:  s.sdkConfiguration.Globals.PlatformVersion,
		Device:           s.sdkConfiguration.Globals.Device,
		Model:            s.sdkConfiguration.Globals.Model,
		DeviceVendor:     s.sdkConfiguration.Globals.DeviceVendor,
		DeviceName:       s.sdkConfiguration.Globals.DeviceName,
		Marketplace:      s.sdkConfiguration.Globals.Marketplace,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(operations.SwitchUserServerList[0], map[string]string{})
	if o.ServerURL != nil {
		baseURL = *o.ServerURL
	}

	opURL, err := utils.GenerateURL(ctx, baseURL, "/home/users/{id}/switch", request, globals)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "switch-user",
		OAuth2Scopes:     nil,
		SecuritySource:   s.sdkConfiguration.Security,
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	utils.PopulateHeaders(ctx, req, request, globals)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	var httpRes *http.Response
	req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err = s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	}

	res := &operations.SwitchUserResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	if httpRes.StatusCode == 200 || httpRes.StatusCode == 201 {
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		var out components.UserPlexAccount
		if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
			return nil, err
		}
		res.UserPlexAccount = &out
	} else {
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
