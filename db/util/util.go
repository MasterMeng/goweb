/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	dbURLRegex = regexp.MustCompile(`(Datasource:\s*)?(\S+):(\S+)@|(Datasource:.*\s)?(user=\S+).*\s(password=\S+)|(Datasource:.*\s)?(password=\S+).*\s(user=\S+)`)
)

// GetDBName gets database name from connection string
func GetDBName(datasource string) string {
	var dbName string
	datasource = strings.ToLower(datasource)

	re := regexp.MustCompile(`(?:\/([^\/?]+))|(?:dbname=([^\s]+))`)
	getName := re.FindStringSubmatch(datasource)
	if getName != nil {
		dbName = getName[1]
		if dbName == "" {
			dbName = getName[2]
		}
	}

	return dbName
}

// MaskDBCred hides DB credentials in connection string
func MaskDBCred(str string) string {
	matches := dbURLRegex.FindStringSubmatch(str)

	// If there is a match, there should be three entries: 1 for
	// the match and 9 for submatches (see dbURLRegex regular expression)
	if len(matches) == 10 {
		matchIdxs := dbURLRegex.FindStringSubmatchIndex(str)
		substr := str[matchIdxs[0]:matchIdxs[1]]
		for idx := 1; idx < len(matches); idx++ {
			if matches[idx] != "" {
				if strings.Index(matches[idx], "user=") == 0 {
					substr = strings.Replace(substr, matches[idx], "user=****", 1)
				} else if strings.Index(matches[idx], "password=") == 0 {
					substr = strings.Replace(substr, matches[idx], "password=****", 1)
				} else {
					substr = strings.Replace(substr, matches[idx], "****", 1)
				}
			}
		}
		str = str[:matchIdxs[0]] + substr + str[matchIdxs[1]:]
	}
	return str
}

// GetBlogDataSource returns a datasource with a unqiue database name
func GetBlogDataSource(dbtype, datasource string) string {
	if dbtype == "sqlite3" {
		ext := filepath.Ext(datasource)
		dbName := strings.TrimSuffix(filepath.Base(datasource), ext)
		datasource = fmt.Sprintf("%s_%s", dbName, ext)
	} else {
		dbName := getDBName(datasource)
		datasource = strings.Replace(datasource, dbName, fmt.Sprintf("%s", dbName), 1)
	}
	return datasource
}

// getDBName gets database name from connection string
func getDBName(datasource string) string {
	var dbName string
	datasource = strings.ToLower(datasource)

	re := regexp.MustCompile(`(?:\/([^\/?]+))|(?:dbname=([^\s]+))`)
	getName := re.FindStringSubmatch(datasource)
	if getName != nil {
		dbName = getName[1]
		if dbName == "" {
			dbName = getName[2]
		}
	}

	return dbName
}

// GetError wraps error passed in with context
func GetError(err error, getType string) error {
	if err.Error() == "sql: no rows in result set" {
		return fmt.Errorf("Failed to get %s: %s", getType, err)
	}
	return fmt.Errorf("Failed to process database request: %s", err)
}
