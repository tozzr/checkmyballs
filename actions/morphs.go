package actions

import (
	"github.com/tozzr/checkmyballs/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Morph)
// DB Table: Plural (morphs)
// Resource: Plural (Morphs)
// Path: Plural (/morphs)
// View Template Folder: Plural (/templates/morphs/)

// MorphsResource is the resource for the Morph model
type MorphsResource struct {
	buffalo.Resource
}

// List gets all Morphs. This function is mapped to the path
// GET /morphs
func (v MorphsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	morphs := &models.Morphs{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	order_by := c.Params().Get("order_by")
	if order_by == "name" {
		order_by = "name asc"
	} else {
		order_by = "rating desc"
	}

	// Retrieve all Morphs from the DB
	if err := q.Order(order_by).All(morphs); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, morphs))
}

// Show gets the data for one Morph. This function is mapped to
// the path GET /morphs/{morph_id}
func (v MorphsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Morph
	morph := &models.Morph{}

	// To find the Morph the parameter morph_id is used.
	if err := tx.Find(morph, c.Param("morph_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, morph))
}

// New renders the form for creating a new Morph.
// This function is mapped to the path GET /morphs/new
func (v MorphsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Morph{}))
}

// Create adds a Morph to the DB. This function is mapped to the
// path POST /morphs
func (v MorphsResource) Create(c buffalo.Context) error {
	// Allocate an empty Morph
	morph := &models.Morph{}

	// Bind morph to the html form elements
	if err := c.Bind(morph); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	f, err := c.File("someFile")
	if err != nil {
		return errors.WithStack(err)
	} else if f.Filename != "" {
		morph.Filename = f.Filename
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(morph)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, morph))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Morph was created successfully")

	// and redirect to the morphs index page
	return c.Render(201, r.Auto(c, morph))
}

// Edit renders a edit form for a Morph. This function is
// mapped to the path GET /morphs/{morph_id}/edit
func (v MorphsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Morph
	morph := &models.Morph{}

	if err := tx.Find(morph, c.Param("morph_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, morph))
}

// Update changes a Morph in the DB. This function is mapped to
// the path PUT /morphs/{morph_id}
func (v MorphsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Morph
	morph := &models.Morph{}

	if err := tx.Find(morph, c.Param("morph_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Morph to the html form elements
	if err := c.Bind(morph); err != nil {
		return errors.WithStack(err)
	}

	f, err := c.File("someFile")
	if err != nil {
		return errors.WithStack(err)
	} else if f.Filename != "" {
		morph.Filename = f.Filename
	}

	verrs, err := tx.ValidateAndUpdate(morph)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, morph))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Morph was updated successfully")

	// and redirect to the morphs index page
	return c.Render(200, r.Auto(c, morph))
}

// Destroy deletes a Morph from the DB. This function is mapped
// to the path DELETE /morphs/{morph_id}
func (v MorphsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Morph
	morph := &models.Morph{}

	// To find the Morph the parameter morph_id is used.
	if err := tx.Find(morph, c.Param("morph_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(morph); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Morph was destroyed successfully")

	// Redirect to the morphs index page
	return c.Render(200, r.Auto(c, morph))
}
