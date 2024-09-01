package pkg2

import (
	"time"

	"github.com/cscoding21/csmap/tests/common"
)

type Project struct {
	ID                *string               `json:"id,omitempty"`
	ProjectBasics     *ProjectBasics        `json:"projectBasics"`
	ProjectValue      *ProjectValue         `json:"projectValue"`
	ProjectCost       *ProjectCost          `json:"projectCost"`
	ProjectDaci       *ProjectDaci          `json:"projectDaci"`
	ControlFields     *common.ControlFields `json:"controlFields"`
	ProjectFeatures   []*ProjectFeature     `json:"projectFeatures,omitempty"`
	ProjectMilestones []*ProjectMilestone   `json:"projectMilestones,omitempty"`
}

type ProjectBasics struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	OwnerEmail  *string    `json:"ownerEmail,omitempty"`
}

type ProjectCost struct {
	Ongoing      *float64 `json:"ongoing,omitempty"`
	BlendedRate  *float64 `json:"blendedRate,omitempty"`
	InitialCost  *float64 `json:"initialCost,omitempty"`
	HourEstimate int      `json:"hourEstimate"`
}

type ProjectDaci struct {
	Driver      []*string `json:"driver,omitempty"`
	Approver    []*string `json:"approver,omitempty"`
	Contributor []*string `json:"contributor,omitempty"`
	Informed    []*string `json:"informed,omitempty"`
}

type ProjectFeature struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

type ProjectFilters struct {
	Status *string `json:"status,omitempty"`
}

type ProjectMilestone struct {
	ID        string                  `json:"id"`
	StartDate *time.Time              `json:"startDate,omitempty"`
	EndDate   *time.Time              `json:"endDate,omitempty"`
	Phase     *ProjectMilestonePhase  `json:"phase"`
	Tasks     []*ProjectMilestoneTask `json:"tasks"`
}

type ProjectMilestonePhase struct {
	ID          string `json:"id"`
	Order       int    `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectMilestoneTask struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	HourEstimate     int        `json:"hourEstimate"`
	Description      *string    `json:"description,omitempty"`
	RequiredSkillIDs []string   `json:"requiredSkillIDs,omitempty"`
	ResourceIDs      []string   `json:"resourceIDs,omitempty"`
	Status           string     `json:"status"`
	StartDate        *time.Time `json:"startDate,omitempty"`
	EndDate          *time.Time `json:"endDate,omitempty"`
}

type ProjectResults struct {
	Results []*Project `json:"results,omitempty"`
}

type ProjectValue struct {
	FundingSource        *string  `json:"fundingSource,omitempty"`
	DiscountRate         *float64 `json:"discountRate,omitempty"`
	YearOneValue         *float64 `json:"yearOneValue,omitempty"`
	YearTwoValue         *float64 `json:"yearTwoValue,omitempty"`
	YearThreeValue       *float64 `json:"yearThreeValue,omitempty"`
	YearFourValue        *float64 `json:"yearFourValue,omitempty"`
	YearFiveValue        *float64 `json:"yearFiveValue,omitempty"`
	NetPresentValue      *float64 `json:"netPresentValue,omitempty"`
	InternalRateOfReturn *float64 `json:"internalRateOfReturn,omitempty"`
}

type Projecttemplate struct {
	ID     string                  `json:"id"`
	Name   string                  `json:"name"`
	Phases []*ProjecttemplatePhase `json:"phases"`
}

type ProjecttemplatePhase struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	Description string `json:"description"`
}
