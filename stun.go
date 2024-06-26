package util

import (
	utilString "github.com/linuxkungfu/go-util/string"
)

// STUN Attributes
// https://www.iana.org/assignments/stun-parameters/stun-parameters.xhtml

// stun
// https://www.rfc-editor.org/rfc/rfc8489.html

const (
	StunIceUfragLen int = 16
	StunIcePwdLen   int = 8
)

func CreateStunIceUsername() string {
	return utilString.CreateRandString(StunIceUfragLen)
}

func CreateStunIcePassword() string {
	return utilString.CreateRandString(StunIcePwdLen)
}
