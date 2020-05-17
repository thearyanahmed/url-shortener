package main

import (
	"errors"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	RedirectNotFound = error.New("Redirect not found.")
	InvalidRedirect = error.New("Invalid redirect.")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepository RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errors.Wrap(InvalidRedirect,"service.Redirect.store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()

	return r.redirectRepo.Store(redirect)
}