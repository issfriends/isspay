package value

const (
	Admin Membership = iota + 1
	Manager
	NormalUser
)

type Membership int8

type Role int8

func (enum *Role) FromString(str string) error {
	switch str {
	case "master", "碩士":
		*enum = Master
	case "faculty", "教職員":
		*enum = Faculty
	case "phd", "博士":
		*enum = Phd
	case "alumni", "校友":
		*enum = Alumni
	default:
	}
	return nil
}

func (enum Role) String() string {
	switch enum {
	case Master:
		return "master"
	case Faculty:
		return "faculty"
	case Professor:
		return "professor"
	case Phd:
		return "phd"
	case Alumni:
		return "alumni"
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
	Interviewee
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
