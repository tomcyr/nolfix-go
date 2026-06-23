package msg

import "encoding/xml"

// UserRequest is the login/logout/status request (<UserReq>).
type UserRequest struct {
	XMLName     xml.Name `xml:"UserReq"`
	UserReqID   string   `xml:"UserReqID,attr,omitempty"`
	UserReqTyp  *int     `xml:"UserReqTyp,attr,omitempty"`
	Username    string   `xml:"Username,attr,omitempty"`
	Password    string   `xml:"Password,attr,omitempty"`
	NewPassword string   `xml:"NewPassword,attr,omitempty"`
}

func (u UserRequest) MsgName() string { return "UserReq" }

// UserRequestType values for UserRequest.UserReqTyp.
const (
	UserReqTypLogin          = 1
	UserReqTypLogout         = 2
	UserReqTypPasswordChange = 3
	UserReqTypUserStatus     = 4
)

// UserResponse is the response to a UserRequest (<UserRsp>).
type UserResponse struct {
	XMLName      xml.Name `xml:"UserRsp"`
	UserReqID    string   `xml:"UserReqID,attr,omitempty"`
	Username     string   `xml:"Username,attr,omitempty"`
	UserStat     *int     `xml:"UserStat,attr,omitempty"`
	UserStatText string   `xml:"UserStatText,attr,omitempty"`
	MktDepth     *int     `xml:"MktDepth,attr,omitempty"`
}

func (u UserResponse) MsgName() string { return "UserRsp" }

// UserStatus values for UserResponse.UserStat.
const (
	UserStatusLoggedIn      = 1
	UserStatusLoggedOut     = 2
	UserStatusNotExists     = 3
	UserStatusWrongPasswd   = 4
	UserStatusPasswdChanged = 5
	UserStatusOther         = 6
	UserStatusStockLogout   = 7
	UserStatusSessionClosed = 8
)

// UserStatText values for UserResponse.UserStatText.
const (
	UserStatTextNol3Closed   = 1
	UserStatTextNol3Offline  = 2
	UserStatTextNol3Online   = 3
	UserStatTextNol3Disabled = 4
)
