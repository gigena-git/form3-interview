package test

import "github.com/gigena-git/form3-interview/pkg/create"

func generateCreateData() *create.AccountDataRequest {
	country := "GB"
	version := int64(0)
	account_classification := "Personal"
	joint_account := false
	account_matching_opt_out := false
	attributes := create.AccountAttributes{
		Country:                 &country,
		BaseCurrency:            "GBP",
		BankID:                  "6011797507804003",
		BankIDCode:              "XWZULDKWYPJCYMAQ",
		Bic:                     "IKSYJY95",
		Name:                    []string{"Lady Misty Emard"},
		AlternativeNames:        []string{"XRIEESSWVW"},
		AccountClassification:   &account_classification,
		JointAccount:            &joint_account,
		AccountMatchingOptOut:   &account_matching_opt_out,
		SecondaryIdentification: "SSQRGELQGT",
	}
	account_data := create.AccountDataRequest{
		Type:           "accounts",
		ID:             "9c36eefe-4105-4c60-886c-b95245293db0",
		OrganisationID: "a82eda54-62aa-446a-8dd1-3cbacf486062",
		Attributes:     &attributes,
		Version:        &version,
	}
	return &account_data
}
