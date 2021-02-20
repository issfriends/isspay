package value

type Membership int8

const (
	Admin Membership = iota + 1
	Manager
	NormalUser
)

type Role int8

func (enum Role) String() string {
	switch enum {
	case Master:
		return "master"
	case Faculty:
		return "phd"
	case Professor:
		return "faculty"
	case Alumni:
		return "professor"
	default:
		return ""
	}
}

const (
	Master Role = iota + 1
	Phd
	Faculty
	Professor
	Alumni
)

type OrderStatus int8

type ProductCategory int8

func (enum ProductCategory) String() string {
	switch enum {
	case Snake:
		return "snake"
	case Drink:
		return "drink"
	default:
		return ""
	}
}

const (
	Snake ProductCategory = iota + 1
	Drink
)
