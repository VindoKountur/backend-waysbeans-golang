package authdto

type LoginReponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
