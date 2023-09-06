package types

import (
	"errors"
	"regexp"
)

type Trigger struct {
	Key   string
	Value string
}

type Triggers []Trigger

const (
	TriggerKeyRegexpStr = `([A-Za-z]*)`
	TriggerValRegexpStr = `([A-Za-z0-9]*)`
	TriggerSep          = "/"
	TriggerRegexpStr    = `^` + TriggerKeyRegexpStr + TriggerSep + TriggerValRegexpStr + `$`
)

var TriggerRegexp = regexp.MustCompile(TriggerRegexpStr)

func ParseTriggerFromBytes(bytes []byte) (Trigger, error) {
	str := string(bytes)
	res := TriggerRegexp.FindStringSubmatch(str)

	if len(res) != 3 {
		return Trigger{}, errors.New("cannot parse trigger bytes")
	}
	return Trigger{Key: res[1], Value: res[2]}, nil
}
