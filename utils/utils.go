package utils

import "github.com/diubrother/gorpc/codes"

func parseTarget(target string) error{
	if target == "" {
		return codes.CONFIGERROR
	}
}