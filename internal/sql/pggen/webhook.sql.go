// Code generated by pggen. DO NOT EDIT.

package pggen

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const insertWebhookSQL = `WITH inserted AS (
    INSERT INTO webhooks (
        webhook_id,
        vcs_id,
        vcs_provider_id,
        secret,
        identifier
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    w.webhook_id,
    w.vcs_id,
    w.vcs_provider_id,
    w.secret,
    w.identifier,
    v.cloud
FROM inserted w
JOIN vcs_providers v USING (vcs_provider_id);`

type InsertWebhookParams struct {
	WebhookID     pgtype.UUID
	VCSID         pgtype.Text
	VCSProviderID pgtype.Text
	Secret        pgtype.Text
	Identifier    pgtype.Text
}

type InsertWebhookRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	Cloud         pgtype.Text `json:"cloud"`
}

// InsertWebhook implements Querier.InsertWebhook.
func (q *DBQuerier) InsertWebhook(ctx context.Context, params InsertWebhookParams) (InsertWebhookRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "InsertWebhook")
	row := q.conn.QueryRow(ctx, insertWebhookSQL, params.WebhookID, params.VCSID, params.VCSProviderID, params.Secret, params.Identifier)
	var item InsertWebhookRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
		return item, fmt.Errorf("query InsertWebhook: %w", err)
	}
	return item, nil
}

// InsertWebhookBatch implements Querier.InsertWebhookBatch.
func (q *DBQuerier) InsertWebhookBatch(batch genericBatch, params InsertWebhookParams) {
	batch.Queue(insertWebhookSQL, params.WebhookID, params.VCSID, params.VCSProviderID, params.Secret, params.Identifier)
}

// InsertWebhookScan implements Querier.InsertWebhookScan.
func (q *DBQuerier) InsertWebhookScan(results pgx.BatchResults) (InsertWebhookRow, error) {
	row := results.QueryRow()
	var item InsertWebhookRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
		return item, fmt.Errorf("scan InsertWebhookBatch row: %w", err)
	}
	return item, nil
}

const updateWebhookVCSIDSQL = `UPDATE webhooks
SET vcs_id = $1
WHERE webhook_id = $2
RETURNING *;`

type UpdateWebhookVCSIDRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
}

// UpdateWebhookVCSID implements Querier.UpdateWebhookVCSID.
func (q *DBQuerier) UpdateWebhookVCSID(ctx context.Context, vcsID pgtype.Text, webhookID pgtype.UUID) (UpdateWebhookVCSIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateWebhookVCSID")
	row := q.conn.QueryRow(ctx, updateWebhookVCSIDSQL, vcsID, webhookID)
	var item UpdateWebhookVCSIDRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.Secret, &item.Identifier, &item.VCSProviderID); err != nil {
		return item, fmt.Errorf("query UpdateWebhookVCSID: %w", err)
	}
	return item, nil
}

// UpdateWebhookVCSIDBatch implements Querier.UpdateWebhookVCSIDBatch.
func (q *DBQuerier) UpdateWebhookVCSIDBatch(batch genericBatch, vcsID pgtype.Text, webhookID pgtype.UUID) {
	batch.Queue(updateWebhookVCSIDSQL, vcsID, webhookID)
}

// UpdateWebhookVCSIDScan implements Querier.UpdateWebhookVCSIDScan.
func (q *DBQuerier) UpdateWebhookVCSIDScan(results pgx.BatchResults) (UpdateWebhookVCSIDRow, error) {
	row := results.QueryRow()
	var item UpdateWebhookVCSIDRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.Secret, &item.Identifier, &item.VCSProviderID); err != nil {
		return item, fmt.Errorf("scan UpdateWebhookVCSIDBatch row: %w", err)
	}
	return item, nil
}

const findWebhooksSQL = `SELECT
    w.webhook_id,
    w.vcs_id,
    w.vcs_provider_id,
    w.secret,
    w.identifier,
    v.cloud
FROM webhooks w
JOIN vcs_providers v USING (vcs_provider_id);`

type FindWebhooksRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	Cloud         pgtype.Text `json:"cloud"`
}

// FindWebhooks implements Querier.FindWebhooks.
func (q *DBQuerier) FindWebhooks(ctx context.Context) ([]FindWebhooksRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindWebhooks")
	rows, err := q.conn.Query(ctx, findWebhooksSQL)
	if err != nil {
		return nil, fmt.Errorf("query FindWebhooks: %w", err)
	}
	defer rows.Close()
	items := []FindWebhooksRow{}
	for rows.Next() {
		var item FindWebhooksRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindWebhooks row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindWebhooks rows: %w", err)
	}
	return items, err
}

// FindWebhooksBatch implements Querier.FindWebhooksBatch.
func (q *DBQuerier) FindWebhooksBatch(batch genericBatch) {
	batch.Queue(findWebhooksSQL)
}

// FindWebhooksScan implements Querier.FindWebhooksScan.
func (q *DBQuerier) FindWebhooksScan(results pgx.BatchResults) ([]FindWebhooksRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindWebhooksBatch: %w", err)
	}
	defer rows.Close()
	items := []FindWebhooksRow{}
	for rows.Next() {
		var item FindWebhooksRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindWebhooksBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindWebhooksBatch rows: %w", err)
	}
	return items, err
}

const findWebhookByIDSQL = `SELECT
    w.webhook_id,
    w.vcs_id,
    w.vcs_provider_id,
    w.secret,
    w.identifier,
    v.cloud
FROM webhooks w
JOIN vcs_providers v USING (vcs_provider_id)
WHERE w.webhook_id = $1;`

type FindWebhookByIDRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	Cloud         pgtype.Text `json:"cloud"`
}

// FindWebhookByID implements Querier.FindWebhookByID.
func (q *DBQuerier) FindWebhookByID(ctx context.Context, webhookID pgtype.UUID) (FindWebhookByIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindWebhookByID")
	row := q.conn.QueryRow(ctx, findWebhookByIDSQL, webhookID)
	var item FindWebhookByIDRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
		return item, fmt.Errorf("query FindWebhookByID: %w", err)
	}
	return item, nil
}

// FindWebhookByIDBatch implements Querier.FindWebhookByIDBatch.
func (q *DBQuerier) FindWebhookByIDBatch(batch genericBatch, webhookID pgtype.UUID) {
	batch.Queue(findWebhookByIDSQL, webhookID)
}

// FindWebhookByIDScan implements Querier.FindWebhookByIDScan.
func (q *DBQuerier) FindWebhookByIDScan(results pgx.BatchResults) (FindWebhookByIDRow, error) {
	row := results.QueryRow()
	var item FindWebhookByIDRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
		return item, fmt.Errorf("scan FindWebhookByIDBatch row: %w", err)
	}
	return item, nil
}

const findWebhookByRepoAndProviderSQL = `SELECT
    w.webhook_id,
    w.vcs_id,
    w.vcs_provider_id,
    w.secret,
    w.identifier,
    v.cloud
FROM webhooks w
JOIN vcs_providers v USING (vcs_provider_id)
WHERE identifier = $1
AND   vcs_provider_id = $2;`

type FindWebhookByRepoAndProviderRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	Cloud         pgtype.Text `json:"cloud"`
}

// FindWebhookByRepoAndProvider implements Querier.FindWebhookByRepoAndProvider.
func (q *DBQuerier) FindWebhookByRepoAndProvider(ctx context.Context, identifier pgtype.Text, vcsProviderID pgtype.Text) ([]FindWebhookByRepoAndProviderRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindWebhookByRepoAndProvider")
	rows, err := q.conn.Query(ctx, findWebhookByRepoAndProviderSQL, identifier, vcsProviderID)
	if err != nil {
		return nil, fmt.Errorf("query FindWebhookByRepoAndProvider: %w", err)
	}
	defer rows.Close()
	items := []FindWebhookByRepoAndProviderRow{}
	for rows.Next() {
		var item FindWebhookByRepoAndProviderRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindWebhookByRepoAndProvider row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindWebhookByRepoAndProvider rows: %w", err)
	}
	return items, err
}

// FindWebhookByRepoAndProviderBatch implements Querier.FindWebhookByRepoAndProviderBatch.
func (q *DBQuerier) FindWebhookByRepoAndProviderBatch(batch genericBatch, identifier pgtype.Text, vcsProviderID pgtype.Text) {
	batch.Queue(findWebhookByRepoAndProviderSQL, identifier, vcsProviderID)
}

// FindWebhookByRepoAndProviderScan implements Querier.FindWebhookByRepoAndProviderScan.
func (q *DBQuerier) FindWebhookByRepoAndProviderScan(results pgx.BatchResults) ([]FindWebhookByRepoAndProviderRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindWebhookByRepoAndProviderBatch: %w", err)
	}
	defer rows.Close()
	items := []FindWebhookByRepoAndProviderRow{}
	for rows.Next() {
		var item FindWebhookByRepoAndProviderRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindWebhookByRepoAndProviderBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindWebhookByRepoAndProviderBatch rows: %w", err)
	}
	return items, err
}

const findWebhooksByProviderSQL = `SELECT
    w.webhook_id,
    w.vcs_id,
    w.vcs_provider_id,
    w.secret,
    w.identifier,
    v.cloud
FROM webhooks w
JOIN vcs_providers v USING (vcs_provider_id)
WHERE w.vcs_provider_id = $1;`

type FindWebhooksByProviderRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	Cloud         pgtype.Text `json:"cloud"`
}

// FindWebhooksByProvider implements Querier.FindWebhooksByProvider.
func (q *DBQuerier) FindWebhooksByProvider(ctx context.Context, vcsProviderID pgtype.Text) ([]FindWebhooksByProviderRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindWebhooksByProvider")
	rows, err := q.conn.Query(ctx, findWebhooksByProviderSQL, vcsProviderID)
	if err != nil {
		return nil, fmt.Errorf("query FindWebhooksByProvider: %w", err)
	}
	defer rows.Close()
	items := []FindWebhooksByProviderRow{}
	for rows.Next() {
		var item FindWebhooksByProviderRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindWebhooksByProvider row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindWebhooksByProvider rows: %w", err)
	}
	return items, err
}

// FindWebhooksByProviderBatch implements Querier.FindWebhooksByProviderBatch.
func (q *DBQuerier) FindWebhooksByProviderBatch(batch genericBatch, vcsProviderID pgtype.Text) {
	batch.Queue(findWebhooksByProviderSQL, vcsProviderID)
}

// FindWebhooksByProviderScan implements Querier.FindWebhooksByProviderScan.
func (q *DBQuerier) FindWebhooksByProviderScan(results pgx.BatchResults) ([]FindWebhooksByProviderRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindWebhooksByProviderBatch: %w", err)
	}
	defer rows.Close()
	items := []FindWebhooksByProviderRow{}
	for rows.Next() {
		var item FindWebhooksByProviderRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindWebhooksByProviderBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindWebhooksByProviderBatch rows: %w", err)
	}
	return items, err
}

const findUnreferencedWebhooksSQL = `SELECT
    w.webhook_id,
    w.vcs_id,
    w.vcs_provider_id,
    w.secret,
    w.identifier,
    v.cloud
FROM webhooks w
JOIN vcs_providers v USING (vcs_provider_id)
WHERE NOT EXISTS (
    SELECT FROM repo_connections c
    WHERE c.webhook_id = w.webhook_id
);`

type FindUnreferencedWebhooksRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	Cloud         pgtype.Text `json:"cloud"`
}

// FindUnreferencedWebhooks implements Querier.FindUnreferencedWebhooks.
func (q *DBQuerier) FindUnreferencedWebhooks(ctx context.Context) ([]FindUnreferencedWebhooksRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindUnreferencedWebhooks")
	rows, err := q.conn.Query(ctx, findUnreferencedWebhooksSQL)
	if err != nil {
		return nil, fmt.Errorf("query FindUnreferencedWebhooks: %w", err)
	}
	defer rows.Close()
	items := []FindUnreferencedWebhooksRow{}
	for rows.Next() {
		var item FindUnreferencedWebhooksRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindUnreferencedWebhooks row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindUnreferencedWebhooks rows: %w", err)
	}
	return items, err
}

// FindUnreferencedWebhooksBatch implements Querier.FindUnreferencedWebhooksBatch.
func (q *DBQuerier) FindUnreferencedWebhooksBatch(batch genericBatch) {
	batch.Queue(findUnreferencedWebhooksSQL)
}

// FindUnreferencedWebhooksScan implements Querier.FindUnreferencedWebhooksScan.
func (q *DBQuerier) FindUnreferencedWebhooksScan(results pgx.BatchResults) ([]FindUnreferencedWebhooksRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindUnreferencedWebhooksBatch: %w", err)
	}
	defer rows.Close()
	items := []FindUnreferencedWebhooksRow{}
	for rows.Next() {
		var item FindUnreferencedWebhooksRow
		if err := rows.Scan(&item.WebhookID, &item.VCSID, &item.VCSProviderID, &item.Secret, &item.Identifier, &item.Cloud); err != nil {
			return nil, fmt.Errorf("scan FindUnreferencedWebhooksBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindUnreferencedWebhooksBatch rows: %w", err)
	}
	return items, err
}

const deleteWebhookByIDSQL = `DELETE
FROM webhooks
WHERE webhook_id = $1
RETURNING *;`

type DeleteWebhookByIDRow struct {
	WebhookID     pgtype.UUID `json:"webhook_id"`
	VCSID         pgtype.Text `json:"vcs_id"`
	Secret        pgtype.Text `json:"secret"`
	Identifier    pgtype.Text `json:"identifier"`
	VCSProviderID pgtype.Text `json:"vcs_provider_id"`
}

// DeleteWebhookByID implements Querier.DeleteWebhookByID.
func (q *DBQuerier) DeleteWebhookByID(ctx context.Context, webhookID pgtype.UUID) (DeleteWebhookByIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteWebhookByID")
	row := q.conn.QueryRow(ctx, deleteWebhookByIDSQL, webhookID)
	var item DeleteWebhookByIDRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.Secret, &item.Identifier, &item.VCSProviderID); err != nil {
		return item, fmt.Errorf("query DeleteWebhookByID: %w", err)
	}
	return item, nil
}

// DeleteWebhookByIDBatch implements Querier.DeleteWebhookByIDBatch.
func (q *DBQuerier) DeleteWebhookByIDBatch(batch genericBatch, webhookID pgtype.UUID) {
	batch.Queue(deleteWebhookByIDSQL, webhookID)
}

// DeleteWebhookByIDScan implements Querier.DeleteWebhookByIDScan.
func (q *DBQuerier) DeleteWebhookByIDScan(results pgx.BatchResults) (DeleteWebhookByIDRow, error) {
	row := results.QueryRow()
	var item DeleteWebhookByIDRow
	if err := row.Scan(&item.WebhookID, &item.VCSID, &item.Secret, &item.Identifier, &item.VCSProviderID); err != nil {
		return item, fmt.Errorf("scan DeleteWebhookByIDBatch row: %w", err)
	}
	return item, nil
}
