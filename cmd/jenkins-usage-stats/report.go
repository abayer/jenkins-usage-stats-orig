package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	sq "github.com/Masterminds/squirrel"
	stats "github.com/abayer/jenkins-usage-stats"
	"github.com/spf13/cobra"
)

// ReportOptions contains the configuration for actually outputting reports
type ReportOptions struct {
	Directory string
	Database  string
}

// NewReportCmd returns the report command
func NewReportCmd(ctx context.Context) *cobra.Command {
	options := &ReportOptions{}

	cobraCmd := &cobra.Command{
		Use:   "report",
		Short: "Generate stats.jenkins.io reports",
		Run: func(cmd *cobra.Command, args []string) {
			if err := options.runReport(ctx); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		DisableAutoGenTag: true,
	}

	cobraCmd.Flags().StringVar(&options.Database, "database", "", "Database URL to import to")
	_ = cobraCmd.MarkFlagRequired("database")
	cobraCmd.Flags().StringVar(&options.Directory, "directory", "", "Directory to output to")
	_ = cobraCmd.MarkFlagRequired("directory")

	return cobraCmd
}

func (ro *ReportOptions) runReport(ctx context.Context) error {
	rawDB, err := sql.Open("postgres", ro.Database)
	if err != nil {
		return err
	}
	defer func() {
		_ = rawDB.Close()
	}()

	db := sq.NewStmtCacheProxy(rawDB)

	now := time.Now()

	return stats.GenerateReport(db, now.Year(), int(now.Month()), ro.Directory)
}