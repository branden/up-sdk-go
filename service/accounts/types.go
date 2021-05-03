package accounts

import "time"

// Account is an Upbound Cloud account.
type Account struct {
	Name string `json:"name,omitempty"`
	Type Type   `json:"type,omitempty"`
}

// Type is either a user or organization.
type Type string

const (
	// AccountOrganization is an organization account.
	AccountOrganization Type = "organization"
	// AccountUser is a user account.
	AccountUser Type = "user"
)

// User is a user on Upbound Cloud.
// TODO(hasheddan): move to user service when implemented.
type User struct {
	ID              uint       `json:"id,omitempty"`
	Username        string     `json:"username,omitempty"`
	FirstName       string     `json:"firstName,omitempty"`
	LastName        string     `json:"lastName,omitempty"`
	Email           string     `json:"email,omitempty"`
	Biography       string     `json:"biography,omitempty"`
	Location        string     `json:"location,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	EnterpriseTrial *time.Time `json:"enterpriseTrial,omitempty"`
	PersonalTrial   *time.Time `json:"personalTrial,omitempty"`
}

// Organization is an organization on Upbound Cloud.
// TODO(hasheddan): move to org service when implemented.
type Organization struct {
	ID                   uint       `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	DisplayName          string     `json:"displayName,omitempty"`
	CreatorID            uint       `json:"creatorId,omitempty"`
	ReservedEnvironments int        `json:"reservedEnvironments"`
	CreatedAt            *time.Time `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time `json:"updatedAt,omitempty"`
	DeletedAt            *time.Time `json:"deletedAt,omitempty"`
	ExpiresAt            *time.Time `json:"expiresAt,omitempty"`
	DeleteAt             *time.Time `json:"deleteAt,omitempty"`
}

// OrganizationPermissionGroup is the type of permission a user has in the
// organization.
type OrganizationPermissionGroup string

const (
	// OrganizationMember denotes basic permission on an organization.
	OrganizationMember OrganizationPermissionGroup = "member"
	// OrganizationOwner denotes full access permission on an organization.
	OrganizationOwner OrganizationPermissionGroup = "owner"
)

// AccountResponse is the API response when requesting information on a
// account.
type AccountResponse struct {
	Account      Account       `json:"account"`
	Organization *Organization `json:"organization,omitempty"`
	User         *User         `json:"user,omitempty"`
}