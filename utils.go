package main

import (
	"fmt"
	"time"
	"os"
	"math"
	"errors"
	"strings"
	"unicode"
	"strconv"
	"encoding/base64"
	"crypto/rand"
	"net/smtp"
)

type Version struct
{
	Major int
	Minor int
	Patch int
	Tag string
	TagID int
}

func (version *Version) String() (out string) {
	out = strconv.Itoa(version.Major) + "." + strconv.Itoa(version.Minor) + "." + strconv.Itoa(version.Patch)
	if version.Tag != "" {
		out += "-" + version.Tag
		if version.TagID != 0 {
			out += strconv.Itoa(version.TagID)
		}
	}
	return
}

// Generate a cryptographically secure set of random bytes..
func GenerateSafeString(length int) (string, error) {
	rb := make([]byte,length)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(rb), nil
}

func relative_time(in string) (string, error) {
	if in == "" {
		return "", nil
	}
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, in)
	//t, err := time.ParseInLocation(layout, in, timeLocation)
	if err != nil {
		return "", err
	}

	diff := time.Since(t)
	hours := diff.Hours()
	seconds := diff.Seconds()
	weeks := int(hours / 24 / 7)
	months := int(hours / 24 / 31)
	switch {
		case months > 11:
			//return t.Format("Mon Jan 2 2006"), err
			return t.Format("Jan 2 2006"), err
		case months > 1: return fmt.Sprintf("%d months ago", months), err
		case months == 1: return "a month ago", err
		case weeks > 1: return fmt.Sprintf("%d weeks ago", weeks), err
		case int(hours / 24) == 7: return "a week ago", err
		case int(hours / 24) == 1: return "1 day ago", err
		case int(hours / 24) > 1: return fmt.Sprintf("%d days ago", int(hours / 24)), err
		case seconds <= 1: return "a moment ago", err
		case seconds < 60: return fmt.Sprintf("%d seconds ago", int(seconds)), err
		case seconds < 120: return "a minute ago", err
		case seconds < 3600: return fmt.Sprintf("%d minutes ago", int(seconds / 60)), err
		case seconds < 7200: return "an hour ago", err
		default: return fmt.Sprintf("%d hours ago", int(seconds / 60 / 60)), err
	}
}

func convert_byte_unit(bytes float64) (float64,string) {
	switch {
		case bytes >= float64(terabyte): return bytes / float64(terabyte), "TB"
		case bytes >= float64(gigabyte): return bytes / float64(gigabyte), "GB"
		case bytes >= float64(megabyte): return bytes / float64(megabyte), "MB"
		case bytes >= float64(kilobyte): return bytes / float64(kilobyte), "KB"
		default: return bytes, " bytes"
	}
}

func convert_byte_in_unit(bytes float64,unit string) (count float64) {
	switch(unit) {
		case "TB": count = bytes / float64(terabyte)
		case "GB": count = bytes / float64(gigabyte)
		case "MB": count = bytes / float64(megabyte)
		case "KB": count = bytes / float64(kilobyte)
		default: count = 0.1
	}

	if count < 0.1 {
		count = 0.1
	}
	return
}

func convert_unit(num int) (int,string) {
	switch {
		case num >= 1000000000000: return 0, "∞"
		case num >= 1000000000: return num / 1000000000, "B"
		case num >= 1000000: return num / 1000000, "M"
		case num >= 1000: return num / 1000, "K"
		default: return num, ""
	}
}

func convert_friendly_unit(num int) (int,string) {
	switch {
		case num >= 1000000000000: return 0, " zillion"
		case num >= 1000000000: return num / 1000000000, " billion"
		case num >= 1000000: return num / 1000000, " million"
		case num >= 1000: return num / 1000, " thousand"
		default: return num, ""
	}
}

func name_to_slug(name string) (slug string) {
	name = strings.TrimSpace(name)
	name = strings.Replace(name,"  "," ",-1)

	for _, char := range name {
		if unicode.IsLower(char) || unicode.IsNumber(char) {
			slug += string(char)
		} else if unicode.IsUpper(char) {
			slug += string(unicode.ToLower(char))
		} else if unicode.IsSpace(char) {
			slug += "-"
		}
	}

	if slug == "" {
		slug = "untitled"
	}
	return slug
}

func SendEmail(email string, subject string, msg string) (res bool) {
	// This hook is useful for plugin_sendmail or for testing tools. Possibly to hook it into some sort of mail server?
	if vhooks["email_send_intercept"] != nil {
		return vhooks["email_send_intercept"](email, subject, msg).(bool)
	}
	body := "Subject: " + subject + "\n\n" + msg + "\n"

	con, err := smtp.Dial(config.SmtpServer + ":" + config.SmtpPort)
	if err != nil {
		return
	}

	if config.SmtpUsername != "" {
		auth := smtp.PlainAuth("",config.SmtpUsername,config.SmtpPassword,config.SmtpServer)
		err = con.Auth(auth)
		if err != nil {
			return
		}
	}

	err = con.Mail(site.Email)
	if err != nil {
		return
	}
	err = con.Rcpt(email)
	if err != nil {
		return
	}

	email_data, err := con.Data()
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(email_data, body)
	if err != nil {
		return
	}

	err = email_data.Close()
	if err != nil {
		return
	}
	err = con.Quit()
	if err != nil {
		return
	}
	return true
}

func weak_password(password string) error {
	if len(password) < 8 {
		return errors.New("Your password needs to be at-least eight characters long.")
	}
	var charMap map[rune]int = make(map[rune]int)
	var numbers, /*letters, */symbols, upper, lower int
	for _, char := range password {
		charItem, ok := charMap[char]
		if ok {
			charItem++
		} else {
			charItem = 1
		}
		charMap[char] = charItem

		if unicode.IsLetter(char) {
			//letters++
			if unicode.IsUpper(char) {
				upper++
			} else {
				lower++
			}
		} else if unicode.IsNumber(char) {
			numbers++
		} else {
			symbols++
		}
	}

	if numbers == 0 {
		return errors.New("You don't have any numbers in your password.")
	}
	/*if letters == 0 {
		return errors.New("You don't have any letters in your password.")
	}*/
	if upper == 0 {
		return errors.New("You don't have any uppercase characters in your password.")
	}
	if lower == 0 {
		return errors.New("You don't have any lowercase characters in your password.")
	}
	if (len(password) / 2) > len(charMap) {
		return errors.New("You don't have enough unique characters in your password.")
	}

	if strings.Contains(strings.ToLower(password),"test") || /*strings.Contains(strings.ToLower(password),"123456") || */strings.Contains(strings.ToLower(password),"123") || strings.Contains(strings.ToLower(password),"password") || strings.Contains(strings.ToLower(password),"qwerty") {
		return errors.New("You may not have 'test', '123', 'password' or 'qwerty' in your password.")
	}
	return nil
}

func create_file(name string) error {
	f, err := os.Create(name)
	f.Close()
	return err
}

func write_file(name string, content string) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	f.Sync()
	f.Close()
	return
}

func word_count(input string) (count int) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0
	}
	in_space := false
	for _, value := range input {
		if unicode.IsSpace(value) {
			if !in_space {
				in_space = true
			}
		} else if in_space {
			count++
			in_space = false
		}
	}
	return count + 1
}

func getLevel(score int) (level int) {
	var base float64 = 25
	var current, prev float64
	exp_factor := 2.8

	for i := 1;;i++ {
		_, bit := math.Modf(float64(i) / 10)
		if bit == 0 {
			exp_factor += 0.1
		}
		current = base + math.Pow(float64(i), exp_factor) + (prev / 3)
		prev = current
		if float64(score) < current {
			break
		}
		level++
	}
	return level
}

func getLevelScore(getLevel int) (score int) {
	var base float64 = 25
	var current, prev float64
	var level int
	exp_factor := 2.8

	for i := 1;;i++ {
		_, bit := math.Modf(float64(i) / 10)
		if bit == 0 {
			exp_factor += 0.1
		}
		current = base + math.Pow(float64(i), exp_factor) + (prev / 3)
		prev = current
		level++
		if level <= getLevel {
			break
		}
	}
	return int(math.Ceil(current))
}

func getLevels(maxLevel int) []float64 {
	var base float64 = 25
	var current, prev float64 // = 0
	exp_factor := 2.8
	var out []float64
	out = append(out, 0)

	for i := 1;i <= maxLevel;i++ {
		_, bit := math.Modf(float64(i) / 10)
		if bit == 0 {
			exp_factor += 0.1
		}
		current = base + math.Pow(float64(i), exp_factor) + (prev / 3)
		prev = current
		out = append(out, current)
	}
	return out
}

func fill_group_id_gap(biggerID int, smallerID int) {
	dummy := Group{ID:0, Name:""}
	for i := smallerID; i > biggerID; i++ {
		groups = append(groups, dummy)
	}
}

func build_slug(slug string, id int) string {
	if slug == "" {
		return strconv.Itoa(id)
	}
	return slug + "." + strconv.Itoa(id)
}

func addModLog(action string, elementID int, elementType string, ipaddress string,  actorID int) (err error) {
	_, err = add_modlog_entry_stmt.Exec(action,elementID,elementType,ipaddress,actorID)
	if err != nil {
		return err
	}
	return nil
}

func addAdminLog(action string, elementID string, elementType int, ipaddress string, actorID int) (err error) {
	_, err = add_adminlog_entry_stmt.Exec(action,elementID,elementType,ipaddress,actorID)
	if err != nil {
		return err
	}
	return nil
}
