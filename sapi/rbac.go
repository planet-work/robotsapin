package sapi

type UserContext struct {
	UserID string
	Right  string
}

func NewUserContext(t *Token) *UserContext {
	uc := &UserContext{UserID: t.UserID.String}
	/*
		if t.GrantType == "password" {
		} else if t.GrantType == "authorization_code" {
			a, err := authorizationQueryByCode(t.AuthorizationCode.String)
			if err == nil {
				for _, x := range a.Roles.Roles {
					switch x {
					case "dnsserver":
						uc.Roles[x] = new(Role)
						uc.Roles[x].tenants = append(uc.Roles[x].tenants, "*")
					default:
						logE.Println("unknown special role")
					}
				}
			}
		}
		uc.Filters = make(map[string]string)
	*/
	return uc
}
