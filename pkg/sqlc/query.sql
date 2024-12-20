-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ?
);

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;


-- name: GetBookEagerLoadAuthor :many
SELECT b.id, b.title, b.author_id, a.name, a.bio
FROM books b
JOIN authors a ON b.author_id = a.id;



-- name: GetAuthorsEagerLoadBooks :many
SELECT sqlc.embed(a), sqlc.embed(b)
FROM authors a
JOIN books b ON b.author_id = a.id;