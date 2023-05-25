package enumx

type UserStatus int

const (
	DISABLED UserStatus = iota
	NORMAL
)

// CheckValueInUserStatusEnum 检查传入的值是否是枚举值
func CheckValueInUserStatusEnum(val int64) bool {
	all := []UserStatus{DISABLED, NORMAL}
	ok := false
	for value := range all {
		if value == int(val) {
			ok = true
			break
		}
	}
	return ok
}
