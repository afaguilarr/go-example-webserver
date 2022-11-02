package postgres

const InsertLocationQuery = `
INSERT INTO locations
  (country, state_or_province, city_or_municipality, neighborhood, zip_code, pet_master_id)
  VALUES ($1, $2, $3, $4, $5, $6)
RETURNING country, state_or_province, city_or_municipality, neighborhood, zip_code, pet_master_id
`
