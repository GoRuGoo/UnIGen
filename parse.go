package main

import (
	"errors"
	"flag"
	"strings"
	"time"
)

type Options struct {
	base   string
	format string
}

func paseOptions() Options {
	base := flag.String("base", "", "Base strgin with placeholder {timestamp}")
	format := flag.String("format", "2005-07-30-12-00-00", "Timestamp format")
	flag.Parse()

	options := Options{
		base:   *base,
		format: *format,
	}

	return options
}

func (o Options) validateOptions() error {
	// Validation base
	if o.base == "" {
		return errors.New("Base string is required")
	}

	// Check that the {timestamp} placeholder is included as it is needed to generate test data
	if !strings.Contains(o.base, "{timestamp}") {
		return errors.New("Base string must contain {timestamp}")
	}

	// Dynamically parse to ensure formatting is correct
	_, err := time.Parse(o.format, "2006-01-02-15-04-05")
	if err != nil {
		return errors.New("Invalid timestamp format")
	}

	return nil
}
