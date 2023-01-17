package parser

import (
	"TVHelper/common"
	"encoding/json"
	"log"
	"strconv"

	"github.com/spf13/cast"
)

type Parser struct {
	Subscribe      []Subscribe    `json:"subscribe,omitempty"`
	MixAds         []string       `json:"mix-ads,omitempty"`
	MixFlags       []string       `json:"mix-flags,omitempty"`
	MixParses      []common.Parse `json:"mix-parses,omitempty"`
	SitesBlacklist []string       `json:"sites-blacklist,omitempty"`
	SitesAppend    []common.Site  `json:"sites-append,omitempty"`
	SitesPrepend   []common.Site  `json:"sites-prepend,omitempty"`
	Lives          []common.Live  `json:"lives,omitempty"`
	Spider         string         `json:"spider,omitempty"`
	Wallpaper      string         `json:"wallpaper,omitempty"`
}

type Subscribe struct {
	Url      string `json:"url,omitempty"`
	MultiJar bool   `json:"multi-jar,omitempty"`
	Jar      string `json:"jar,omitempty"`
	AlwaysOn bool   `json:"always-on,omitempty"`
}

func (s *Subscribe) UnmarshalJSON(data []byte) error {
	type TempSubscribe Subscribe
	sr := struct {
		*TempSubscribe
		MultiJar interface{} `json:"multi-jar,omitempty"`
		AlwaysOn interface{} `json:"always-on,omitempty"`
	}{
		TempSubscribe: (*TempSubscribe)(s),
	}
	if err := json.Unmarshal(data, &sr); err != nil {
		return err
	}
	switch sr.MultiJar.(type) {
	case bool:
		s.MultiJar = sr.MultiJar.(bool)
	case string:
		boolValue, err := strconv.ParseBool(sr.MultiJar.(string))
		if err != nil {
			log.Println(err)
		}
		s.MultiJar = boolValue
	default:
		s.MultiJar = false
		if cast.ToInt(sr.MultiJar) > 0 {
			s.MultiJar = true
		}
	}
	switch sr.AlwaysOn.(type) {
	case bool:
		s.AlwaysOn = sr.AlwaysOn.(bool)
	case string:
		boolValue, err := strconv.ParseBool(sr.AlwaysOn.(string))
		if err != nil {
			log.Println(err)
		}
		s.AlwaysOn = boolValue
	default:
		s.AlwaysOn = false
		if cast.ToInt(sr.AlwaysOn) > 0 {
			s.AlwaysOn = true
		}
	}
	return nil
}
