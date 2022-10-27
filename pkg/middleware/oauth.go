package middleware

import (
	"context"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/middleware/util"
)

func SetTicketCookie(w http.ResponseWriter, r *http.Request, ticket string) {
	util.SetCookie(w, "session", map[string]string{
		"ticket": ticket,
	})
}

/// 招聘的代理走这个 ops
func OAuthAdminV2(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		/// 判断是否是调试
		name := r.Header.Get(constant.AppOpsAuthNameHeaderKey)
		if memberID, ok := constant.OpsAuthNameMap[name]; ok {
			ctx := context.WithValue(r.Context(), constant.MemberIDKey, memberID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		registerTicket := util.GetCookie(r, "session")["ticket"]
		if registerTicket == "" {
			panic(errors.ErrNotAuthorized)
		}
		ticket, err := util2.ParseTicket(registerTicket)
		util.PanicIf(err)
		if ticket.MemberID == 0 {
			panic(errors.ErrNotAuthorized)
		}
		_, isAdmin := constant.OpsAuthMemberIDMap[ticket.MemberID]
		_, isAdminV2 := constant.OpsAuthMemberIDMapV2[ticket.MemberID]
		/// 说明是超级管理员  说明是招聘的管理员
		if (isAdmin && !isAdminV2) || (!isAdmin && isAdminV2) {
			ctx = context.WithValue(ctx, constant.MemberIDKey, ticket.MemberID)
		} else {
			panic(errors.ErrMemberAccessDenied)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func OAuthAdmin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		/// 判断是否是调试
		name := r.Header.Get(constant.AppOpsAuthNameHeaderKey)
		if memberID, ok := constant.OpsAuthNameMap[name]; ok {
			ctx := context.WithValue(r.Context(), constant.MemberIDKey, memberID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		registerTicket := util.GetCookie(r, "session")["ticket"]
		if registerTicket == "" {
			panic(errors.ErrNotAuthorized)
		}
		ticket, err := util2.ParseTicket(registerTicket)
		util.PanicIf(err)
		if ticket.MemberID == 0 {
			panic(errors.ErrNotAuthorized)
		}
		if _, ok := constant.OpsAuthMemberIDMap[ticket.MemberID]; !ok {
			panic(errors.ErrMemberAccessDenied)
		}
		ctx := context.WithValue(r.Context(), constant.MemberIDKey, ticket.MemberID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func OAuthRegister(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		/// 判断是否是调试
		name := r.Header.Get(constant.AppOpsAuthNameHeaderKey)
		if memberID, ok := constant.OpsAuthNameMap[name]; ok {
			ctx := context.WithValue(r.Context(), constant.MemberIDKey, memberID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		registerTicket := util.GetCookie(r, "session")["ticket"]
		if registerTicket == "" {
			panic(errors.ErrNotAuthorized)
		}
		ticket, err := util2.ParseTicket(registerTicket)
		util.PanicIf(err)
		if ticket.MemberID == 0 {
			panic(errors.ErrNotAuthorized)
		}
		ctx := context.WithValue(r.Context(), constant.MemberIDKey, ticket.MemberID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func OAuthGuest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		registerTicket := util.GetCookie(r, "session")["ticket"]
		if registerTicket != "" {
			ticket, _ := util2.ParseTicket(registerTicket)
			ctx = context.WithValue(r.Context(), constant.MemberIDKey, ticket.MemberID)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetMemberID(ctx context.Context) (int64, error) {
	if ctx == nil {
		return -1, errors.ErrInvalidTicket
	}
	if MemberID, ok := ctx.Value(constant.MemberIDKey).(int64); ok {
		return MemberID, nil
	}
	return -1, errors.ErrInvalidTicket
}

func MustGetMemberID(ctx context.Context) int64 {
	if memberID, err := GetMemberID(ctx); err != nil {
		panic(err)
	} else {
		if memberID <= 0 {
			panic(errors.ErrInvalidTicket)
		}
		return memberID
	}
}
