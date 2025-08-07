package postgres

const (
	findExampleByID = "find example by id"
	createExample   = "create example"
	saveExample     = "update example"
)

func exampleQueries() map[string]string {
	return map[string]string{
		findExampleByID: `SELECT * FROM examples WHERE id = $1`,
		createExample:   `INSERT INTO examples (id, name) VALUES ($1, $2)`,
		saveExample:     `UPDATE quotas SET name = $1  WHERE id = $2`,
	}
}
