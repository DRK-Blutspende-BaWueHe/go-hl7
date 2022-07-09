package hl7v23

// ODS - Dietary orders, supplements, and preferences
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ODS
type ODS struct {
	Type                           string   `hl7:"1"`
	ServicePeriod                  []CE     `hl7:"2"`
	DietSupplementOrPreferenceCode []CE     `hl7:"3"`
	TextInstruction                []string `hl7:"4"`
}
