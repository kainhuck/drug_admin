package util

import "drug_api/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecretManager = []byte(setting.AppSetting.JwtSecretManager)
	jwtSecretCustomer = []byte(setting.AppSetting.JwtSecretCustomer)
	jwtSecretEmployee = []byte(setting.AppSetting.JwtSecretEmployee)
}