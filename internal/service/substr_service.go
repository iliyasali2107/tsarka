package service

import (
	"context"
	"regexp"
)

type SubstrService interface {
	Find(ctx context.Context, substr string) string
	Validate(ctx context.Context, substr string) error
}

type substrService struct{}

func NewSubstrService() SubstrService {
	return &substrService{}
}

func (ss *substrService) Find(ctx context.Context, substr string) string {
	lastOccurrence := make(map[byte]int)
	longestSubstring := ""
	currentSubstring := ""
	start := 0

	for end := 0; end < len(substr); end++ {
		if index, ok := lastOccurrence[substr[end]]; ok && index >= start {
			start = index + 1
		}

		lastOccurrence[substr[end]] = end

		currentSubstring = substr[start : end+1]

		if len(currentSubstring) > len(longestSubstring) {
			longestSubstring = currentSubstring
		}
	}

	return longestSubstring
}

func (ss *substrService) Validate(ctx context.Context, substr string) error {
	regexPattern := "^[a-zA-Z0-9]+$"
	match, err := regexp.MatchString(regexPattern, substr)
	if err != nil {
		return err
	}

	if !match {
		return ErrMatch
	}

	return nil
}
