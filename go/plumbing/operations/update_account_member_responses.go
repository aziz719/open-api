// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// UpdateAccountMemberReader is a Reader for the UpdateAccountMember structure.
type UpdateAccountMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccountMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAccountMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAccountMemberOK creates a UpdateAccountMemberOK with default headers values
func NewUpdateAccountMemberOK() *UpdateAccountMemberOK {
	return &UpdateAccountMemberOK{}
}

/*UpdateAccountMemberOK handles this case with default header values.

OK
*/
type UpdateAccountMemberOK struct {
	Payload *models.Member
}

func (o *UpdateAccountMemberOK) Error() string {
	return fmt.Sprintf("[PUT /{account_slug}/members/{member_id}][%d] updateAccountMemberOK  %+v", 200, o.Payload)
}

func (o *UpdateAccountMemberOK) GetPayload() *models.Member {
	return o.Payload
}

func (o *UpdateAccountMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Member)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccountMemberDefault creates a UpdateAccountMemberDefault with default headers values
func NewUpdateAccountMemberDefault(code int) *UpdateAccountMemberDefault {
	return &UpdateAccountMemberDefault{
		_statusCode: code,
	}
}

/*UpdateAccountMemberDefault handles this case with default header values.

error
*/
type UpdateAccountMemberDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update account member default response
func (o *UpdateAccountMemberDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAccountMemberDefault) Error() string {
	return fmt.Sprintf("[PUT /{account_slug}/members/{member_id}][%d] updateAccountMember default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAccountMemberDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAccountMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
