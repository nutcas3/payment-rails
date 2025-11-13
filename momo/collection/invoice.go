package collection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	invoicePath = "/collection/v2_0/invoice"
)

// CreateInvoice is used to create an invoice
//
// See [CreateInvoice] docs for more information.
//
// [CreateInvoice]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=CreateInvoice
func (c Collection) CreateInvoice(ctx context.Context, refID uuid.UUID, callbackURL string, body types.CreateInvoiceInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, map[string]string{
		callbackHeader: callbackURL,
		refHeader:      refID.String(),
	})
	if err != nil {
		return err
	}

	err = c.backend.Call(
		ctx,
		http.MethodPost,
		invoicePath,
		headers,
		nil,
		body,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// CancelInvoice is used to delete an invoice with given ID.
//
// See [CancelInvoice] docs for more information.
//
// [CancelInvoice]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=CancelInvoice
func (c Collection) CancelInvoice(ctx context.Context, invoiceID, transactionID uuid.UUID, callbackURL string) error {
	if transactionID == uuid.Nil || invoiceID == uuid.Nil {
		return fmt.Errorf("transactionID and invoiceID are required")
	}

	headers, err := c.getHeaders(ctx, map[string]string{
		callbackHeader: callbackURL,
		refHeader:      transactionID.String(),
	})
	if err != nil {
		return err
	}

	err = c.backend.Call(
		ctx,
		http.MethodDelete,
		invoicePath,
		headers,
		&common.Params{
			Path: []string{invoiceID.String()},
		},
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetInvoiceStatus is used to get the status of an invoice.
//
// See [GetInvoiceStatus] docs for more information.
//
// [GetInvoiceStatus]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetInvoiceStatus
func (c Collection) GetInvoiceStatus(ctx context.Context, refID uuid.UUID) (*types.InvoiceStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.InvoiceStatus

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		invoicePath,
		headers,
		&common.Params{
			Path: []string{refID.String()},
		},
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
