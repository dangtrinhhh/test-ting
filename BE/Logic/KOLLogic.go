package Logic

import (
	"log"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex int, pageSize int) ([]*DTO.KolDTO, error) {
	var kols []*DTO.KolDTO

	// Calculate offset to split page if necessary
	offset := (pageIndex - 1) * pageSize

	// Get data by query
	query := `
		SELECT "KolID", "UserProfileID", "Language", "Education", "ExpectedSalary", "ExpectedSalaryEnable", "ChannelSettingTypeID", 
			   "IDFrontURL", "IDBackURL", "PortraitURL", "RewardID", "PaymentMethodID", "TestimonialsID", "VerificationStatus", 
			   "Enabled", "ActiveDate", "Active", "CreatedBy", "CreatedDate", "ModifiedBy", "ModifiedDate", "IsRemove", "IsOnBoarding", 
			   "Code", "PortraitRightURL", "PortraitLeftURL", "LivenessStatus"
		FROM "KOL" LIMIT ? OFFSET ?`

	err := Initializers.DB.Raw(query, pageSize, offset).Scan(&kols).Error

	if err != nil {
		log.Printf("Error when query data from KOL table: %v", err)
		return nil, err
	}

	// Return list KOLs và no error
	return kols, nil
}
