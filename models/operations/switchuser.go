package operations

import (
	"github.com/LukeHagar/plexgo/models/components"
	"net/http"
)

var SwitchUserServerList = []string{
	"https://plex.tv/api/v2",
}

type SwitchUserGlobals struct {
	Accepts          *components.Accepts `default:"application/json" header:"style=simple,explode=false,name=accepts"`
	ClientIdentifier *string             `header:"style=simple,explode=false,name=X-Plex-Client-Identifier"`
	Product          *string             `header:"style=simple,explode=false,name=X-Plex-Product"`
	Version          *string             `header:"style=simple,explode=false,name=X-Plex-Version"`
	Platform         *string             `header:"style=simple,explode=false,name=X-Plex-Platform"`
	PlatformVersion  *string             `header:"style=simple,explode=false,name=X-Plex-Platform-Version"`
	Device           *string             `header:"style=simple,explode=false,name=X-Plex-Device"`
	Model            *string             `header:"style=simple,explode=false,name=X-Plex-Model"`
	DeviceVendor     *string             `header:"style=simple,explode=false,name=X-Plex-Device-Vendor"`
	DeviceName       *string             `header:"style=simple,explode=false,name=X-Plex-Device-Name"`
	Marketplace      *string             `header:"style=simple,explode=false,name=X-Plex-Marketplace"`
}

type SwitchUserRequestBody struct {
	Pin *string `json:"pin,omitempty"`
}

type SwitchUserRequest struct {
	ID          string                `pathParam:"style=simple,explode=false,name=id"`
	RequestBody SwitchUserRequestBody `request:"mediaType=application/json"`
}

type SwitchUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UserPlexAccount *components.UserPlexAccount
}