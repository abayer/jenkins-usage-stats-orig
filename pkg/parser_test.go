package pkg_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/abayer/jenkins-usage-stats/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDailyJSON(t *testing.T) {
	fooFile := filepath.Join("testdata", "foo.json.gz")

	reports, err := pkg.ParseDailyJSON(fooFile)
	require.NoError(t, err)
	assert.Len(t, reports, 2)
	assert.Equal(t, "32b68faa8644852c4ad79540b4bfeb1caf63284811f4f9d6c2bc511f797218c8", reports[0].Install)
	assert.Equal(t, uint64(50), reports[0].Jobs["hudson-maven-MavenModuleSet"])
	assert.Len(t, reports[0].Plugins, 75)
	assert.Equal(t, "1.8", reports[0].Nodes[0].JVMVersion)
	assert.Equal(t, "11", reports[1].Nodes[0].JVMVersion)

	ts, err := reports[0].Timestamp()
	require.NoError(t, err)
	assert.Equal(t, time.Date(2021, time.October, 30, 23, 59, 54, 0, time.UTC), ts)
}

func TestFilterPrivateFromReport(t *testing.T) {
	report := &pkg.JSONReport{
		Plugins: []pkg.JSONPlugin{
			{
				Name:    "legit-plugin",
				Version: "1.2.3",
			},
			{
				Name:    "privateplugin-something",
				Version: "1.2.3",
			},
			{
				Name:    "other-legit-plugin",
				Version: "2.3.4 (private)",
			},
			{
				Name:    "final-legit-plugin",
				Version: "4.5.6",
			},
		},
	}

	pkg.FilterPrivateFromReport(report)

	assert.Len(t, report.Plugins, 2)
	assert.Equal(t, report.Plugins[0].Name, "legit-plugin")
	assert.Equal(t, report.Plugins[1].Name, "final-legit-plugin")
}