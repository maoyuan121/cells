// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	Token
	RevokeTokenRequest
	RevokeTokenResponse
	PruneTokensRequest
	PruneTokensResponse
	ID
	GetLoginRequest
	GetLoginResponse
	CreateLoginRequest
	CreateLoginResponse
	AcceptLoginRequest
	AcceptLoginResponse
	GetConsentRequest
	GetConsentResponse
	CreateConsentRequest
	CreateConsentResponse
	AcceptConsentRequest
	AcceptConsentResponse
	CreateLogoutRequest
	CreateLogoutResponse
	AcceptLogoutRequest
	AcceptLogoutResponse
	CreateAuthCodeRequest
	CreateAuthCodeResponse
	VerifyTokenRequest
	VerifyTokenResponse
	ExchangeRequest
	ExchangeResponse
	PasswordCredentialsTokenRequest
	PasswordCredentialsTokenResponse
	RefreshTokenRequest
	RefreshTokenResponse
	PersonalAccessToken
	PatGenerateRequest
	PatGenerateResponse
	PatListRequest
	PatListResponse
	PatRevokeRequest
	PatRevokeResponse
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthTokenRevoker service

type AuthTokenRevokerClient interface {
	// Revoker invalidates the current token and specifies if the invalidation is due to a refresh or a revokation
	Revoke(ctx context.Context, in *RevokeTokenRequest, opts ...client.CallOption) (*RevokeTokenResponse, error)
}

type authTokenRevokerClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenRevokerClient(serviceName string, c client.Client) AuthTokenRevokerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenRevokerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenRevokerClient) Revoke(ctx context.Context, in *RevokeTokenRequest, opts ...client.CallOption) (*RevokeTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenRevoker.Revoke", in)
	out := new(RevokeTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenRevoker service

type AuthTokenRevokerHandler interface {
	// Revoker invalidates the current token and specifies if the invalidation is due to a refresh or a revokation
	Revoke(context.Context, *RevokeTokenRequest, *RevokeTokenResponse) error
}

func RegisterAuthTokenRevokerHandler(s server.Server, hdlr AuthTokenRevokerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenRevoker{hdlr}, opts...))
}

type AuthTokenRevoker struct {
	AuthTokenRevokerHandler
}

func (h *AuthTokenRevoker) Revoke(ctx context.Context, in *RevokeTokenRequest, out *RevokeTokenResponse) error {
	return h.AuthTokenRevokerHandler.Revoke(ctx, in, out)
}

// Client API for AuthTokenPruner service

type AuthTokenPrunerClient interface {
	// PruneTokens clear revoked tokens
	PruneTokens(ctx context.Context, in *PruneTokensRequest, opts ...client.CallOption) (*PruneTokensResponse, error)
}

type authTokenPrunerClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenPrunerClient(serviceName string, c client.Client) AuthTokenPrunerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenPrunerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenPrunerClient) PruneTokens(ctx context.Context, in *PruneTokensRequest, opts ...client.CallOption) (*PruneTokensResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenPruner.PruneTokens", in)
	out := new(PruneTokensResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenPruner service

type AuthTokenPrunerHandler interface {
	// PruneTokens clear revoked tokens
	PruneTokens(context.Context, *PruneTokensRequest, *PruneTokensResponse) error
}

func RegisterAuthTokenPrunerHandler(s server.Server, hdlr AuthTokenPrunerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenPruner{hdlr}, opts...))
}

type AuthTokenPruner struct {
	AuthTokenPrunerHandler
}

func (h *AuthTokenPruner) PruneTokens(ctx context.Context, in *PruneTokensRequest, out *PruneTokensResponse) error {
	return h.AuthTokenPrunerHandler.PruneTokens(ctx, in, out)
}

// Client API for LoginProvider service

type LoginProviderClient interface {
	GetLogin(ctx context.Context, in *GetLoginRequest, opts ...client.CallOption) (*GetLoginResponse, error)
	CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...client.CallOption) (*CreateLoginResponse, error)
	AcceptLogin(ctx context.Context, in *AcceptLoginRequest, opts ...client.CallOption) (*AcceptLoginResponse, error)
}

type loginProviderClient struct {
	c           client.Client
	serviceName string
}

func NewLoginProviderClient(serviceName string, c client.Client) LoginProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &loginProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *loginProviderClient) GetLogin(ctx context.Context, in *GetLoginRequest, opts ...client.CallOption) (*GetLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LoginProvider.GetLogin", in)
	out := new(GetLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginProviderClient) CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...client.CallOption) (*CreateLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LoginProvider.CreateLogin", in)
	out := new(CreateLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginProviderClient) AcceptLogin(ctx context.Context, in *AcceptLoginRequest, opts ...client.CallOption) (*AcceptLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LoginProvider.AcceptLogin", in)
	out := new(AcceptLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginProvider service

type LoginProviderHandler interface {
	GetLogin(context.Context, *GetLoginRequest, *GetLoginResponse) error
	CreateLogin(context.Context, *CreateLoginRequest, *CreateLoginResponse) error
	AcceptLogin(context.Context, *AcceptLoginRequest, *AcceptLoginResponse) error
}

func RegisterLoginProviderHandler(s server.Server, hdlr LoginProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&LoginProvider{hdlr}, opts...))
}

type LoginProvider struct {
	LoginProviderHandler
}

func (h *LoginProvider) GetLogin(ctx context.Context, in *GetLoginRequest, out *GetLoginResponse) error {
	return h.LoginProviderHandler.GetLogin(ctx, in, out)
}

func (h *LoginProvider) CreateLogin(ctx context.Context, in *CreateLoginRequest, out *CreateLoginResponse) error {
	return h.LoginProviderHandler.CreateLogin(ctx, in, out)
}

func (h *LoginProvider) AcceptLogin(ctx context.Context, in *AcceptLoginRequest, out *AcceptLoginResponse) error {
	return h.LoginProviderHandler.AcceptLogin(ctx, in, out)
}

// Client API for ConsentProvider service

type ConsentProviderClient interface {
	GetConsent(ctx context.Context, in *GetConsentRequest, opts ...client.CallOption) (*GetConsentResponse, error)
	CreateConsent(ctx context.Context, in *CreateConsentRequest, opts ...client.CallOption) (*CreateConsentResponse, error)
	AcceptConsent(ctx context.Context, in *AcceptConsentRequest, opts ...client.CallOption) (*AcceptConsentResponse, error)
}

type consentProviderClient struct {
	c           client.Client
	serviceName string
}

func NewConsentProviderClient(serviceName string, c client.Client) ConsentProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &consentProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *consentProviderClient) GetConsent(ctx context.Context, in *GetConsentRequest, opts ...client.CallOption) (*GetConsentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ConsentProvider.GetConsent", in)
	out := new(GetConsentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentProviderClient) CreateConsent(ctx context.Context, in *CreateConsentRequest, opts ...client.CallOption) (*CreateConsentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ConsentProvider.CreateConsent", in)
	out := new(CreateConsentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentProviderClient) AcceptConsent(ctx context.Context, in *AcceptConsentRequest, opts ...client.CallOption) (*AcceptConsentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ConsentProvider.AcceptConsent", in)
	out := new(AcceptConsentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsentProvider service

type ConsentProviderHandler interface {
	GetConsent(context.Context, *GetConsentRequest, *GetConsentResponse) error
	CreateConsent(context.Context, *CreateConsentRequest, *CreateConsentResponse) error
	AcceptConsent(context.Context, *AcceptConsentRequest, *AcceptConsentResponse) error
}

func RegisterConsentProviderHandler(s server.Server, hdlr ConsentProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ConsentProvider{hdlr}, opts...))
}

type ConsentProvider struct {
	ConsentProviderHandler
}

func (h *ConsentProvider) GetConsent(ctx context.Context, in *GetConsentRequest, out *GetConsentResponse) error {
	return h.ConsentProviderHandler.GetConsent(ctx, in, out)
}

func (h *ConsentProvider) CreateConsent(ctx context.Context, in *CreateConsentRequest, out *CreateConsentResponse) error {
	return h.ConsentProviderHandler.CreateConsent(ctx, in, out)
}

func (h *ConsentProvider) AcceptConsent(ctx context.Context, in *AcceptConsentRequest, out *AcceptConsentResponse) error {
	return h.ConsentProviderHandler.AcceptConsent(ctx, in, out)
}

// Client API for LogoutProvider service

type LogoutProviderClient interface {
	CreateLogout(ctx context.Context, in *CreateLogoutRequest, opts ...client.CallOption) (*CreateLogoutResponse, error)
	AcceptLogout(ctx context.Context, in *AcceptLogoutRequest, opts ...client.CallOption) (*AcceptLogoutResponse, error)
}

type logoutProviderClient struct {
	c           client.Client
	serviceName string
}

func NewLogoutProviderClient(serviceName string, c client.Client) LogoutProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &logoutProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *logoutProviderClient) CreateLogout(ctx context.Context, in *CreateLogoutRequest, opts ...client.CallOption) (*CreateLogoutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LogoutProvider.CreateLogout", in)
	out := new(CreateLogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logoutProviderClient) AcceptLogout(ctx context.Context, in *AcceptLogoutRequest, opts ...client.CallOption) (*AcceptLogoutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LogoutProvider.AcceptLogout", in)
	out := new(AcceptLogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogoutProvider service

type LogoutProviderHandler interface {
	CreateLogout(context.Context, *CreateLogoutRequest, *CreateLogoutResponse) error
	AcceptLogout(context.Context, *AcceptLogoutRequest, *AcceptLogoutResponse) error
}

func RegisterLogoutProviderHandler(s server.Server, hdlr LogoutProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&LogoutProvider{hdlr}, opts...))
}

type LogoutProvider struct {
	LogoutProviderHandler
}

func (h *LogoutProvider) CreateLogout(ctx context.Context, in *CreateLogoutRequest, out *CreateLogoutResponse) error {
	return h.LogoutProviderHandler.CreateLogout(ctx, in, out)
}

func (h *LogoutProvider) AcceptLogout(ctx context.Context, in *AcceptLogoutRequest, out *AcceptLogoutResponse) error {
	return h.LogoutProviderHandler.AcceptLogout(ctx, in, out)
}

// Client API for AuthCodeProvider service

type AuthCodeProviderClient interface {
	CreateAuthCode(ctx context.Context, in *CreateAuthCodeRequest, opts ...client.CallOption) (*CreateAuthCodeResponse, error)
}

type authCodeProviderClient struct {
	c           client.Client
	serviceName string
}

func NewAuthCodeProviderClient(serviceName string, c client.Client) AuthCodeProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authCodeProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authCodeProviderClient) CreateAuthCode(ctx context.Context, in *CreateAuthCodeRequest, opts ...client.CallOption) (*CreateAuthCodeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthCodeProvider.CreateAuthCode", in)
	out := new(CreateAuthCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthCodeProvider service

type AuthCodeProviderHandler interface {
	CreateAuthCode(context.Context, *CreateAuthCodeRequest, *CreateAuthCodeResponse) error
}

func RegisterAuthCodeProviderHandler(s server.Server, hdlr AuthCodeProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthCodeProvider{hdlr}, opts...))
}

type AuthCodeProvider struct {
	AuthCodeProviderHandler
}

func (h *AuthCodeProvider) CreateAuthCode(ctx context.Context, in *CreateAuthCodeRequest, out *CreateAuthCodeResponse) error {
	return h.AuthCodeProviderHandler.CreateAuthCode(ctx, in, out)
}

// Client API for AuthTokenVerifier service

type AuthTokenVerifierClient interface {
	// Verifies a token and returns claims
	Verify(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error)
}

type authTokenVerifierClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenVerifierClient(serviceName string, c client.Client) AuthTokenVerifierClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenVerifierClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenVerifierClient) Verify(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenVerifier.Verify", in)
	out := new(VerifyTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenVerifier service

type AuthTokenVerifierHandler interface {
	// Verifies a token and returns claims
	Verify(context.Context, *VerifyTokenRequest, *VerifyTokenResponse) error
}

func RegisterAuthTokenVerifierHandler(s server.Server, hdlr AuthTokenVerifierHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenVerifier{hdlr}, opts...))
}

type AuthTokenVerifier struct {
	AuthTokenVerifierHandler
}

func (h *AuthTokenVerifier) Verify(ctx context.Context, in *VerifyTokenRequest, out *VerifyTokenResponse) error {
	return h.AuthTokenVerifierHandler.Verify(ctx, in, out)
}

// Client API for AuthCodeExchanger service

type AuthCodeExchangerClient interface {
	Exchange(ctx context.Context, in *ExchangeRequest, opts ...client.CallOption) (*ExchangeResponse, error)
}

type authCodeExchangerClient struct {
	c           client.Client
	serviceName string
}

func NewAuthCodeExchangerClient(serviceName string, c client.Client) AuthCodeExchangerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authCodeExchangerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authCodeExchangerClient) Exchange(ctx context.Context, in *ExchangeRequest, opts ...client.CallOption) (*ExchangeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthCodeExchanger.Exchange", in)
	out := new(ExchangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthCodeExchanger service

type AuthCodeExchangerHandler interface {
	Exchange(context.Context, *ExchangeRequest, *ExchangeResponse) error
}

func RegisterAuthCodeExchangerHandler(s server.Server, hdlr AuthCodeExchangerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthCodeExchanger{hdlr}, opts...))
}

type AuthCodeExchanger struct {
	AuthCodeExchangerHandler
}

func (h *AuthCodeExchanger) Exchange(ctx context.Context, in *ExchangeRequest, out *ExchangeResponse) error {
	return h.AuthCodeExchangerHandler.Exchange(ctx, in, out)
}

// Client API for PasswordCredentialsToken service

type PasswordCredentialsTokenClient interface {
	PasswordCredentialsToken(ctx context.Context, in *PasswordCredentialsTokenRequest, opts ...client.CallOption) (*PasswordCredentialsTokenResponse, error)
}

type passwordCredentialsTokenClient struct {
	c           client.Client
	serviceName string
}

func NewPasswordCredentialsTokenClient(serviceName string, c client.Client) PasswordCredentialsTokenClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &passwordCredentialsTokenClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *passwordCredentialsTokenClient) PasswordCredentialsToken(ctx context.Context, in *PasswordCredentialsTokenRequest, opts ...client.CallOption) (*PasswordCredentialsTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PasswordCredentialsToken.PasswordCredentialsToken", in)
	out := new(PasswordCredentialsTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PasswordCredentialsToken service

type PasswordCredentialsTokenHandler interface {
	PasswordCredentialsToken(context.Context, *PasswordCredentialsTokenRequest, *PasswordCredentialsTokenResponse) error
}

func RegisterPasswordCredentialsTokenHandler(s server.Server, hdlr PasswordCredentialsTokenHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PasswordCredentialsToken{hdlr}, opts...))
}

type PasswordCredentialsToken struct {
	PasswordCredentialsTokenHandler
}

func (h *PasswordCredentialsToken) PasswordCredentialsToken(ctx context.Context, in *PasswordCredentialsTokenRequest, out *PasswordCredentialsTokenResponse) error {
	return h.PasswordCredentialsTokenHandler.PasswordCredentialsToken(ctx, in, out)
}

// Client API for AuthTokenRefresher service

type AuthTokenRefresherClient interface {
	Refresh(ctx context.Context, in *RefreshTokenRequest, opts ...client.CallOption) (*RefreshTokenResponse, error)
}

type authTokenRefresherClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenRefresherClient(serviceName string, c client.Client) AuthTokenRefresherClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenRefresherClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenRefresherClient) Refresh(ctx context.Context, in *RefreshTokenRequest, opts ...client.CallOption) (*RefreshTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenRefresher.Refresh", in)
	out := new(RefreshTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenRefresher service

type AuthTokenRefresherHandler interface {
	Refresh(context.Context, *RefreshTokenRequest, *RefreshTokenResponse) error
}

func RegisterAuthTokenRefresherHandler(s server.Server, hdlr AuthTokenRefresherHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenRefresher{hdlr}, opts...))
}

type AuthTokenRefresher struct {
	AuthTokenRefresherHandler
}

func (h *AuthTokenRefresher) Refresh(ctx context.Context, in *RefreshTokenRequest, out *RefreshTokenResponse) error {
	return h.AuthTokenRefresherHandler.Refresh(ctx, in, out)
}

// Client API for PersonalAccessTokenService service

type PersonalAccessTokenServiceClient interface {
	Generate(ctx context.Context, in *PatGenerateRequest, opts ...client.CallOption) (*PatGenerateResponse, error)
	Revoke(ctx context.Context, in *PatRevokeRequest, opts ...client.CallOption) (*PatRevokeResponse, error)
	List(ctx context.Context, in *PatListRequest, opts ...client.CallOption) (*PatListResponse, error)
}

type personalAccessTokenServiceClient struct {
	c           client.Client
	serviceName string
}

func NewPersonalAccessTokenServiceClient(serviceName string, c client.Client) PersonalAccessTokenServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &personalAccessTokenServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *personalAccessTokenServiceClient) Generate(ctx context.Context, in *PatGenerateRequest, opts ...client.CallOption) (*PatGenerateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PersonalAccessTokenService.Generate", in)
	out := new(PatGenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personalAccessTokenServiceClient) Revoke(ctx context.Context, in *PatRevokeRequest, opts ...client.CallOption) (*PatRevokeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PersonalAccessTokenService.Revoke", in)
	out := new(PatRevokeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personalAccessTokenServiceClient) List(ctx context.Context, in *PatListRequest, opts ...client.CallOption) (*PatListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PersonalAccessTokenService.List", in)
	out := new(PatListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonalAccessTokenService service

type PersonalAccessTokenServiceHandler interface {
	Generate(context.Context, *PatGenerateRequest, *PatGenerateResponse) error
	Revoke(context.Context, *PatRevokeRequest, *PatRevokeResponse) error
	List(context.Context, *PatListRequest, *PatListResponse) error
}

func RegisterPersonalAccessTokenServiceHandler(s server.Server, hdlr PersonalAccessTokenServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PersonalAccessTokenService{hdlr}, opts...))
}

type PersonalAccessTokenService struct {
	PersonalAccessTokenServiceHandler
}

func (h *PersonalAccessTokenService) Generate(ctx context.Context, in *PatGenerateRequest, out *PatGenerateResponse) error {
	return h.PersonalAccessTokenServiceHandler.Generate(ctx, in, out)
}

func (h *PersonalAccessTokenService) Revoke(ctx context.Context, in *PatRevokeRequest, out *PatRevokeResponse) error {
	return h.PersonalAccessTokenServiceHandler.Revoke(ctx, in, out)
}

func (h *PersonalAccessTokenService) List(ctx context.Context, in *PatListRequest, out *PatListResponse) error {
	return h.PersonalAccessTokenServiceHandler.List(ctx, in, out)
}
