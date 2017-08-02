package docker_parser

import (
	"testing"
	"github.com/stretchr/testify/require"
	"strings"
)

func TestTemplate(t *testing.T) {
	is := require.New(t)

	content := strings.Join([]string{
		"", // empty line
		"# This is a comment before everything",
		"   # THis is a comment with spaces before # and a comment after",
		"  fRoM toto # this is another comment",
		"FROM # this should be untouched",
		"",
	}, "\n")

	parser, err := Parse(content, false)

	is.NoError(err)
	is.Equal(content, parser.GetDockerFileContent())

}

func TestWithAll(t *testing.T) {

	is := require.New(t)

	content := strings.Join([]string{
		"TAG toto",
		"TAG titi",
		"CACHED_FROM cache_place",
		"CONTEXT ../",
		"FROM_FILE ./Dockerfile.ori",
		"rest of the file",
		"",
	}, "\n")

	parser, err := Parse(content, false)
	is.NoError(err)

	contentRes := strings.Join([]string{
		"# TAG toto",
		"# TAG titi",
		"# CACHED_FROM cache_place",
		"# CONTEXT ../",
		"# FROM_FILE ./Dockerfile.ori",
		"FROM " + parser.TmpTag,
		"rest of the file",
		"",
	}, "\n")

	is.Equal(parser.GetDockerFileContent(), contentRes)
	is.Equal(parser.Tags, []string{"toto", "titi"})
	is.Equal(parser.CachedFrom, "cache_place")
	is.Equal(parser.ContextPath, "../")
	is.Equal(parser.FromFile, "./Dockerfile.ori")

}
