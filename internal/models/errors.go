package models

import "errors"

var (
	ErrGameNotFound   = errors.New("game not found")
	ErrInvalidCountry = errors.New("invalid country")
)
