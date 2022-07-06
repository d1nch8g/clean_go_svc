package middleware

import (
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

var fieldsToExclude = []string{
	`password`,
	`login`,
	`bytes`,
}

func ctxTagInterceptor() grpc.UnaryServerInterceptor {
	return grpc_ctxtags.UnaryServerInterceptor(
		grpc_ctxtags.WithFieldExtractor(fieldLogger),
	)
}

func fieldLogger(fullMethod string, req any) map[string]interface{} {
	result := grpc_ctxtags.CodeGenRequestFieldExtractor(fullMethod, req)
	for _, v := range fieldsToExclude {
		delete(result, v)
	}
	return result
}
