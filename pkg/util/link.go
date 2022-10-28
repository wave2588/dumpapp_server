package util

import (
	"fmt"
	"net/url"
	"regexp"

	"dumpapp_server/pkg/common/util"
	"github.com/pkg/errors"
)

type LinkParseHostPattern struct {
	Hosts    *Set
	Patterns []LinkParseURLPatten
}

type LinkParseURLPatten struct {
	Re      *regexp.Regexp
	Type    string
	SubType string
}

var InSiteLinkParsePatterns = []LinkParseHostPattern{
	{
		Hosts: NewSet("dumpapp.com", "sign.dumpapp.com"),
		Patterns: []LinkParseURLPatten{
			{Re: regexp.MustCompile(`^/user/(?P<id>\d+)/xx`), Type: "home", SubType: "home"},
		},
	},
}

var OutSitesLinkParsePatterns = []LinkParseHostPattern{
	{
		Hosts: NewSet("music.163.com"),
		Patterns: []LinkParseURLPatten{
			{Re: regexp.MustCompile(`/song/(?P<token>\d+)`), Type: "netease", SubType: "music"},
			{Re: regexp.MustCompile(`/m/song\?id=(?P<token>\d+)`), Type: "netease", SubType: "music"},
			{Re: regexp.MustCompile(`/#/song\?id=(?P<token>\d+)`), Type: "netease", SubType: "music"},
		},
	},
}

func ParseLinkInfo(link string) (linkType, subType, token string) {
	up, err := url.Parse(link)
	util.PanicIf(err)
	if up == nil {
		util.PanicIf(errors.New(fmt.Sprintf("ParseLinkInfo_up_is_nil  link: %s", link)))
	}

	for _, hostPattern := range InSiteLinkParsePatterns {
		if !hostPattern.Hosts.Exists(up.Host) {
			continue
		}
		for _, pattern := range hostPattern.Patterns {
			fmt.Println(up.Path)
			matchMap, hasMatch := RegexpNamedMatch(pattern.Re, up.Path)
			if hasMatch {
				return pattern.Type, pattern.SubType, matchMap["id"]
			}
		}
	}
	return "", "", ""
}
