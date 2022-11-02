package postgres

const InsertPetMasterQuery = `
INSERT INTO pet_masters (name, contact_number, user_id) VALUES ($1, $2, $3)
RETURNING id, name, contact_number, user_id
`
