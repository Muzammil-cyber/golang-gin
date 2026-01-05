package graph

import (
	"github.com/muzammil-cyber/golang-gin/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	VideoService service.VideoService
	JWTService   service.JWTService
}
