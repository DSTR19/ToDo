package users_postgres_repository

import (
	core_postgres_pool "github.com/DSTR19/ToDo/internal/core/repository/postgres/pool"
)

type UsersPostgresRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersPostgresRepository(pool core_postgres_pool.Pool) *UsersPostgresRepository {
	return &UsersPostgresRepository{
		pool: pool,
	}
}
