package gen

import "github.com/nicolerobin/sqlx_gen/parser"

type Join []string

// Key describes cache key
type Key struct {
	// VarLeft describes the variable of cache key expression which likes cacheUserIdPrefix
	VarLeft string
	// VarRight describes the value of cache key expression which likes "cache:user:id:"
	VarRight string
	// VarExpression describes the cache key expression which likes cacheUserIdPrefix = "cache:user:id:"
	VarExpression string
	// KeyLeft describes the variable of key definition expression which likes userKey
	KeyLeft string
	// KeyRight describes the value of key definition expression which likes fmt.Sprintf("%s%v", cacheUserPrefix, user)
	KeyRight string
	// DataKeyRight describes data key likes fmt.Sprintf("%s%v", cacheUserPrefix, data.User)
	DataKeyRight string
	// KeyExpression describes key expression likes userKey := fmt.Sprintf("%s%v", cacheUserPrefix, user)
	KeyExpression string
	// DataKeyExpression describes data key expression likes userKey := fmt.Sprintf("%s%v", cacheUserPrefix, data.User)
	DataKeyExpression string
	// FieldNameJoin describes the filed slice of table
	FieldNameJoin Join
	// Fields describes the fields of table
	Fields []*parser.Field
}
