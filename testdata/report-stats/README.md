## report-stats

This directory contains daily gzipped data to be ingested for report function testing. If any changes are made here, the directory should be ingested into a fresh database and then dumped to `testdata/fixtures` by running `make dump-fixtures`.

Any changes to `testdata/fixtures` will mean that `report_test.go`'s test will need to be re-run with the `UPDATE_GOLDEN` env var set to `true` to update the contents of `testdata/reports` as well.
