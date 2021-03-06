package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JAccountFetchEmailReader is a Reader for the JAccountFetchEmail structure.
type JAccountFetchEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountFetchEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountFetchEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountFetchEmailOK creates a JAccountFetchEmailOK with default headers values
func NewJAccountFetchEmailOK() *JAccountFetchEmailOK {
	return &JAccountFetchEmailOK{}
}

/*JAccountFetchEmailOK handles this case with default header values.

OK
*/
type JAccountFetchEmailOK struct {
	Payload JAccountFetchEmailOKBody
}

func (o *JAccountFetchEmailOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.fetchEmail/{id}][%d] jAccountFetchEmailOK  %+v", 200, o.Payload)
}

func (o *JAccountFetchEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JAccountFetchEmailOKBody j account fetch email o k body
swagger:model JAccountFetchEmailOKBody
*/
type JAccountFetchEmailOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JAccountFetchEmailOKBody) UnmarshalJSON(raw []byte) error {

	var jAccountFetchEmailOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &jAccountFetchEmailOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = jAccountFetchEmailOKBodyAO0

	var jAccountFetchEmailOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jAccountFetchEmailOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jAccountFetchEmailOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JAccountFetchEmailOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jAccountFetchEmailOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountFetchEmailOKBodyAO0)

	jAccountFetchEmailOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountFetchEmailOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j account fetch email o k body
func (o *JAccountFetchEmailOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
