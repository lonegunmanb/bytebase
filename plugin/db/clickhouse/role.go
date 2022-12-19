package clickhouse

import (
	"context"

	"github.com/pkg/errors"

	v1pb "github.com/bytebase/bytebase/proto/generated-go/v1"
)

// CreateRole creates the role.
func (*Driver) CreateRole(_ context.Context, _ *v1pb.DatabaseRoleUpsert) (*v1pb.DatabaseRole, error) {
	return nil, errors.Errorf("create role for ClickHouse is not implemented yet")
}

// UpdateRole updates the role.
func (*Driver) UpdateRole(_ context.Context, _ string, _ *v1pb.DatabaseRoleUpsert) (*v1pb.DatabaseRole, error) {
	return nil, errors.Errorf("update role for ClickHouse is not implemented yet")
}

// FindRole finds the role by name.
func (*Driver) FindRole(_ context.Context, _ string) (*v1pb.DatabaseRole, error) {
	return nil, errors.Errorf("find role for ClickHouse is not implemented yet")
}

// DeleteRole deletes the role by name.
func (*Driver) DeleteRole(_ context.Context, _ string) error {
	return errors.Errorf("delete role for ClickHouse is not implemented yet")
}
