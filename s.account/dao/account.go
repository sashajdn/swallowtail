package dao

import (
	"context"
	"time"

	"github.com/imdario/mergo"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"

	"swallowtail/libraries/gerrors"
	"swallowtail/s.account/domain"
	accountproto "swallowtail/s.account/proto"
)

// ListAccounts returns a list of all domain accounts from the underlying datastore.
func ListAccounts(ctx context.Context) ([]*domain.Account, error) {
	var (
		sql = `
		SELECT * FROM s_account_accounts`
		accounts []*domain.Account
	)
	err := db.Select(ctx, &accounts, sql)
	if err != nil {
		return nil, terrors.Propagate(err)
	}

	return accounts, nil
}

// ListFuturesMembers returns a list of all domain accounts from the underlying datastore, of whom are futures members.
func ListFuturesMembers(ctx context.Context) ([]*domain.Account, error) {
	var (
		sql = `
		SELECT * FROM s_account_accounts
		WHERE is_futures_member=true`
		accounts []*domain.Account
	)
	err := db.Select(ctx, &accounts, sql)
	if err != nil {
		return nil, terrors.Propagate(err)
	}

	return accounts, nil
}

// ReadAccountByUserID returns the a domain account from the underlying datastore, by UserID.
func ReadAccountByUserID(ctx context.Context, userID string) (*domain.Account, error) {
	var (
		sql      = `SELECT * FROM s_account_accounts WHERE user_id=$1`
		accounts []*domain.Account
	)

	err := db.Select(ctx, &accounts, sql, userID)
	if err != nil {
		return nil, gerrors.Propagate(err, gerrors.ErrUnknown, nil)
	}

	switch len(accounts) {
	case 0:
		return nil, gerrors.NotFound("account_not_found", nil)
	case 1:
		return accounts[0], nil
	}

	// We shouldn't have this case, since this is enforced at the persistence layer.
	slog.Critical(ctx, "More than one account with identical id", map[string]string{
		"userID": userID,
	})

	return accounts[0], nil
}

// ReadAccountByUsername returns the a domain account from the underlying datastore, by username.
func ReadAccountByUsername(ctx context.Context, username string) (*domain.Account, error) {
	var (
		sql      = `SELECT * FROM s_account_accounts WHERE username=$1`
		accounts []*domain.Account
	)
	err := db.Select(ctx, &accounts, sql, username)
	if err != nil {
		return nil, terrors.Propagate(err)
	}
	if len(accounts) > 1 {
		// We shouldn't have this case, since this is enforced at the persistence layer.
		slog.Error(ctx, "More than one account with identical username", map[string]string{
			"username": username,
		})
	}
	return accounts[0], nil
}

// CreateAccount creates a new account in the datastore.
func CreateAccount(ctx context.Context, account *domain.Account) error {
	var (
		sql = `
		INSERT INTO
			s_account_accounts(username, password, email, user_id, phone_number,
				created, updated, last_payment_timestamp,
				high_priority_pager, low_priority_pager, is_admin, is_futures_member) 
		VALUES
			($1, $2, $3 ,$4, $5, $6, $7, $8, $9, $10, $11, $12)`
	)

	var (
		lowPriorityPager  = account.LowPriorityPager
		highPriorityPager = account.HighPriorityPager
	)

	if lowPriorityPager == "" {
		lowPriorityPager = accountproto.PagerType_DISCORD.String()
	}
	if highPriorityPager == "" {
		highPriorityPager = accountproto.PagerType_DISCORD.String()
	}

	now := time.Now().UTC()
	if _, err := (db.Exec(
		ctx, sql,
		account.Username, account.Password, account.Email, account.UserID, account.PhoneNumber,
		now, now, now,
		highPriorityPager, lowPriorityPager, account.IsAdmin, account.IsFuturesMember,
	)); err != nil {
		return terrors.Propagate(err)
	}
	return nil
}

// UpdateAccount updates an already existing account; it will perform a mutation on the existing account.
// AccountID must be provided at least to the passed domain account struct, `mutation`.
func UpdateAccount(ctx context.Context, mutation *domain.Account) (*domain.Account, error) {
	var (
		sql = `
		UPDATE s_account_accounts
		SET username=$1, password=$2, email=$3, phone_number=$4, high_priority_pager=$5, low_priority_pager=$6, is_futures_member=$7, is_admin=$8, updated=$9
		WHERE user_id=$10`
	)
	if mutation.UserID == "" {
		return nil, terrors.PreconditionFailed("mutation-without-id", "Account mutation requires at least the account ID", nil)
	}

	account, err := ReadAccountByUserID(ctx, mutation.UserID)
	if err != nil {
		return nil, err
	}

	if err := mergo.MergeWithOverwrite(account, mutation); err != nil {
		return nil, gerrors.Augment(gerrors.Propagate(err, gerrors.ErrUnknown, nil), "failed_to_merge_accounts", nil)
	}

	if _, err := (db.Exec(
		ctx, sql,
		account.Username, account.Password, account.Email, account.PhoneNumber, account.HighPriorityPager, account.LowPriorityPager, account.IsFuturesMember, account.IsAdmin, account.Updated,
		account.UserID,
	)); err != nil {
		return nil, gerrors.Propagate(err, gerrors.ErrUnknown, nil)
	}
	return account, nil
}
