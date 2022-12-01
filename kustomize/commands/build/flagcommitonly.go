package build

import "github.com/spf13/pflag"

// AddFlagCommitOnly adds the --commit-only flag.
func AddFlagCommitOnly(set *pflag.FlagSet) {
	set.BoolVar(
		&theFlags.enable.commitOnly,
		"commit-only",
		false,
		"Enable use of the commit only mode.")
}

func isCommitOnlyEnabled() bool {
	return theFlags.enable.commitOnly
}
