package models

import (
	"sync"
)


var (
	PointsMap = make(map[string]int)
	MapMutex sync.Mutex
)