package operations

import (
	"github.com/LukeHagar/plexgo/models/components"
	"net/http"
)

var GetHomeUsersServerList = []string{
	"https://plex.tv/api/v2",
}

type GetHomeUsersGlobals struct {
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

type GetHomeUsersRequest struct {
}

type HomeUserSubscription struct {
	State string `json:"state"`
	Type  string `json:"type"`
}

type HomeUser struct {
	ID                 int64                 `json:"id"`
	UUID               string                `json:"uuid"`
	Title              string                `json:"title"`
	Username           string                `json:"username"`
	Email              string                `json:"email"`
	FriendlyName       string                `json:"friendlyName"`
	Thumb              string                `json:"thumb"`
	HasPassword        bool                  `json:"hasPassword"`
	Restricted         bool                  `json:"restricted"`
	UpdatedAt          int64                 `json:"updatedAt"`
	RestrictionProfile *string               `json:"restrictionProfile"`
	Admin              bool                  `json:"admin"`
	Guest              bool                  `json:"guest"`
	Protected          bool                  `json:"protected"`
	Pin                *string               `json:"pin"`
	Subscription       HomeUserSubscription `json:"subscription"`
}

type GetHomeUsersResponseBody struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	GuestUserID   int64      `json:"guestUserID"`
	GuestUserUUID string     `json:"guestUserUUID"`
	GuestEnabled  bool       `json:"guestEnabled"`
	Subscription  bool       `json:"subscription"`
	Users         []HomeUser `json:"users"`
}

type GetHomeUsersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Object *GetHomeUsersResponseBody
}
