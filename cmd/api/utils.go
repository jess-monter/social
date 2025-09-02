package main

import "strconv"

func parseID(param string) (int64, error) {
	intID, err := strconv.ParseInt(param, 10, 64)
	if err != nil || intID < 1 {
		return 0, err
	}
	return intID, nil
}
