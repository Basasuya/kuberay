package utils

import (
	"fmt"
	"net/http"
	"time"
)

var (
	DefaultHttpProxyPort = 8000
	HealthCheckPath      = "/-/healthz"
)

type RayHttpProxyClientInterface interface {
	InitClient()
	CheckHealth() error
	SetHostIp(hostIp string)
}

// GetRayHttpProxyClientFunc Used for unit tests.
var GetRayHttpProxyClientFunc = GetRayHttpProxyClient

func GetRayHttpProxyClient() RayHttpProxyClientInterface {
	return &RayHttpProxyClient{}
}

type RayHttpProxyClient struct {
	client       http.Client
	httpProxyURL string
}

func (r *RayHttpProxyClient) InitClient() {
	r.client = http.Client{
		Timeout: 20 * time.Millisecond,
	}
}

func (r *RayHttpProxyClient) SetHostIp(hostIp string) {
	r.httpProxyURL = fmt.Sprint("http://", hostIp, ":", DefaultHttpProxyPort)
}

func (r *RayHttpProxyClient) CheckHealth() error {
	req, err := http.NewRequest("GET", r.httpProxyURL+HealthCheckPath, nil)
	if err != nil {
		return err
	}

	resp, err := r.client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
