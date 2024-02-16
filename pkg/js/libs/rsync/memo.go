package rsync

import (
	"errors"
	"fmt"

	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/protocolstate"
)

func memoized(host string, port int) (IsRsyncResponse, error) {
	hash := "isRsync:" + host + ":" + fmt.Sprint(port)

	v, err, _ := protocolstate.Memoizer.Do(hash, func() (interface{}, error) {
		return isRsync(host, port)
	})
	if err != nil {
		return IsRsyncResponse{}, err
	}
	if value, ok := v.(IsRsyncResponse); ok {
		return value, nil
	}

	return IsRsyncResponse{}, errors.New("could not convert cached result to IsRsyncResponse")
}
