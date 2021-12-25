package players

type InternalProfile struct {
	DefaultID  int  `json:"defaultId"`
	VerifiedID int  `json:"verifiedId"`
	Verified   bool `json:"verified"`
}

// Converts the internal profile to an external profile.
func (p *InternalProfile) Export() ExternalProfile {
	var externalProfile ExternalProfile
	externalProfile.ID = p.DefaultID
	externalProfile.Verified = p.Verified
	if p.Verified {
		externalProfile.ID = p.VerifiedID
	}
	return externalProfile
}

// ExternalProfile represents a user's profile which is safe to share with other packages.
type ExternalProfile struct {
	Verified bool
	ID       int
}
