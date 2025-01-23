package banana

import "errors"

var (
	UserConflict = errors.New("Nguoi dung da ton tai")
	SignUpFail   = errors.New("Dang ky that bai")
	UserNotFound = errors.New("Khong tim thay nguoi dung")
)
