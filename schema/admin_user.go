package schema

import (
	"github.com/gofrs/uuid"
	"time"
)

type GoTrueIdentity map[string]interface{}

type User struct {
	ID uuid.UUID `json:"id"`

	Aud   string `json:"aud"`
	Role  string `json:"role"`
	Email string `json:"email"`

	EmailConfirmedAt *time.Time `json:"email_confirmed_at,omitempty"`
	InvitedAt        *time.Time `json:"invited_at,omitempty"`

	Phone            string     `json:"phone"`
	PhoneConfirmedAt *time.Time `json:"phone_confirmed_at,omitempty"`

	ConfirmationSentAt *time.Time `json:"confirmation_sent_at,omitempty"`

	// For backward compatibility only. Use EmailConfirmedAt or PhoneConfirmedAt instead.
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`

	RecoverySentAt *time.Time `json:"recovery_sent_at,omitempty"`

	EmailChange       string     `json:"new_email,omitempty"`
	EmailChangeSentAt *time.Time `json:"email_change_sent_at,omitempty"`

	PhoneChange       string     `json:"new_phone,omitempty"`
	PhoneChangeSentAt *time.Time `json:"phone_change_sent_at,omitempty"`

	ReauthenticationSentAt *time.Time `json:"reauthentication_sent_at,omitempty"`

	LastSignInAt *time.Time `json:"last_sign_in_at,omitempty"`

	AppMetaData  map[string]interface{} `json:"app_metadata"`
	UserMetaData map[string]interface{} `json:"user_metadata"`

	Factors    map[string]interface{} `json:"factors,omitempty"`
	Identities *[]GoTrueIdentity      `json:"identities,omitempty"`

	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	BannedUntil *time.Time `json:"banned_until,omitempty"`
}
