package errors

import (
	"fmt"
	"github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// HandleGRPCError извлекает код и сообщение из ErrorResponse
func HandleGRPCError(err error) (int32, string, error) {
	st, ok := status.FromError(err)
	if !ok {
		return 0, "Unknown error", fmt.Errorf("not a gRPC error: %v", err)
	}

	if len(st.Details()) > 0 {
		for _, details := range st.Details() {
			switch detail := details.(type) {
			case *anypb.Any:
				var errResponse error_v1.ErrorResponse
				if err = proto.Unmarshal(detail.Value, &errResponse); err == nil {
					return errResponse.Code, errResponse.Message, nil
				} else {
					return 0, "", fmt.Errorf("failed to unmarshal error details: %v", err)
				}
			default:
				return 0, "", fmt.Errorf("unknown error detail type")
			}
		}
	}

	return int32(st.Code()), st.Message(), nil
}
