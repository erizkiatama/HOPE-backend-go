package auth

import (
	"HOPE-backend/internal/entity/auth"
	"context"
	"fmt"
)

func (r *repository) CreateUser(ctx context.Context, user auth.User) (*auth.User, error) {
	rows, err := r.db.NamedQueryContext(ctx,
		`WITH new_user AS (
				INSERT INTO "auth".users (email, password, name, alias, is_verified, secret_key) 
				VALUES (:email, :password, :name, :alias, :is_verified, :secret_key) RETURNING id
		    ),
			profile AS (
			    INSERT INTO "auth".user_profiles (photo, user_id, total_audio_played, total_time_played, longest_streak)
				VALUES (:photo, (SELECT id from new_user), :total_audio_played, :total_time_played, :longest_streak)
			)
			INSERT INTO "auth".user_roles (user_id, role_name)
			VALUES ((SELECT id from new_user), :role)
			RETURNING user_id AS id`,
		user,
	)
	if err != nil {
		return nil, fmt.Errorf("[AuthRepo.Create][010006] Failed: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&user.Id)
		if err != nil {
			return nil, fmt.Errorf("[AuthRepo.Create][010007] Failed: %v", err)
		}
	}
	return &user, nil
}
