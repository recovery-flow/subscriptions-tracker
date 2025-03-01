package cli

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/sirupsen/logrus"
)

func runMigration(ctx context.Context, cfg config.Config, direction string) error {
	cmd := exec.Command(
		"migrate",
		"-path", "internal/service/infra/data/repository/sqldb/migrations",
		"-database", cfg.Database.SQL.URL,
		"-verbose", direction,
	)

	cmd.Stdout = logrus.StandardLogger().Out
	cmd.Stderr = logrus.StandardLogger().Out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run migration %s: %v", direction, err)
	}
	return nil
}

func MigrateUp(ctx context.Context, cfg config.Config) error {
	return runMigration(ctx, cfg, "up")
}

func MigrateDown(ctx context.Context, cfg config.Config) error {
	return runMigration(ctx, cfg, "down")
}
