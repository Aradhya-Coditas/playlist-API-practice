package models

type BFFSearchEquityGroupRequest struct {
	ExchangeName string `json:"exchange" validate:"required" example:"NSE"`
}
type BFFSearchEquityGroupResponse struct {
	Groups []string `json:"groups" example:"AA,AB,AC,AD,AG,AH,AI,AJ,AK,AL,AM,AN,EQ"`
}
