// Code generated by "make api"; DO NOT EDIT.
package roles

import (
	"context"
	"fmt"

	"github.com/hashicorp/watchtower/api"
)

func (s Role) CreateRole(ctx context.Context, r *Role) (*Role, *api.Error, error) {
	if s.Client == nil {
		return nil, nil, fmt.Errorf("nil client in CreateRole request")
	}

	var opts []api.Option
	if s.Scope.Id != "" {
		// If it's explicitly set here, override anything that might be in the
		// client
		opts = append(opts, api.WithScopeId(s.Scope.Id))
	}

	req, err := s.Client.NewRequest(ctx, "POST", "roles", r, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating CreateRole request: %w", err)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during CreateRole call: %w", err)
	}

	target := new(Role)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding CreateRole repsonse: %w", err)
	}

	target.Client = s.Client

	return target, apiErr, nil
}