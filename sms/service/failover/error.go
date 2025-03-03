package failover

import "errors"

var ErrAllFailed = errors.New("all sms service providers were polled, but failed to send")
