package core

import (
	"regexp"
	"strings"
)

func suggest(s string) string {
	syntax := map[string][]string{
		"package": {"package <name:string>"},
		"for":     {"for i:=0; i<n:int>; i<++|+=n>", "<index:int>, <item:any> := range <collection:[]any> {"},
		"if":      {"if [opt: !] <x:any> ( '>'|'<'|'=='|'>='|'<='|'!=' | <boolean>) <value> {"},
		"=":       {"[-]", "[-]", ":[-]", "+[-]", "-[-]", "/[-]", "*[-]", "%[-]", "&[-]", "|[-]", "![-]"},
	}

	for key, entry := range syntax {
		if strings.Contains(s, key) {
			for index, option := range entry {
				if strings.Contains(option, "[-]") {
					entry[index] = searchAndReplace(option, "[-]", key)
				}
			}
			return strings.Join(entry, " or ")
		}
	}
	return s
}

func searchAndReplace(haystack, kneedle, replacer string) string {
	sb := []byte(haystack)
	re := regexp.MustCompile(kneedle)
	indices := re.FindIndex([]byte(haystack))
	if len(indices) == 2 {
		print("ISAYA\n")
		start := indices[0]
		if start > 0 { start -= 1 }
		stop := indices[1]
		if stop != len(sb)-1 { stop += 1 }
		ret := sb[:start]
		ret = append(ret, []byte(replacer)...) 
		ret = append(ret, sb[stop+1:]...)
		return string(ret)
	}
	return haystack + "<< " + kneedle + " >>"
}
