//This will parse docker for some specific values
// We are cheating on the parsing to make it simplier and because we need only FROM and above (we should not have above)
// So no complicated json or run parsing

package docker_parser

import (
	"strings"
	"regexp"
	"errors"
	"math/rand"
	"strconv"
	"fmt"
)

type Parser struct {
	content           string   // original content
	Tags              []string // list of Tags from the command TAG
	FromFile          string   // FROM_FILE
	CachedFrom        string   // CACHED_FROM
	ContextPath       string   // CONTEXT
	TmpTag            string
	DockerfileContent string // The docker file without thoses parameters
}

func Parse(content string, debug bool) (res Parser, err error) {
	parser := Parser{
		content,
		[]string{},
		"",
		"",
		"",
		"docker_extended_tmp/" + strconv.FormatInt(rand.Int63(), 16),
		"",
	}

	newContent := []string{}
	oriContentSplited := strings.Split(content, "\n")

loop:
	for i, line := range oriContentSplited {
		withoutComment := removeComments(line)
		cleaned := strings.TrimSpace(withoutComment)
		firstWord, rest := extractWordAndValue(cleaned)

		switch firstWord {
		case "":
			// the word is empty, we add the original line
			newContent = append(newContent, line)
			break
		case "FROM":
			newContent = append(newContent, line)
			newContent = append(newContent, oriContentSplited[i+1:]...)
			break loop
		case "FROM_FILE":
			if rest == "" {
				return parser, errors.New("FROM FILE expect a parameter")
			}
			parser.FromFile = rest
			newContent = append(newContent, "# "+line)
			newContent = append(newContent, "FROM "+parser.TmpTag)
			newContent = append(newContent, oriContentSplited[i+1:]...)
			break loop
		case "CACHED_FROM":
			if rest == "" {
				return parser, errors.New("CACHED_FROM need a value, found empty string")
			}
			if parser.CachedFrom != "" {
				return parser, errors.New("Only one call of CACHED_FROM is allowed")
			}
			parser.CachedFrom = rest
			newContent = append(newContent, "# "+line, )
		case "TAG":
			if rest == "" {
				return parser, errors.New("TAG need a value, found empty string")
			}
			parser.Tags = append(parser.Tags, rest)
			newContent = append(newContent, "# "+line)
		case "CONTEXT":
			if rest == "" {
				return parser, errors.New("CONTEXT need a value, found empty string")
			}
			if parser.ContextPath != "" {
				return parser, errors.New("Only one call of CONTEXT is allowed")
			}
			parser.ContextPath = rest
			newContent = append(newContent, "# "+line)
		default:
			return parser, errors.New("Command " + firstWord + " appears before the FROM or FROM_FILE command")
		}

	}

	parser.DockerfileContent = strings.Join(newContent, "\n")

	if debug {
		fmt.Println("#####################")
		fmt.Println("Dockerfile to execute")
		fmt.Println("#####################")
		fmt.Println(parser.DockerfileContent)
	}


	return parser, nil

}

func removeComments(content string) string {
	return regexp.MustCompile(`#.*`).ReplaceAllString(content, "")
}

func extractWordAndValue(content string) (firstWord string, rest string) {
	// Extract the first word and a value. It expect to have only a single value
	firstWordReg := regexp.MustCompile(`^[a-zA-Z_]+`)
	firstWord = strings.ToUpper(firstWordReg.FindString(content))
	rest = strings.TrimSpace(firstWordReg.ReplaceAllString(content, ""))
	return
}

func (p Parser) GetDockerFileContent() string {
	return p.DockerfileContent
}
