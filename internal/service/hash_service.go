package service

import (
	"fmt"
	"hash/crc64"
	"math/bits"
	"strconv"
	"sync"
	"time"

	"tsarka/internal/config"
)

var timeMutex sync.Mutex

type HashService interface {
	Hash(str string) string
	SendRequestToRequests(req Request) bool
	GetRequestFromRequests() Request
	IsBusy() bool
	GetBufferLen() int
}

type hashService struct {
	MaxRequests  int
	Interval     time.Duration
	CalcDuration time.Duration
	HashResults  sync.Map
	Requests     chan Request
}

func NewHashService(cfg config.Config) HashService {
	return &hashService{
		MaxRequests:  cfg.HashMaxRequests,
		Interval:     cfg.HashInternval,
		CalcDuration: cfg.HashCalcDuration,
		HashResults:  sync.Map{},
		Requests:     make(chan Request, cfg.HashMaxRequests),
	}
}

func (hs *hashService) Hash(str string) string {
	crcHash := crc64.Checksum([]byte(str), crc64.MakeTable(crc64.ISO))
	startTime := time.Now()
	var ones int
	for time.Since(startTime) < hs.CalcDuration {
		currentTime := myTimeNow()
		hash := crcHash & uint64(currentTime.UnixNano())
		ones = bits.OnesCount64(hash)
		time.Sleep(hs.Interval)
		fmt.Println(ones)
	}
	return strconv.Itoa(ones)
}

func myTimeNow() time.Time {
	timeMutex.Lock()
	defer timeMutex.Unlock()

	return time.Now()
}

type Request struct {
	Id  string
	Str string
}

func (hs *hashService) SendRequestToRequests(req Request) bool {
	select {
	case hs.Requests <- req:
		return true
	default:
		return false
	}
}

func (hs *hashService) GetRequestFromRequests() Request {
	return <-hs.Requests
}

func (hs *hashService) IsBusy() bool {
	return len(hs.Requests) == hs.MaxRequests-1
}

func (hs *hashService) GetBufferLen() int {
	return len(hs.Requests)
}


// 760720 