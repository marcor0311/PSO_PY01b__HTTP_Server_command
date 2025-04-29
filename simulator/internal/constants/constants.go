package constants

import "time"

const (
    StationCutting    = "Cutting"
    StationAssembling = "Assembling"
    StationPackaging  = "Packaging"
)

const (
	CuttingMinTime    = 2 * time.Second
	CuttingMaxTime    = 3 * time.Second
	AssemblingMinTime = 1 * time.Second
	AssemblingMaxTime = 2 * time.Second
	PackagingMinTime  = 0 * time.Second
	PackagingMaxTime  = 1 * time.Second
)