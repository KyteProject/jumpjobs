package models

import "time"

// Job is a job posting structure
type Job struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Location     string    `json:"location"`
	Company      string    `json:"company"`
	CompanyURL   string    `json:"company_url"`
	Pay          string    `json:"pay"`
	Currency     string    `json:"currency"`
	Description  string    `json:"description"`
	JobLevel     string    `json:"job_level"`
	Email        string    `json:"email"`
	StripeToken  string    `json:"stripe_token,omitempty"`
	Type         string    `json:"type"`
	CurrencyCode string    `json:"currency_code"`
	CompanyIcon  string    `json:"company_icon,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
