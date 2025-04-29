package model

import "time"

type Product struct {
    Id               int
    ArrivalTime      time.Time
    EnteredCut       time.Time
    ExitedCut        time.Time
    EnteredAssemble  time.Time
    ExitedAssemble   time.Time
    EnteredPackage   time.Time
    ExitedPackage    time.Time
    RemainingTime    time.Duration 
}