package simple

type Database struct {
	Name string
}

type PostgresSQL Database
type MongoDB Database

func NewDatabasePostgresSQL() *PostgresSQL {
	return (*PostgresSQL)(&Database{Name: "postgresSQl"})
}

func NewDatabaseMongoDB() *MongoDB {
	return (*MongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	postgresSQL *PostgresSQL
	mongoDb     *MongoDB
}

func NewDatabaseRepository(postgresSQL *PostgresSQL, mongoDb *MongoDB) *DatabaseRepository {
	return &DatabaseRepository{postgresSQL: postgresSQL, mongoDb: mongoDb}
}
