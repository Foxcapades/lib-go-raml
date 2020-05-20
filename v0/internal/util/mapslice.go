package util

import "gopkg.in/yaml.v2"

func FindKey(mp yaml.MapSlice, key interface{}) (interface{}, bool) {
	for i := range mp {
		if mp[i].Key == key {
			return mp[i].Value, true
		}
	}
	return nil, false
}
