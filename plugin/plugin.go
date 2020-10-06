// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	Debug bool   `envconfig:"PLUGIN_DEBUG"` // Debug enables Netlify debugging.
	Path  string `envconfig:"PLUGIN_PATH"`  // Path provides the upload path.
	Site  string `envconfig:"PLUGIN_SITE"`  // Side provides the Netlify site id.
	Token string `envconfig:"PLUGIN_TOKEN"` // Token provides the Netlify token.
	Prod  bool   `envconfig:"PLUGIN_PROD"`  // Prod instructs netlify to deploy to prod.

	Level string `envconfig:"PLUGIN_LOG_LEVEL"` // Level defines the log level.
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	if args.Token == "" {
		return errors.New("Missing Netlify Token")
	}

	if args.Site == "" {
		return errors.New("Missing Netlify Site")
	}

	envs := []string{
		fmt.Sprintf("NETLIFY_AUTH_TOKEN=%s", args.Token),
		fmt.Sprintf("NETLIFY_SITE_ID=%s", args.Site),
	}

	// Print the version number fo rht ecommand line tools
	cmd := exec.Command("netlify", "--version")
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	trace(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	flags := []string{"deploy"}
	if args.Debug {
		flags = append(flags, "--debug")
	}
	if args.Path != "" {
		flags = append(flags, fmt.Sprintf("--dir=%s", args.Path))
	} else {
		flags = append(flags, fmt.Sprintf("--dir=./"))
	}

	cmd = exec.Command("netlify", flags...)
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	trace(cmd)
	return cmd.Run()
}

// trace prints the command to the stdout.
func trace(cmd *exec.Cmd) {
	fmt.Printf("$ %s\n", strings.Join(cmd.Args, " "))
}
