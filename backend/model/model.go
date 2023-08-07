package model

import (
	"fmt"
	"net/url"
	"strings"
)

type PasswordCard struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

func (p *PasswordCard) Validate() error {
	if strings.TrimSpace(p.ID) == "" {
		return fmt.Errorf("invalid id")
	}

	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("invalid name")
	}

	if strings.TrimSpace(p.Username) == "" {
		return fmt.Errorf("username can't be empty")
	}

	if strings.TrimSpace(p.Password) == "" {
		return fmt.Errorf("password can't be empty")
	}

	if strings.TrimSpace(p.URL) == "" {
		return fmt.Errorf("invalid URL")
	}

	_, err := url.Parse(p.URL)
	if err != nil {
		return fmt.Errorf("invalid URL provided: %w", err)
	}

	return nil
}
