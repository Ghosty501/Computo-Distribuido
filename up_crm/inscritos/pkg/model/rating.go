package model

type InscritoID string

type InscritoType string

const (
	InscritoTypeInst = InscritoType("inscrito")
)

type UserID string
type InscritosValue int

type Inscritos struct {
	InscritoID    string         `json:"inscritoID"`
	InscritosType string         `json:"inscritoType"`
	UserID        UserID         `json:"userId"`
	Value         InscritosValue `json:"value"`
}
