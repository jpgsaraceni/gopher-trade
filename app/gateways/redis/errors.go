package redis

import "errors"

var errType = errors.New("failed to convert redis reply to string")
