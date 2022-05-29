package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leg100/otf"
	"github.com/leg100/otf/sql/pggen"
)

var (
	_ otf.WorkspaceStore = (*WorkspaceDB)(nil)
)

type WorkspaceDB struct {
	*pgxpool.Pool
}

func NewWorkspaceDB(conn *pgxpool.Pool) *WorkspaceDB {
	return &WorkspaceDB{
		Pool: conn,
	}
}

func (db WorkspaceDB) Create(ws *otf.Workspace) error {
	q := pggen.NewQuerier(db.Pool)
	ctx := context.Background()
	_, err := q.InsertWorkspace(ctx, pggen.InsertWorkspaceParams{
		ID:                         pgtype.Text{String: ws.ID(), Status: pgtype.Present},
		CreatedAt:                  ws.CreatedAt(),
		UpdatedAt:                  ws.UpdatedAt(),
		Name:                       pgtype.Text{String: ws.Name(), Status: pgtype.Present},
		AllowDestroyPlan:           ws.AllowDestroyPlan(),
		CanQueueDestroyPlan:        ws.CanQueueDestroyPlan(),
		Environment:                pgtype.Text{String: ws.Environment(), Status: pgtype.Present},
		Description:                pgtype.Text{String: ws.Description(), Status: pgtype.Present},
		ExecutionMode:              pgtype.Text{String: string(ws.ExecutionMode()), Status: pgtype.Present},
		FileTriggersEnabled:        ws.FileTriggersEnabled(),
		GlobalRemoteState:          ws.GlobalRemoteState(),
		MigrationEnvironment:       pgtype.Text{String: ws.MigrationEnvironment(), Status: pgtype.Present},
		SourceName:                 pgtype.Text{String: ws.SourceName(), Status: pgtype.Present},
		SourceURL:                  pgtype.Text{String: ws.SourceURL(), Status: pgtype.Present},
		SpeculativeEnabled:         ws.SpeculativeEnabled(),
		StructuredRunOutputEnabled: ws.StructuredRunOutputEnabled(),
		TerraformVersion:           pgtype.Text{String: ws.TerraformVersion(), Status: pgtype.Present},
		TriggerPrefixes:            ws.TriggerPrefixes(),
		QueueAllRuns:               ws.QueueAllRuns(),
		AutoApply:                  ws.AutoApply(),
		WorkingDirectory:           pgtype.Text{String: ws.WorkingDirectory(), Status: pgtype.Present},
		OrganizationID:             pgtype.Text{String: ws.OrganizationID(), Status: pgtype.Present},
	})
	if err != nil {
		return databaseError(err)
	}
	return nil
}

func (db WorkspaceDB) Update(spec otf.WorkspaceSpec, fn func(*otf.Workspace) error) (*otf.Workspace, error) {
	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	q := pggen.NewQuerier(tx)
	// retrieve workspace
	ws, err := db.getWorkspaceForUpdate(ctx, tx, spec)
	if err != nil {
		return nil, databaseError(err)
	}
	// update workspace
	if err := fn(ws); err != nil {
		return nil, err
	}
	// persist update
	_, err = q.UpdateWorkspaceByID(ctx, pggen.UpdateWorkspaceByIDParams{
		ID:                         pgtype.Text{String: ws.ID(), Status: pgtype.Present},
		UpdatedAt:                  ws.UpdatedAt(),
		AllowDestroyPlan:           ws.AllowDestroyPlan(),
		Description:                pgtype.Text{String: ws.Description(), Status: pgtype.Present},
		ExecutionMode:              pgtype.Text{String: string(ws.ExecutionMode()), Status: pgtype.Present},
		Name:                       pgtype.Text{String: ws.Name(), Status: pgtype.Present},
		QueueAllRuns:               ws.QueueAllRuns(),
		SpeculativeEnabled:         ws.SpeculativeEnabled(),
		StructuredRunOutputEnabled: ws.StructuredRunOutputEnabled(),
		TerraformVersion:           pgtype.Text{String: ws.TerraformVersion(), Status: pgtype.Present},
		TriggerPrefixes:            ws.TriggerPrefixes(),
		WorkingDirectory:           pgtype.Text{String: ws.WorkingDirectory(), Status: pgtype.Present},
	})
	if err != nil {
		return nil, err
	}
	return ws, tx.Commit(ctx)
}

// Lock the specified workspace.
func (db WorkspaceDB) Lock(spec otf.WorkspaceSpec, opts otf.WorkspaceLockOptions) (*otf.Workspace, error) {
	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	q := pggen.NewQuerier(tx)
	// retrieve workspace
	ws, err := db.getWorkspaceForUpdate(ctx, tx, spec)
	if err != nil {
		return nil, databaseError(err)
	}
	// lock the workspace
	if err := ws.Lock(opts.Requestor); err != nil {
		return nil, err
	}
	// persist to db
	params, err := otf.MarshalWorkspaceLockParams(ws)
	if err != nil {
		return nil, err
	}
	_, err = q.UpdateWorkspaceLockByID(ctx, params)
	if err != nil {
		return nil, databaseError(err)
	}
	// return ws with new lock
	return ws, tx.Commit(ctx)
}

// Unlock the specified workspace; the caller has the opportunity to check the
// current locker passed into the provided callback. If an error is returned the
// unlock will not go ahead.
func (db WorkspaceDB) Unlock(spec otf.WorkspaceSpec, opts otf.WorkspaceUnlockOptions) (*otf.Workspace, error) {
	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	q := pggen.NewQuerier(tx)
	// retrieve workspace
	ws, err := db.getWorkspaceForUpdate(ctx, tx, spec)
	if err != nil {
		return nil, databaseError(err)
	}
	// unlock workspace
	if err := ws.Unlock(opts.Requestor, opts.Force); err != nil {
		return nil, err
	}
	// persist to db
	params, err := otf.MarshalWorkspaceLockParams(ws)
	if err != nil {
		return nil, err
	}
	_, err = q.UpdateWorkspaceLockByID(ctx, params)
	if err != nil {
		return nil, databaseError(err)
	}
	// return ws with new lock
	return ws, tx.Commit(ctx)
}

func (db WorkspaceDB) List(opts otf.WorkspaceListOptions) (*otf.WorkspaceList, error) {
	q := pggen.NewQuerier(db.Pool)
	batch := &pgx.Batch{}
	ctx := context.Background()

	q.FindWorkspacesBatch(batch, pggen.FindWorkspacesParams{
		OrganizationName:    pgtype.Text{String: opts.OrganizationName, Status: pgtype.Present},
		Prefix:              pgtype.Text{String: opts.Prefix, Status: pgtype.Present},
		Limit:               opts.GetLimit(),
		Offset:              opts.GetOffset(),
		IncludeOrganization: includeOrganization(opts.Include),
	})
	q.CountWorkspacesBatch(batch,
		pgtype.Text{String: opts.Prefix, Status: pgtype.Present},
		pgtype.Text{String: opts.OrganizationName, Status: pgtype.Present},
	)

	results := db.Pool.SendBatch(ctx, batch)
	defer results.Close()

	rows, err := q.FindWorkspacesScan(results)
	if err != nil {
		return nil, err
	}
	count, err := q.CountWorkspacesScan(results)
	if err != nil {
		return nil, err
	}

	var items []*otf.Workspace
	for _, r := range rows {
		ws, err := otf.UnmarshalWorkspaceDBResult(otf.WorkspaceDBResult(r))
		if err != nil {
			return nil, err
		}
		items = append(items, ws)
	}

	return &otf.WorkspaceList{
		Items:      items,
		Pagination: otf.NewPagination(opts.ListOptions, *count),
	}, nil
}

func (db WorkspaceDB) Get(spec otf.WorkspaceSpec) (*otf.Workspace, error) {
	ctx := context.Background()
	q := pggen.NewQuerier(db.Pool)

	if spec.ID != nil {
		result, err := q.FindWorkspaceByID(ctx,
			includeOrganization(spec.Include),
			pgtype.Text{String: *spec.ID, Status: pgtype.Present},
		)
		if err != nil {
			return nil, databaseError(err)
		}
		return otf.UnmarshalWorkspaceDBResult(otf.WorkspaceDBResult(result))
	} else if spec.Name != nil && spec.OrganizationName != nil {
		result, err := q.FindWorkspaceByName(ctx, pggen.FindWorkspaceByNameParams{
			Name:                pgtype.Text{String: *spec.Name, Status: pgtype.Present},
			OrganizationName:    pgtype.Text{String: *spec.OrganizationName, Status: pgtype.Present},
			IncludeOrganization: includeOrganization(spec.Include),
		})
		if err != nil {
			return nil, databaseError(err)
		}
		return otf.UnmarshalWorkspaceDBResult(otf.WorkspaceDBResult(result))
	} else {
		return nil, fmt.Errorf("no workspace spec provided")
	}
}

// Delete deletes a specific workspace, along with its child records (runs etc).
func (db WorkspaceDB) Delete(spec otf.WorkspaceSpec) error {
	ctx := context.Background()
	q := pggen.NewQuerier(db.Pool)

	var result pgconn.CommandTag
	var err error

	if spec.ID != nil {
		result, err = q.DeleteWorkspaceByID(ctx, pgtype.Text{String: *spec.ID, Status: pgtype.Present})
	} else if spec.Name != nil && spec.OrganizationName != nil {
		result, err = q.DeleteWorkspaceByName(ctx,
			pgtype.Text{String: *spec.Name, Status: pgtype.Present},
			pgtype.Text{String: *spec.OrganizationName, Status: pgtype.Present},
		)
	} else {
		return fmt.Errorf("no workspace spec provided")
	}
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return otf.ErrResourceNotFound
	}

	return nil
}

// Lock the specified workspace.
func (db WorkspaceDB) getWorkspaceForUpdate(ctx context.Context, tx pgx.Tx, spec otf.WorkspaceSpec) (*otf.Workspace, error) {
	q := pggen.NewQuerier(tx)
	workspaceID, err := getWorkspaceID(ctx, q, spec)
	if err != nil {
		return nil, err
	}
	result, err := q.FindWorkspaceByIDForUpdate(ctx, workspaceID)
	if err != nil {
		return nil, err
	}
	return otf.UnmarshalWorkspaceDBResult(otf.WorkspaceDBResult(result))
}

func getWorkspaceID(ctx context.Context, q *pggen.DBQuerier, spec otf.WorkspaceSpec) (pgtype.Text, error) {
	if spec.ID != nil {
		return pgtype.Text{String: *spec.ID, Status: pgtype.Present}, nil
	}
	if spec.Name != nil && spec.OrganizationName != nil {
		return q.FindWorkspaceIDByName(ctx,
			pgtype.Text{String: *spec.Name, Status: pgtype.Present},
			pgtype.Text{String: *spec.OrganizationName, Status: pgtype.Present},
		)
	}
	return pgtype.Text{}, otf.ErrInvalidWorkspaceSpec
}
