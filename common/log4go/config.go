// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	ErrAttr  = "LoadConfiguration: Error: Required attribute %s for filter missing in %s\n"
	ErrFile  = "LoadConfiguration: Error: Could not parse XML configuration in %q: %s\n"
	ErrChild = "LoadConfiguration: Error: Required child <%s> for filter missing in %s\n"
)

type xmlProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type xmlFilter struct {
	Enabled  bool          `xml:"enabled,attr"`
	Type     string        `xml:"type"`
	Level    string        `xml:"level"`
	Property []xmlProperty `xml:"property"`
}

type xmlLoggerConfig struct {
	Filter []xmlFilter `xml:"filter"`
}

func (this Log4go) LoadConfiguration(filename string) error {
	this.Close()
	// Open the configuration file
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	xc := new(xmlLoggerConfig)
	if err := xml.Unmarshal(contents, xc); err != nil {
		return fmt.Errorf(ErrFile, filename, err)
	}

	for _, xmlfilt := range xc.Filter {
		var (
			lvl  level
			call func(string, []xmlProperty) *LogWriter
		)

		// If we're disabled (syntax and correctness checks only), don't add to logger
		if !xmlfilt.Enabled {
			continue
		}

		switch xmlfilt.Level {
		case "DEBUG":
			lvl = DEBUG
		case "TRACE":
			lvl = TRACE
		case "INFO":
			lvl = INFO
		case "WARNING":
			lvl = WARNING
		case "ERROR":
			lvl = ERROR
		default:
			return fmt.Errorf(ErrChild, "level", filename)
		}

		switch xmlfilt.Type {
		case "console":
			call = xmlToConsoleLogWriter
		case "file":
			call = xmlToFileLogWriter
		default:
			return fmt.Errorf(ErrChild, "type", filename)
		}

		this.filters[lvl] = append(this.filters[lvl], call(levelStrings[lvl], xmlfilt.Property))
	}
	return nil
}

// Parse a number with K/M/G suffixes based on thousands (1000) or 2^10 (1024)
// Parse a number with S/M/H suffixes based on thousands (60)
func strToNumSuffix(str string, mult int) int {
	num := 1
	if len(str) < 2 {
		num *= mult
		parsed, _ := strconv.Atoi(str)
		return parsed * num
	}

	switch str[len(str)-1] {
	case 'G', 'g', 'H', 'h':
		num *= mult
		fallthrough
	case 'M', 'm':
		num *= mult
		fallthrough
	case 'K', 'k', 'S', 's':
		num *= mult
		str = str[0 : len(str)-1]
	}
	parsed, _ := strconv.Atoi(str)
	return parsed * num
}

func xmlToFileLogWriter(level string, props []xmlProperty) *LogWriter {
	file := ""
	format := "[%D %T] [%L] (%S) %M"
	maxsize := 0
	delay := 0
	// Parse properties
	for _, prop := range props {
		switch prop.Name {
		case "filename":
			file = strings.Trim(prop.Value, " \r\n")
		case "format":
			format = strings.Trim(prop.Value, " \r\n")
		case "maxsize":
			maxsize = strToNumSuffix(strings.Trim(prop.Value, " \r\n"), 1024)
		case "delay":
			delay = strToNumSuffix(strings.Trim(prop.Value, " \r\n"), 60)
		}
	}
	// Check properties
	if len(file) == 0 {
		return nil
	}
	out := NewFile(file, maxsize, delay)
	in := NewBuffer(format, level)
	writer := NewWriter(in, out)
	return &LogWriter{in, out, writer, log.New(in, "", 0)}
}

func xmlToConsoleLogWriter(level string, props []xmlProperty) *LogWriter {
	format := "[%D %T] [%L] (%S) %M"
	// Parse properties
	for _, prop := range props {
		switch prop.Name {
		case "format":
			format = strings.Trim(prop.Value, " \r\n")
		}
	}
	out := NewConsole()
	in := NewBuffer(format, level)
	writer := NewWriter(in, out)
	return &LogWriter{in, out, writer, log.New(in, "", 0)}
}
