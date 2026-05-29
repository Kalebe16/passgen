package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

const uppercaseChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowercaseChars string = "abcdefghijklmnopqrstuvwxyz"
const numberChars string = "0123456789"
const symbolChars string = "!@#$%^&*()-_=+"

type Options struct {
	Version   bool
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Symbols   bool
}

func generatePassword(length int, uppercase bool, lowercase bool, numbers bool, symbols bool) string {
	var chars string = ""
	var password string = ""

	if uppercase {
		chars += uppercaseChars
	}
	if lowercase {
		chars += lowercaseChars
	}
	if numbers {
		chars += numberChars
	}
	if symbols {
		chars += symbolChars
	}

	for range length {
		password += string(chars[rand.Intn(len(chars))])
	}

	return password
}

func validateOptions(opts Options) error {
	if opts.Version {
		return nil
	}

	if opts.Length < 1 || opts.Length > 1000 {
		return errors.New("error: '--length' must be >= 1 and <= 1000")
	}

	if !opts.Uppercase && !opts.Lowercase && !opts.Numbers && !opts.Symbols {
		return errors.New("error: one of '--uppercase', '--lowercase', '--numbers', '--symbols' is required")
	}

	return nil
}

func run(opts Options, copy func(string) error) error {
	if opts.Version {
		fmt.Fprintln(os.Stdout, "v"+Version)
		return nil
	}

	var password string = generatePassword(opts.Length, opts.Uppercase, opts.Lowercase, opts.Numbers, opts.Symbols)
	if err := copy(password); err != nil {
		fmt.Fprintln(os.Stderr, "error: failed to copy password to clipboard")
		return err
	}
	fmt.Fprintln(os.Stdout, "password copied to clipboard")
	return nil
}

func cli() {
	var opts Options

	rootCmd := &cobra.Command{
		Use:           "passgen",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateOptions(opts)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts, clipboard.WriteAll)
		},
	}
	rootCmd.Flags().BoolVarP(&opts.Version, "version", "v", false, "show version")
	rootCmd.Flags().IntVarP(&opts.Length, "length", "l", 16, "password length")
	rootCmd.Flags().BoolVarP(&opts.Uppercase, "uppercase", "u", false, "include uppercase chars")
	rootCmd.Flags().BoolVarP(&opts.Lowercase, "lowercase", "L", false, "include lowercase chars")
	rootCmd.Flags().BoolVarP(&opts.Numbers, "numbers", "n", false, "include numeric chars")
	rootCmd.Flags().BoolVarP(&opts.Symbols, "symbols", "s", false, "include symbol chars")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func main() {
	cli()
}
