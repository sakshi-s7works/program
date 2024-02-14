// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex;not null"`
    FullName  string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}


type UserProfile struct {
	gorm.Model
	UserID                  uint   `json:"user_id"`
	ProfileImageURL         string `json:"profile_image_url"`
	Headline                string `json:"headline"`
	PhoneNumber             string `json:"phone_number"`
	ShortBio                string `json:"short_bio" gorm:"type:text"`
	CurrentCityID           uint   `json:"current_city_id"`
	LinkedInUsername        string `json:"linkedin_username"`
	CompanyURL              string `json:"company_url"`
	DietaryRestrictions     string `json:"dietary_restrictions"`
	Designation             string `json:"designation"`
	JoinSpecificGroup       bool   `json:"join_specific_group"`
	SpecificGroupCriteria   string `json:"specific_group_criteria"`
	Gender                  string `json:"gender"`
	Ethnicity               string `json:"ethnicity"`
	AgeRange                string `json:"age_range"`
	InterestInSpecialEvents string `json:"interest_in_special_events"`
	ReferredBy              string `json:"referred_by"`
	CreatedAt               time.Time
}