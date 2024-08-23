package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func RequiresExactIntArgs(nArgs int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		nValidArgs := 0
		var err error
		for _, arg := range args {
			if _, err := strconv.Atoi(arg); err == nil {
				nValidArgs += 1
			}
		}
		if len(args) != nArgs || nValidArgs == 0 {
			err = errors.New(
				fmt.Sprintf(
					"%s: '%s' requires exactly %d int %s; got %d int %s.",
					cmd.Root().Name(),
					cmd.CommandPath(),
					nArgs,
					pluralize("argument", nArgs),
					nValidArgs,
					pluralize("argument", nValidArgs),
				),
			)
		}

		return err
	}
}

func RequiresMinArgs(min int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) >= min {
			return nil
		}
		return errors.New(
			fmt.Sprintf(
				"%s: '%s' requires at least %d %s; got %d %s.",
				cmd.Root().Name(),
				cmd.CommandPath(),
				min,
				pluralize("argument", min),
				len(args),
				pluralize("argument", len(args)),
			),
		)
	}
}

func RequiresRangeArgs(min int, max int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {

		if len(args) >= min && len(args) <= max {
			return nil
		}

		return errors.New(
			fmt.Sprintf(
				"%s: '%s' requires at least %d %s and at most %d %s; got %d %s",
				cmd.Root().Name(),
				cmd.CommandPath(),
				min,
				pluralize("argument", min),
				max,
				pluralize("argument", max),
				len(args),
				pluralize("argument", len(args)),
			),
		)
	}
}

func pluralize(w string, n int) string {
	if n <= 1 {
		return w
	}
	return w + "s"
}
