

-- name: CreateUser :one
INSERT INTO public.users (name, email, role, password_hash)
	VALUES (@name::text, @email::text, @role::smallint, @password_hash::text)
	RETURNING id;


-- name: GetUser :one
SELECT *
	FROM public.users
	WHERE
		id = @id
		AND NOT is_deleted;


-- name: UpdateUser :exec
UPDATE public.users
	SET
		name = COALESCE( NULLIF( @name::text, '' ), name ),
		email = COALESCE( NULLIF( @email::text, '' ), email ),
		role = COALESCE( NULLIF( @role::smallint, 0 ), role ),
		update_timestamp = NOW()
	WHERE
		id = @id
		AND NOT is_deleted;


-- name: DeleteUser :exec
UPDATE public.users
	SET
		is_deleted = TRUE,
		delete_timestamp = NOW()
	WHERE
		id = @id
		AND NOT is_deleted;

