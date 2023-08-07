package user

import (
	"context"
	"fmt"
)

func (r *repository) VerifyUser(ctx context.Context, id uint64) error {
	query := r.db.Rebind(`UPDATE "auth".users SET is_verified = true WHERE id = ?`)
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("[UserRepo.VerifyUser][010013] Failed: %v", err)
	}

	return nil
}
