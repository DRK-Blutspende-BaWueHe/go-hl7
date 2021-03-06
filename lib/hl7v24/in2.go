package hl7v24

import "time"

// HL7 v2.4 - IN2 - Insurance Additional Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN2
type IN2 struct {
	InsuredEmployeeID                    []CX        `hl7:"1" json:"InsuredEmployeeID,omitempty"`
	InsuredSSN                           string      `hl7:"2" json:"InsuredSSN,omitempty"`
	InsuredEmployerName                  []XCN       `hl7:"3" json:"InsuredEmployerName,omitempty"`
	EmployerInformationData              string      `hl7:"4" json:"EmployerInformationData,omitempty"`
	MailClaimParty                       []string    `hl7:"5" json:"MailClaimParty,omitempty"`
	MedicareHealthInsCardNumber          string      `hl7:"6" json:"MedicareHealthInsCardNumber,omitempty"`
	MedicaidCaseName                     []XPN       `hl7:"7" json:"MedicaidCaseName,omitempty"`
	MedicaidCaseNumber                   string      `hl7:"8" json:"MedicaidCaseNumber,omitempty"`
	MilitarySponsorName                  []XPN       `hl7:"9" json:"MilitarySponsorName,omitempty"`
	MilitaryIDNumber                     string      `hl7:"10" json:"MilitaryIDNumber,omitempty"`
	DependentOfMilitaryRecipient         CE          `hl7:"11" json:"DependentOfMilitaryRecipient,omitempty"`
	MilitaryOrganization                 string      `hl7:"12" json:"MilitaryOrganization,omitempty"`
	MilitaryStation                      string      `hl7:"13" json:"MilitaryStation,omitempty"`
	MilitaryService                      string      `hl7:"14" json:"MilitaryService,omitempty"`
	MilitaryRank                         string      `hl7:"15" json:"MilitaryRank,omitempty"`
	MilitaryStatus                       string      `hl7:"16" json:"MilitaryStatus,omitempty"`
	MilitaryRetireDate                   time.Time   `hl7:"17,shortdate" json:"MilitaryRetireDate,omitempty"`
	MilitaryNonAvailCertOnFile           string      `hl7:"18" json:"MilitaryNonAvailCertOnFile,omitempty"`
	BabyCoverage                         string      `hl7:"19" json:"BabyCoverage,omitempty"`
	CombineBabyBill                      string      `hl7:"20" json:"CombineBabyBill,omitempty"`
	BloodDeductible                      string      `hl7:"21" json:"BloodDeductible,omitempty"`
	SpecialCoverageApprovalName          []XPN       `hl7:"22" json:"SpecialCoverageApprovalName,omitempty"`
	SpecialCoverageApprovalTitle         string      `hl7:"23" json:"SpecialCoverageApprovalTitle,omitempty"`
	NonCoveredInsuranceCode              []string    `hl7:"24" json:"NonCoveredInsuranceCode,omitempty"`
	PayorID                              []CX        `hl7:"25" json:"PayorID,omitempty"`
	PayorSubscriberID                    []CX        `hl7:"26" json:"PayorSubscriberID,omitempty"`
	EligibilitySource                    int         `hl7:"27" json:"EligibilitySource,omitempty"`
	RoomCoverageTypeAmount               []RMC       `hl7:"28" json:"RoomCoverageTypeAmount,omitempty"`
	PolicyTypeAmount                     []PTA       `hl7:"29" json:"PolicyTypeAmount,omitempty"`
	DailyDeductible                      DDI         `hl7:"30" json:"DailyDeductible,omitempty"`
	LivingDependency                     string      `hl7:"31" json:"LivingDependency,omitempty"`
	AmbulatoryStatus                     []string    `hl7:"32" json:"AmbulatoryStatus,omitempty"`
	Citizenship                          []CE        `hl7:"33" json:"Citizenship,omitempty"`
	PrimaryLanguage                      CE          `hl7:"34" json:"PrimaryLanguage,omitempty"`
	LivingArrangement                    string      `hl7:"35" json:"LivingArrangement,omitempty"`
	PublicityIndicator                   CE          `hl7:"36" json:"PublicityIndicator,omitempty"`
	ProtectionIndicator                  string      `hl7:"37" json:"ProtectionIndicator,omitempty"`
	StudentIndicator                     string      `hl7:"38" json:"StudentIndicator,omitempty"`
	Religion                             CE          `hl7:"39" json:"Religion,omitempty"`
	MothersMaidenName                    []XPN       `hl7:"40" json:"MothersMaidenName,omitempty"`
	Nationality                          CE          `hl7:"41" json:"Nationality,omitempty"`
	EthnicGroup                          []CE        `hl7:"42" json:"EthnicGroup,omitempty"`
	MaritalStatus                        []CE        `hl7:"43" json:"MaritalStatus,omitempty"`
	InsuredEmploymentStartDate           time.Time   `hl7:"44,shortdate" json:"InsuredEmploymentStartDate,omitempty"`
	EmploymentStopDate                   time.Time   `hl7:"45,shortdate" json:"EmploymentStopDate,omitempty"`
	JobTitle                             string      `hl7:"46" json:"JobTitle,omitempty"`
	JobCode                              JCC         `hl7:"47" json:"JobCode,omitempty"`
	JobStatus                            string      `hl7:"48" json:"JobStatus,omitempty"`
	EmployerContactPersonName            []XPN       `hl7:"49" json:"EmployerContactPersonName,omitempty"`
	EmployerContactPersonPhone           []XTN       `hl7:"50" json:"EmployerContactPersonPhone,omitempty"`
	EmployerContactReason                string      `hl7:"51" json:"EmployerContactReason,omitempty"`
	InsuredContactPersonName             []XPN       `hl7:"52" json:"InsuredContactPersonName,omitempty"`
	InsuredContactPersonTelephone        []XTN       `hl7:"53" json:"InsuredContactPersonTelephone,omitempty"`
	InsuredContactPersonReason           []string    `hl7:"54" json:"InsuredContactPersonReason,omitempty"`
	RelationshipToThePatientStartDate    time.Time   `hl7:"55,shortdate" json:"RelationshipToThePatientStartDate,omitempty"`
	RelationshipToThePatientStopDate     []time.Time `hl7:"56,shortdate" json:"RelationshipToThePatientStopDate,omitempty"`
	InsuranceCompanyContactReason        string      `hl7:"57" json:"InsuranceCompanyContactReason,omitempty"`
	InsuranceCompanyContactPhone         XTN         `hl7:"58" json:"InsuranceCompanyContactPhone,omitempty"`
	PolicyScope                          string      `hl7:"59" json:"PolicyScope,omitempty"`
	PolicySource                         string      `hl7:"60" json:"PolicySource,omitempty"`
	PatientMemberNumber                  CX          `hl7:"61" json:"PatientMemberNumber,omitempty"`
	GuarantorRelationshipToInsured       CE          `hl7:"62" json:"GuarantorRelationshipToInsured,omitempty"`
	InsuredsTelephoneNumberHome          []XTN       `hl7:"63" json:"InsuredsTelephoneNumberHome,omitempty"`
	InsuredEmployerTelephoneNumber       []XTN       `hl7:"64" json:"InsuredEmployerTelephoneNumber,omitempty"`
	MilitaryHandicappedProgram           CE          `hl7:"65" json:"MilitaryHandicappedProgram,omitempty"`
	SuspendFlag                          string      `hl7:"66" json:"SuspendFlag,omitempty"`
	CopayLimitFlag                       string      `hl7:"67" json:"CopayLimitFlag,omitempty"`
	StoplossLimitFlag                    string      `hl7:"68" json:"StoplossLimitFlag,omitempty"`
	InsuredOrganizationNameAndID         []XON       `hl7:"69" json:"InsuredOrganizationNameAndID,omitempty"`
	InsuredEmployerOrganizationNameAndID []XON       `hl7:"70" json:"InsuredEmployerOrganizationNameAndID,omitempty"`
	Race                                 []CE        `hl7:"71" json:"Race,omitempty"`
	HCFAPatientRelationshipToInsured     CE          `hl7:"72" json:"HCFAPatientRelationshipToInsured,omitempty"`
}
