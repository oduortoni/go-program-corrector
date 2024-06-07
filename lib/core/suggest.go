package core

import (
	"regexp"
	"strings"
)

func suggest(s string) string {
	syntax := []Suggestion{
		{"package", []string{"package <name:string>"}},
		{"for", []string{"for i:=0; i<n:int>; i<++|+=n>", "<index:int>, <item:any> := range <collection:[]any> {"}},
		{"if", []string{"if [opt: !] <x:any> ( '>'|'<'|'=='|'>='|'<='|'!=' | <boolean>) <value> {"}},
		/*{"=", []string{"[-]", "[-]", ":[-]", "+[-]", "-[-]", "/[-]", "*[-]", "%[-]", "&[-]", "|[-]", "![-]"}},*/
	}

	for _, entry := range syntax {
		if strings.Contains(s, entry.Key) {
			for index, option := range entry.Suggest {
				if strings.Contains(option, "[-]") {
					entry.Suggest[index] = searchAndReplace(option, "[-]", entry.Key)
				}
			}
			return strings.Join(entry.Suggest, " or ")
		}
	}
	return s
}

func searchAndReplace(haystack, kneedle, replacer string) string {
	sb := []byte(haystack)
	re := regexp.MustCompile(kneedle)
	indices := re.FindIndex([]byte(haystack))
	if len(indices) == 2 {
		start := indices[0]
		if start > 0 {
			start -= 1
		}
		stop := indices[1]
		if stop != len(sb)-1 {
			stop += 1
		}
		ret := sb[:start]
		ret = append(ret, []byte(replacer)...)
		ret = append(ret, sb[stop+1:]...)
		return string(ret)
	}
	return haystack + "<< " + kneedle + " >>"
}
