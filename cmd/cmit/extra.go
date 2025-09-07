package cmit

var ExtraArgs []string

func init() {
	rootCmd.Flags().StringSliceVar(&ExtraArgs, "extra", []string{}, "Additional git commit flags")
}
