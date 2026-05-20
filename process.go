package ninerouter

import (
	"context"
	"os"
	"os/exec"
	"strconv"
)

type StartOptions struct {
	Bin        string
	Port       int
	Host       string
	NoBrowser  bool
	SkipUpdate bool
	Tray       bool
	Log        bool
	Env        []string
}

func Start(ctx context.Context, opts StartOptions) (*exec.Cmd, error) {
	bin := opts.Bin
	if bin == "" {
		bin = "9router"
	}
	args := []string{}
	if opts.Port > 0 {
		args = append(args, "--port", strconv.Itoa(opts.Port))
	}
	if opts.Host != "" {
		args = append(args, "--host", opts.Host)
	}
	if opts.NoBrowser {
		args = append(args, "--no-browser")
	}
	if opts.SkipUpdate {
		args = append(args, "--skip-update")
	}
	if opts.Tray {
		args = append(args, "--tray")
	}
	if opts.Log {
		args = append(args, "--log")
	}
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Env = append(os.Environ(), opts.Env...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}
