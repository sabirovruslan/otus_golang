package hw10programoptimization

import (
	"bufio"
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

func getUserEmails(r io.Reader) (result userEmails, err error) {
	var user UserEmail
	var i int
	scaner := bufio.NewReader(r)
	for {
		line, _, err := scaner.ReadLine()
		if err != nil {
			break
		}
		if err = json.Unmarshal([]byte(line), &user); err != nil {
			continue
		}
		result[i] = user
		i++
	}
	return
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
