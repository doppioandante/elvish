package edit

import (
	"testing"

	"github.com/elves/elvish/edit/ui"
)

var (
	theHistList = newHistlist([]string{"ls", "echo lalala", "ls"})

	histlistDedupFilterTests = []listingFilterTestCases{
		{"", []shown{
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
		{"l", []shown{
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
	}

	histlistNoDedupFilterTests = []listingFilterTestCases{
		{"", []shown{
			{"0", ui.Unstyled("ls")},
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
		{"l", []shown{
			{"0", ui.Unstyled("ls")},
			{"1", ui.Unstyled("echo lalala")},
			{"2", ui.Unstyled("ls")}}},
	}
)

func TestHistlist(t *testing.T) {
	testListingFilter(t, "theHistList", theHistList, histlistDedupFilterTests)
	theHistList.provider.(*histlist).toggleDedup()
	testListingFilter(t, "theHistList", theHistList, histlistNoDedupFilterTests)
}
