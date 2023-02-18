package companies

type companyType int

const (
	AGGREGATE companyType = iota + 1
	CONTRACT
	INTERMEDIATED
)

func (c companyType) String() string {
	switch c {
	case AGGREGATE:
		return "AGGREGATE"
	case CONTRACT:
		return "CONTRACT"
	case INTERMEDIATED:
		return "INTERMEDIATED"
	}
	panic(InvalidCompanyTypeErr)
}

func NewCompanyType(c string) companyType {
	switch c {
	case "AGGREGATE":
		return AGGREGATE
	case "CONTRACT":
		return CONTRACT
	case "INTERMEDIATED":
		return INTERMEDIATED
	}
	panic(InvalidCompanyTypeErr)
}

func getCompanyTypes() []string {
	return []string{"AGGREGATE", "CONTRACT", "INTERMEDIATED"}
}

func isCompanyType(ct string) bool {
	for _, _ct := range getCompanyTypes() {
		if ct == _ct {
			return true
		}
	}
	return false
}
