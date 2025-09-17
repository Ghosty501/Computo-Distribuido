package model

type InteresadoID string

type InteresadoType string

const (
	InteresadoTypeVisit = InteresadoType("interesado")
)

type UserID string
type InteresadosValue int

type Interesados struct {
	InteresadoID   string           `json:"interesadoID"`
	InteresadoType string           `json:"interesadoType"`
	UserID         UserID           `json:"userId"`
	Value          InteresadosValue `json:"value"`
}
