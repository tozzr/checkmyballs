package models

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
)

type Morph struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      string       `json:"name" db:"name" order_by:"name asc"`
	Filename  string       `json:"filename" db:"filename"`
	File      binding.File `db:"-" form:"someFile"`
	Dominant  bool         `json:"dominant" db:"dominant"`
	Rating    int          `json:"rating" db:"rating"`
}

// String is not required by pop and may be deleted
func (m Morph) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Morphs is not required by pop and may be deleted
type Morphs []Morph

// String is not required by pop and may be deleted
func (m Morphs) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Morph) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Morph) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Morph) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (m *Morph) AfterCreate(tx *pop.Connection) error {
	return moveFile(m.File, tx)
}

func (m *Morph) AfterUpdate(tx *pop.Connection) error {
	return moveFile(m.File, tx)
}

func moveFile(file binding.File, tx *pop.Connection) error {
	if file.Filename == "" {
		return nil
	}
	dir := filepath.Join(".", "/public/uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, file.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	return err
}
