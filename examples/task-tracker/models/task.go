// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Task a structure describing a complete task.
//
// A Task is the main entity in this application. Everything revolves around tasks and managing them.
//
//
// swagger:model Task
type Task struct {
	TaskCard

	// The attached files.
	//
	// An issue can have at most 20 files attached to it.
	//
	Attachments map[string]TaskAttachmentsAnon `json:"attachments,omitempty"`

	// The 5 most recent items for this issue.
	//
	// The detail view of an issue includes the 5 most recent comments.
	// This field is read only, comments are added through a separate process.
	//
	// Read Only: true
	Comments []*Comment `json:"comments"`

	// The time at which this issue was last updated.
	//
	// This field is read only so it's only sent as part of the response.
	//
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// last updated by
	LastUpdatedBy *UserCard `json:"lastUpdatedBy,omitempty"`

	// reported by
	ReportedBy *UserCard `json:"reportedBy,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Task) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TaskCard
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TaskCard = aO0

	// AO1
	var dataAO1 struct {
		Attachments map[string]TaskAttachmentsAnon `json:"attachments,omitempty"`

		Comments []*Comment `json:"comments"`

		LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

		LastUpdatedBy *UserCard `json:"lastUpdatedBy,omitempty"`

		ReportedBy *UserCard `json:"reportedBy,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Attachments = dataAO1.Attachments

	m.Comments = dataAO1.Comments

	m.LastUpdated = dataAO1.LastUpdated

	m.LastUpdatedBy = dataAO1.LastUpdatedBy

	m.ReportedBy = dataAO1.ReportedBy

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Task) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TaskCard)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Attachments map[string]TaskAttachmentsAnon `json:"attachments,omitempty"`

		Comments []*Comment `json:"comments"`

		LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

		LastUpdatedBy *UserCard `json:"lastUpdatedBy,omitempty"`

		ReportedBy *UserCard `json:"reportedBy,omitempty"`
	}

	dataAO1.Attachments = m.Attachments

	dataAO1.Comments = m.Comments

	dataAO1.LastUpdated = m.LastUpdated

	dataAO1.LastUpdatedBy = m.LastUpdatedBy

	dataAO1.ReportedBy = m.ReportedBy

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TaskCard
	if err := m.TaskCard.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateAttachments(formats strfmt.Registry) error {

	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for k := range m.Attachments {

		if swag.IsZero(m.Attachments[k]) { // not required
			continue
		}
		if val, ok := m.Attachments[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateLastUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedBy) { // not required
		return nil
	}

	if m.LastUpdatedBy != nil {
		if err := m.LastUpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Task) validateReportedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportedBy) { // not required
		return nil
	}

	if m.ReportedBy != nil {
		if err := m.ReportedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reportedBy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TaskCard
	if err := m.TaskCard.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Attachments {

		if val, ok := m.Attachments[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Task) contextValidateComments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "comments", "body", []*Comment(m.Comments)); err != nil {
		return err
	}

	for i := 0; i < len(m.Comments); i++ {

		if m.Comments[i] != nil {
			if err := m.Comments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *Task) contextValidateLastUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.LastUpdatedBy != nil {
		if err := m.LastUpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUpdatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Task) contextValidateReportedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ReportedBy != nil {
		if err := m.ReportedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reportedBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskAttachmentsAnon task attachments anon
//
// swagger:model TaskAttachmentsAnon
type TaskAttachmentsAnon struct {

	// The content type of the file.
	//
	// The content type of the file is inferred from the upload request.
	//
	// Read Only: true
	ContentType string `json:"contentType,omitempty"`

	// Extra information to attach to the file.
	//
	// This is a free form text field with support for github flavored markdown.
	//
	// Min Length: 3
	Description string `json:"description,omitempty"`

	// The name of the file.
	//
	// This name is inferred from the upload request.
	//
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The file size in bytes.
	//
	// This property was generated during the upload request of the file.
	// Read Only: true
	Size float64 `json:"size,omitempty"`

	// The url to download or view the file.
	//
	// This URL is generated on the server, based on where it was able to store the file when it was uploaded.
	//
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this task attachments anon
func (m *TaskAttachmentsAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAttachmentsAnon) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", m.Description, 3); err != nil {
		return err
	}

	return nil
}

func (m *TaskAttachmentsAnon) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this task attachments anon based on the context it is used
func (m *TaskAttachmentsAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAttachmentsAnon) contextValidateContentType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "contentType", "body", string(m.ContentType)); err != nil {
		return err
	}

	return nil
}

func (m *TaskAttachmentsAnon) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *TaskAttachmentsAnon) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", float64(m.Size)); err != nil {
		return err
	}

	return nil
}

func (m *TaskAttachmentsAnon) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAttachmentsAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAttachmentsAnon) UnmarshalBinary(b []byte) error {
	var res TaskAttachmentsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
