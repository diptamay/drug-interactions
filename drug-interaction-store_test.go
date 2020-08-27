package main

import "testing"

var testCases = []struct {
	drugs    []string
	expected string
}{
	{[]string{"sildenafil", "tamsulosin", "valaciclovir"}, "MODERATE: Sildenafil may potentiate the hypotensive effect of alpha blockers, resulting in symptomatic hypotension in some patients."},
	{[]string{"sildenafil", "ibuprofen"}, "No interaction"},
	{[]string{"valaciclovir", "doxepin", "ticlopidine", "ibuprofen"}, "MAJOR: Valaciclovir may decrease the excretion rate of Doxepin which could result in a higher serum level."},
	{[]string{"abc", "enalaprilat"}, "No interaction"},
	{[]string{"abc", "xyz"}, "No interaction"},
	{[]string{"sildenafil", "nicardipine", "enalaprilat"}, "MODERATE: Sildenafil may increase the antihypertensive activities of Enalaprilat."},
	{[]string{"sildenafil", "nicardipine", "enalaprilat", "echinacea"}, "MAJOR: The metabolism of Sildenafil can be increased when combined with Echinacea."},
	{[]string{"sildenafil", "nicardipine", "enalaprilat", "norethisterone", "echinacea"}, "MAJOR: The metabolism of Sildenafil can be increased when combined with Echinacea."},
	{[]string{"SILDENAFIL", "nicardipine", "enalaprilat"}, "MODERATE: Sildenafil may increase the antihypertensive activities of Enalaprilat."},
}

func TestGetImpact(t *testing.T) {
	const filePath = "./interactions.json"
	var interactions = Interactions{}
	_ = interactions.BuildStore(filePath)

	for _, test := range testCases {
		interaction, _ := interactions.FindDrugsImpact(test.drugs)

		observed := GetImpactString(interaction)
		if observed == test.expected {
			t.Logf("PASS: %v", test.drugs)
		} else {
			t.Errorf("FAIL: %v\nExpected: %v\n, Actual: %v\n", test.drugs, test.expected, observed)
		}
	}
}
