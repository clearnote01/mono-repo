package parser

import (
	"github.com/gc-cli/api"
	"github.com/gc-cli/models"
	"github.com/repeale/fp-go"
)

// library which takes api data and enriches it

func GetErrorGroupsWithResolutionStatus(status string) *models.GroupStats {
	res := api.GetErrorGroups(0, 5)

	filterWithResolutionStatus := fp.Filter(func(x models.GroupStat) bool { return x.Group.ResolutionStatus == status })

	groupsWithResolutionStatus := filterWithResolutionStatus(res.ErrorGroupStats)
	res.ErrorGroupStats = groupsWithResolutionStatus
	return res
}

func GetErrorGroupsOpenStatus() *models.GroupStats {
	return GetErrorGroupsWithResolutionStatus("OPEN")
}
