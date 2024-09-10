package company

import (
	"upsider-coding-test/shared"

	"github.com/google/uuid"
)

type (
	CompanyID string
	PartnerID string
)

func NewCompanyID() CompanyID {
	id := uuid.New().String()
	return CompanyID(id)
}
func ParseCompanyID(id string) (CompanyID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", &shared.ArgumentError{Field: "companyID", Err: "invalid company id"}
	}
	return CompanyID(id), nil
}
func (id CompanyID) String() string {
	return string(id)
}

func NewPartnerID() PartnerID {
	id := uuid.New().String()
	return PartnerID(id)
}
func ParsePartnerID(id string) (PartnerID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return "", &shared.ArgumentError{Field: "partnerID", Err: "invalid partner id"}
	}
	return PartnerID(id), nil
}
func (id PartnerID) String() string {
	return string(id)
}
