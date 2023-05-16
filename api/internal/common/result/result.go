package result

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pure-go-zero-admin/api/internal/common/errorx"
)

// Response is the default http result, handling response object and its error.
func Response(w http.ResponseWriter, r *http.Request, err error, data ...any) {
	var (
		ctx     = r.Context()
		resp    any
		code    int
		message string
	)

	if err != nil {
		// Error returned
		if e, ok := err.(*errorx.CodeError); ok {
			code = e.Code
		} else if e, ok := status.FromError(err); ok {
			if v := e.Code(); v == codes.DeadlineExceeded {
				code = http.StatusGatewayTimeout
			} else {
				code = int(v)
			}
		}

		switch code {
		case http.StatusPreconditionFailed:
			message = err.Error()
		default:
			if v, ok := errorx.Message(code); ok {
				message = v
			} else {
				code = http.StatusInternalServerError
				message, _ = errorx.Message(http.StatusInternalServerError)
			}
		}

		resp = &errorx.CodeError{
			Code:      code,
			Message:   message,
			Method:    r.Method,
			Path:      r.RequestURI,
			Timestamp: time.Now(),
		}

		logx.WithContext(ctx).Slowf("【API-SLOW】request : %s, path : %s, err : %s", r.Method, r.RequestURI, err.Error())
	} else {
		// Success returned
		code = http.StatusOK

		if len(data) == 0 {
			resp = true
		} else if len(data) == 1 {
			resp = data[0]
		} else {
			resp = data
		}

		logx.WithContext(ctx).Infof("【API-INFO】request : %s, path : %s, ok", r.Method, r.RequestURI)
	}

	httpx.WriteJsonCtx(ctx, w, code, resp)
}
