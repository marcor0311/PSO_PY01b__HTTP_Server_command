package constants

import "time"

const (
    StationCutting    = "Cutting"
    StationAssembling = "Assembling"
    StationPackaging  = "Packaging"
)

const (
	CuttingMinTime    = 1 * time.Second
	CuttingMaxTime    = 4 * time.Second
	AssemblingMinTime = 2 * time.Second
	AssemblingMaxTime = 3 * time.Second
	PackagingMinTime  = 0 * time.Second
	PackagingMaxTime  = 3 * time.Second
)

