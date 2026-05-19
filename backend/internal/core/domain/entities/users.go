package entities

import "time"

type Users struct {
	ID                 uint      `json:"id"`
	Name               string    `json:"cmuitaccount_name"`
	Email              string    `json:"cmuitaccount"`
	Password           string    `json:"-"`
	Role               string    `json:"role"`
	PreNameID          string    `json:"prename_id"`
	PreNameTH          string    `json:"prename_th"`
	PreNameEN          string    `json:"prename_en"`
	FirstNameTH        string    `json:"firstname_th"`
	FirstNameEN        string    `json:"firstname_en"`
	LastNameTH         string    `json:"lastname_th"`
	LastNameEN         string    `json:"lastname_en"`
	OrganizationCode   string    `json:"organization_code"`
	OrganizationNameTH string    `json:"organization_name_th"`
	OrganizationNameEN string    `json:"organization_name_en"`
	ITAccountTypeID    string    `json:"itaccounttype_id"`
	ITAccountTypeTH    string    `json:"itaccounttype_th"`
	ITAccountTypeEN    string    `json:"itaccounttype_en"`
	IsFirstTime        bool      `json:"is_first_time"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Role struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

type UsersLoginRequest struct {
	Email string `json:"email"`
}

type UsersLoginResponse struct {
	Message string `json:"message"`
}
