package mywife

import (
	"testing"

	"github.com/li-peifeng/metatube-sdk-go/provider/internal/testkit"
)

func TestMyWife_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"1542",
		"1882",
		"1888",
	})
}
