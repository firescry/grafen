package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Instance struct {
	Uri              string        `json:"uri"`
	Title            string        `json:"title"`
	ShortDescription string        `json:"short_description"`
	Description      string        `json:"description"`
	Email            string        `json:"email"`
	Version          string        `json:"version"`
	Urls             Urls          `json:"urls"`
	Stats            Stats         `json:"stats"`
	Thumbnail        string        `json:"thumbnail"`
	Languages        Languages     `json:"languages"`
	Registrations    bool          `json:"registrations"`
	ApprovalRequired bool          `json:"approval_required"`
	InvitesEnabled   bool          `json:"invites_enabled"`
	Configuration    Configuration `json:"configuration"`
	ContactAccount   Contact       `json:"contact_account"`
	Rules            Rules         `json:"rules"`
}

type Urls struct {
	StreamingApi string `json:"streaming_api"`
}

type Stats struct {
	UserCount   int `json:"user_count"`
	StatusCount int `json:"status_count"`
	DomainCount int `json:"domain_count"`
}

type Languages []string

type Configuration struct {
	Accounts         Accounts         `json:"accounts"`
	Statuses         Statuses         `json:"statuses"`
	MediaAttachments MediaAttachments `json:"media_attachments"`
	Polls            Polls            `json:"polls"`
}

type Accounts struct {
	MaxFeaturedTags int `json:"max_featured_tags"`
}

type Statuses struct {
	MaxCharacters            int `json:"max_characters"`
	MaxMediaAttachments      int `json:"max_media_attachments"`
	CharactersReservedPerUrl int `json:"characters_reserved_per_url"`
}

type MediaAttachments struct {
	SupportedMimeTypes  []string `json:"supported_mime_types"`
	ImageSizeLimit      int      `json:"image_size_limit"`
	ImageMatrixLimit    int      `json:"image_matrix_limit"`
	VideoSizeLimit      int      `json:"video_size_limit"`
	VideoFrameRateLimit int      `json:"video_frame_rate_limit"`
	VideoMatrixLimit    int      `json:"video_matrix_limit"`
}

type Polls struct {
	MaxOptions             int `json:"max_options"`
	MaxCharactersPerOption int `json:"max_characters_per_option"`
	MinExpiration          int `json:"min_expiration"`
	MaxExpiration          int `json:"max_expiration"`
}

type Contact struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Acct           string `json:"acct"`
	DisplayName    string `json:"display_name"`
	Locked         bool   `json:"locked"`
	Bot            bool   `json:"bot"`
	Discoverable   bool   `json:"discoverable"`
	Group          bool   `json:"group"`
	CreatedAt      string `json:"created_at"`
	Note           string `json:"note"`
	Url            string `json:"url"`
	Avatar         string `json:"avatar"`
	AvatarStatic   string `json:"avatar_static"`
	Header         string `json:"header"`
	HeaderStatic   string `json:"header_static"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	StatusesCount  int    `json:"statuses_count"`
	LastStatusAt   string `json:"last_status_at"`
	Emojis         Emojis `json:"emojis"`
	Fields         Fields `json:"fields"`
}

type Emojis []string

type Fields []Field

type Field struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	VerifiedAt int    `json:"verified_at"`
}

func GetInstance(c *gin.Context) {
	urls := Urls{}
	stats := Stats{}
	languages := Languages{}
	configuration := Configuration{
		MediaAttachments: MediaAttachments{
			SupportedMimeTypes: []string{},
		},
	}
	contact := Contact{
		Emojis: Emojis{},
		Fields: Fields{Field{}},
	}
	rules := Rules{Rule{}}
	instance := Instance{
		Urls:           urls,
		Stats:          stats,
		Languages:      languages,
		Configuration:  configuration,
		ContactAccount: contact,
		Rules:          rules,
	}
	c.JSON(http.StatusOK, instance)
}
