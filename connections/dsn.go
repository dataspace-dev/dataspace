package connections

import "dataspace/db/types"

// BuildDsn builds the DSN string from the given db.Conection struct
func BuildDsn(dbData *types.Connection) string {
	return "host=" + dbData.Host + " port=" + dbData.Port + " dbname=" + dbData.Dbname + " user=" + dbData.User + " password=" + dbData.Pass + " sslmode=" + dbData.SSLMode
}
