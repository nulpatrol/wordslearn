package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// makeTime is a helper to create a time duration
func makeTime(h int, m int, s int, ms int) time.Time {
	return time.Date(0, 1, 1, h, m, s, ms*1000*1000, time.UTC)
}

// parseSrtTime parses a srt subtitle time (duration since start of film)
func parseSrtTime(in string) (time.Time, error) {
	// . and , to :
	in = strings.Replace(in, ",", ":", -1)
	in = strings.Replace(in, ".", ":", -1)

	if strings.Count(in, ":") == 2 {
		in += ":000"
	}

	r1 := regexp.MustCompile("([0-9]+):([0-9]+):([0-9]+):([0-9]+)")
	matches := r1.FindStringSubmatch(in)
	if len(matches) < 5 {
		return time.Now(), fmt.Errorf("[srt] Regexp didnt match: %s", in)
	}
	h, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Now(), err
	}
	m, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Now(), err
	}
	s, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Now(), err
	}
	ms, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Now(), err
	}

	return makeTime(h, m, s, ms), nil
}

// parseVttTime parses a VTT time (duration since start of film)
// specification at https://www.w3.org/TR/webvtt1/#cue-timings-and-settings-parsing
func parseVttTime(in string) (time.Time, error) {
	// . and , to :
	in = strings.Replace(in, ",", ":", -1)
	in = strings.Replace(in, ".", ":", -1)

	if strings.Count(in, ":") == 2 {
		in = "00:" + in
	}

	r1 := regexp.MustCompile("([0-9]+):([0-9]+):([0-9]+):([0-9]+)")
	matches := r1.FindStringSubmatch(in)
	if len(matches) < 5 {
		return time.Now(), fmt.Errorf("[vtt] Regexp didnt match: %s", in)
	}
	h, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Now(), err
	}
	m, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Now(), err
	}
	s, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Now(), err
	}
	ms, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Now(), err
	}

	return makeTime(h, m, s, ms), nil
}
