// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateManifestParams create manifest params
//
// swagger:model create-manifest-params
type CreateManifestParams struct {

	// base64 encoded manifest content.
	// Required: true
	Content *string `json:"content"`

	// The name of the manifest to customize the installed OCP cluster.
	// Required: true
	// Pattern: ^[^/]*\.(yaml|yml|json)$
	FileName *string `json:"file_name"`

	// The folder that contains the files. Manifests can be placed in 'manifests' or 'openshift' directories.
	// Enum: [manifests openshift]
	Folder *string `json:"folder,omitempty"`
}

// Validate validates this create manifest params
func (m *CreateManifestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateManifestParams) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *CreateManifestParams) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("file_name", "body", m.FileName); err != nil {
		return err
	}

	if err := validate.Pattern("file_name", "body", string(*m.FileName), `^[^/]*\.(yaml|yml|json)$`); err != nil {
		return err
	}

	return nil
}

var createManifestParamsTypeFolderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manifests","openshift"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createManifestParamsTypeFolderPropEnum = append(createManifestParamsTypeFolderPropEnum, v)
	}
}

const (

	// CreateManifestParamsFolderManifests captures enum value "manifests"
	CreateManifestParamsFolderManifests string = "manifests"

	// CreateManifestParamsFolderOpenshift captures enum value "openshift"
	CreateManifestParamsFolderOpenshift string = "openshift"
)

// prop value enum
func (m *CreateManifestParams) validateFolderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createManifestParamsTypeFolderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateManifestParams) validateFolder(formats strfmt.Registry) error {

	if swag.IsZero(m.Folder) { // not required
		return nil
	}

	// value enum
	if err := m.validateFolderEnum("folder", "body", *m.Folder); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateManifestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateManifestParams) UnmarshalBinary(b []byte) error {
	var res CreateManifestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
