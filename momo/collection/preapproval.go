package collection

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

var ErrAccHolderIDAndTypeRequired = errors.New("accountHolderId and accountHolderIdType are required")

const (
	preApprovalPathv2 = "/collection/v2_0/preapproval"
	preApprovalPathv1 = "/collection/v1_0/preapproval"
)

// CancelPreApproval is used to cancel an approved pre-approval.
//
// See [CancelPreApproval] docs for more information.
//
// [CancelPreApproval]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=CancelPreApproval
func (c Collection) CancelPreApproval(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("refID is required")
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return err
	}

	err = c.backend.Call(
		ctx,
		http.MethodDelete,
		preApprovalPathv1,
		headers,
		&common.Params{
			Path: []string{id.String()},
		},
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetPreApprovalStatus is used to get the status of a pre-approval.
//
// See [GetPreApprovalStatus] docs for more information.
//
// [GetPreApprovalStatus]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetPreApprovalStatus
func (c Collection) GetPreApprovalStatus(ctx context.Context, refID uuid.UUID) (*types.PreApprovalStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.PreApprovalStatus

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		preApprovalPathv2,
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

// PreApproval is used to create a pre-approval.
//
// See [PreApproval] docs for more information.
//
// [PreApproval]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=PreApproval
func (c Collection) PreApproval(
	ctx context.Context,
	refID uuid.UUID,
	callbackURL string,
	body types.PreApprovalInput,
) error {
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
		preApprovalPathv2,
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

// GetApprovedPreApprovals is used to get approved pre-approvals of an account holder.
//
// See [GetApprovedPreApprovals] docs for more information.
//
// [GetApprovedPreApprovals]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetApprovedPreApprovals
func (c Collection) GetApprovedPreApprovals(
	ctx context.Context,
	accHolderIDType,
	accHolderID string,
) ([]*types.PreApprovalDetails, error) {
	if accHolderID == "" || accHolderIDType == "" {
		return nil, ErrAccHolderIDAndTypeRequired
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp []*types.PreApprovalDetails

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		preApprovalPathv1,
		headers,
		&common.Params{
			Path: []string{accHolderIDType, accHolderID},
		},
		nil,
		resp,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
