-- name: GetEmployee :one
SELECT * FROM employee
WHERE id=coalesce($1, id) LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employee;

-- name: CreateEmployee :one
INSERT INTO employee (
    id, guid, first_name, last_name, patronymic, position_code, photo_guid, updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id=$1;

-- name: UpdateEmployee :one
UPDATE employee SET (
    guid, first_name, last_name, patronymic, position_code, photo_guid, updated
) = (
    coalesce($2, guid), coalesce($3, first_name), coalesce($4, last_name), coalesce($5, patronymic),
    coalesce($6, position_code), coalesce($7, photo_guid), coalesce($8, updated)
)
WHERE id=$1
RETURNING *;