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

const (
	Completed OrderStatus = iota + 1
	Canceled
)

func (enum OrderStatus) String() string {
	switch enum {
	case Completed:
		return "completed"
	case Canceled:
		return "canceled"
	default:
		return ""
	}
}

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

type AccountStatus int8

const (
	NormalStatus AccountStatus = iota + 1
	ForbiddenStatus
)
