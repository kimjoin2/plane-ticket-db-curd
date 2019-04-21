package dto

type TicketPriceDTO struct {
	BaseNation string `json:"baseNation"`
	TargetNation string `json:"targetNation"`
	BaseAirport string `json:"baseAirPort"`
	TargetAirport string `json:"targetAirport"`
	BaseDepartTime string `json:"baseDepartTime"`
	TargetArrivalTime string `json:"targetArrivalTime"`
	TargetDepartTime string `json:"targetDepartTime"`
	BaseArrivalTime string `json:"baseArrivalTime"`
	BaseAirline string `json:"baseAirline"`
	TargetAirline string `json:"targetAirline"`
	Price int32 `json:"price"`
}