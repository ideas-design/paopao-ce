// Code generated by go-mir. DO NOT EDIT.

package v1

import (
	"errors"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

type WebCore interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Index(c *gin.Context) error
	Articles(c *gin.Context) error
	Login(c *gin.Context, req *LoginReq) (*LoginResp, error)
	Logout(c *gin.Context) error

	mustEmbedUnimplementedWebCoreServant()
}

type WebCoreBinding interface {
	BindLogin(c *gin.Context) (*LoginReq, error)

	mustEmbedUnimplementedWebCoreBinding()
}

type WebCoreRender interface {
	RenderIndex(c *gin.Context, err error)
	RenderArticles(c *gin.Context, err error)
	RenderLogin(c *gin.Context, data *LoginResp, err error)
	RenderLogout(c *gin.Context, err error)

	mustEmbedUnimplementedWebCoreRender()
}

// RegisterWebCoreServant register WebCore servant to gin
func RegisterWebCoreServant(e *gin.Engine, s WebCore, b WebCoreBinding, r WebCoreRender) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/index/", func(c *gin.Context) {
		r.RenderIndex(c, s.Index(c))
	})
	router.Handle("GET", "/articles/:category/", func(c *gin.Context) {
		r.RenderArticles(c, s.Articles(c))
	})
	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
		}
		resp, err := s.Login(c, req)
		r.RenderLogin(c, resp, err)
	})
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		r.RenderLogout(c, s.Logout(c))
	})
}

// UnimplementedWebCoreServant can be embedded to have forward compatible implementations.
type UnimplementedWebCoreServant struct{}

// UnimplementedWebCoreBinding can be embedded to have forward compatible implementations.
type UnimplementedWebCoreBinding struct{}

// UnimplementedWebCoreRender can be embedded to have forward compatible implementations.
type UnimplementedWebCoreRender struct{}

func (UnimplementedWebCoreServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedWebCoreServant) Index(c *gin.Context) error {
	return errors.New("method Index not implemented")
}

func (UnimplementedWebCoreServant) Articles(c *gin.Context) error {
	return errors.New("method Index not implemented")
}

func (UnimplementedWebCoreServant) Login(c *gin.Context, req *LoginReq) (*LoginResp, error) {
	return nil, errors.New("method Login not implemented")
}

func (UnimplementedWebCoreServant) Logout(c *gin.Context) error {
	return errors.New("method Logout not implemented")
}

func (UnimplementedWebCoreServant) mustEmbedUnimplementedWebCoreServant() {}

func (UnimplementedWebCoreBinding) BindLogin(c *gin.Context) (*LoginReq, error) {
	return nil, errors.New("method BindLogin not implemented")
}

func (UnimplementedWebCoreBinding) mustEmbedUnimplementedWebCoreBinding() {}

func (UnimplementedWebCoreRender) RenderIndex(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (UnimplementedWebCoreRender) RenderArticles(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (UnimplementedWebCoreRender) RenderLogin(c *gin.Context, data *LoginResp, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogin not implemented")
}

func (UnimplementedWebCoreRender) RenderLogout(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "method RenderLogout not implemented")
}

func (UnimplementedWebCoreRender) mustEmbedUnimplementedWebCoreRender() {}
