package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type UserEmail struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	ue, err := getUserEmails(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(ue, domain)
}

type userEmails [100_000]UserEmail

func getUserEmails(r io.Reader) (userEmails, error) {
	var user UserEmail
	var i int
	var result userEmails
	scaner := bufio.NewReader(r)
	for {
		line, _, err := scaner.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return result, nil
		}
		if err = json.Unmarshal([]byte(line), &user); err != nil {
			return result, err
		}
		result[i] = user
		i++
	}
	return result, nil
}

func countDomains(ue userEmails, domain string) (DomainStat, error) {
	result := make(DomainStat)
	var key string

	for _, user := range ue {
		matched := strings.Contains(user.Email, domain)

		if matched {
			key = strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			num := result[key]
			num++
			result[key] = num
		}
	}
	return result, nil
}
