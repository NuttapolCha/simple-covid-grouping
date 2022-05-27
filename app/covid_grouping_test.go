package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummarize(t *testing.T) {
	newInt := func(i int) *int {
		return &i
	}
	newStr := func(s string) *string {
		return &s
	}

	input := &CovidCaseResp{
		Data: []CovidCaseRespData{
			{
				Age:      newInt(20),
				Province: newStr("Bangkok"),
			},
			{
				Age:      newInt(1),
				Province: nil,
			},
			{
				Age:      newInt(3),
				Province: newStr("Pathumtani"),
			},
			{
				Age:      newInt(60),
				Province: newStr("Bangkok"),
			},
			{
				Age:      newInt(67),
				Province: nil,
			},
			{
				Age:      nil,
				Province: newStr("Bangkok"),
			},
			{
				Age:      newInt(15),
				Province: newStr("Bangkok"),
			},
			{
				Age:      newInt(54),
				Province: newStr("Bangkok"),
			},
			{
				Age:      newInt(42),
				Province: newStr("Pathumtani"),
			},
			{
				Age:      newInt(70),
				Province: newStr("Chaingmai"),
			},
			{
				Age:      newInt(36),
				Province: newStr("Chaingmai"),
			},
			{
				Age:      newInt(49),
				Province: newStr("Chaingmai"),
			},
		},
	}
	expected := &CovidCaseSummary{
		Province: map[string]int{
			"Bangkok":    5,
			"Pathumtani": 2,
			"Chaingmai":  3,
			"N/A":        2,
		},
		AgeGroup: map[ageGroup]int{
			junior:  4,
			mid:     5,
			senior:  2,
			unknown: 1,
		},
	}

	actual := input.Summarize()
	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual)
}
