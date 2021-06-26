package handler

import (
	"swallowtail/s.account/domain"
	accountproto "swallowtail/s.account/proto"

	"github.com/monzo/terrors"
)

func getIdentifierFromAccount(account *domain.Account, pagerType string) (string, error) {
	errParams := map[string]string{
		"account_id": account.AccountID,
	}

	switch pagerType {
	case accountproto.PagerType_DISCORD.String():
		if account.DiscordID == "" {
			return "", terrors.PreconditionFailed("account-discord-id-missing", "Cannot page account via discord; discord id not set", errParams)
		}
		return account.DiscordID, nil
	case accountproto.PagerType_EMAIL.String():
		if account.Email == "" {
			return "", terrors.PreconditionFailed("account-email-missing", "Cannot page account via email; email not set", errParams)
		}
		return account.Email, nil
	case accountproto.PagerType_PHONE.String():
		if account.PhoneNumber == "" {
			return "", terrors.PreconditionFailed("account-phone-number-missing", "Cannot page account via phone number; phone number not set", errParams)
		}
		return account.PhoneNumber, nil
	case accountproto.PagerType_SMS.String():
		if account.PhoneNumber == "" {
			return "", terrors.PreconditionFailed("account-phone-number-missing", "Cannot page account via sms; sms not set", errParams)
		}
		return account.PhoneNumber, nil
	}

	errParams["pager_type"] = pagerType
	return "", terrors.PreconditionFailed("unknown-pager-type", "Cannot page account; unknown pager type", errParams)
}
