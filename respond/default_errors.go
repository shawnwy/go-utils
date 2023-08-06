package respond

import "github.com/shawnwy/go-utils/v5/errors"

var (
	ErrServer        = errors.NewWithCode(10101, "server not available")
	ErrParamBind     = errors.NewWithCode(10102, "failed to bind parameters")
	ErrAuthorization = errors.NewWithCode(10103, "failed to authorization")
	ErrRBAC          = errors.NewWithCode(10104, "failed to access")
	ErrMongoConnect  = errors.NewWithCode(10105, "failed to connect mongo")
	ErrMySQLConnect  = errors.NewWithCode(10106, "failed to connect mysql")
	ErrRedisConnect  = errors.NewWithCode(10107, "failed to connect redis")
	ErrSocketConnect = errors.NewWithCode(10108, "failed to connect socket")
	ErrKafkaConnect  = errors.NewWithCode(10109, "failed to connect kafka")
	ErrNatsConnect   = errors.NewWithCode(10110, "failed to connect nats")
)
