package cluster_element

import (
	"errors"

	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
)

// TODO: clusterという集約ルートを作成するべき

const (
	CLUSTER_NOT_ASSIGN = 0  // クラスタリング未処理
	CLUSTER_LACK       = -1 // MinCountに満たないクラスタ
)

var (
	errorUpdateAssignedNumber = errors.New("update assigned number error")
)

type ClusterElement interface {
	Identifier() Identifier
	DbscanProfileIdentifier() dbscan_profile.Identifier
	SpotIdentifier() spot.Identifier

	AssignedNumber() int
	Paths() []Identifier

	UpdatePaths([]Identifier) error
	UpdateAssignedNumber(int) error
	// 無効クラスタに所属させる
	LackAssign() error
	// 無所属にする
	Unassign() error

	OverWrite(ClusterElementPreference) error

	IsAssigned() bool
}

type ClusterElementPreference struct {
	AssignedNumber int
	Paths          []Identifier
}

func New(
	i Identifier,
	dsi dbscan_profile.Identifier,
	si spot.Identifier,
) ClusterElement {
	return &clusterElement{
		identifier:              i,
		dbscanProfileIdentifier: dsi,
		spotIdentifier:          si,
		assignedNumber:          CLUSTER_NOT_ASSIGN,
		paths:                   []Identifier{},
	}
}

func Restore(
	i Identifier,
	dsi dbscan_profile.Identifier,
	si spot.Identifier,
	p ClusterElementPreference,
) (ClusterElement, error) {
	ce := New(i, dsi, si)
	if err := ce.OverWrite(p); err != nil {
		return nil, err
	}
	return ce, nil
}

type clusterElement struct {
	identifier              Identifier
	dbscanProfileIdentifier dbscan_profile.Identifier
	spotIdentifier          spot.Identifier
	assignedNumber          int
	paths                   []Identifier
}

func (e *clusterElement) Identifier() Identifier {
	return e.identifier
}
func (e *clusterElement) DbscanProfileIdentifier() dbscan_profile.Identifier {
	return e.dbscanProfileIdentifier
}
func (e *clusterElement) SpotIdentifier() spot.Identifier {
	return e.spotIdentifier
}
func (e *clusterElement) AssignedNumber() int {
	return e.assignedNumber
}
func (e *clusterElement) Paths() []Identifier {
	return e.paths
}

func (e *clusterElement) UpdatePaths(paths []Identifier) error {
	e.paths = paths
	return nil
}

func (e *clusterElement) UpdateAssignedNumber(n int) error {
	if n < 1 {
		return errorUpdateAssignedNumber
	}
	e.assignedNumber = n
	return nil
}

func (e *clusterElement) LackAssign() error {
	e.assignedNumber = CLUSTER_LACK
	return nil
}

func (e *clusterElement) Unassign() error {
	e.assignedNumber = CLUSTER_NOT_ASSIGN
	return nil
}

func (e *clusterElement) OverWrite(p ClusterElementPreference) error {
	switch p.AssignedNumber {
	case CLUSTER_NOT_ASSIGN:
		if err := e.Unassign(); err != nil {
			return err
		}
	case CLUSTER_LACK:
		if err := e.LackAssign(); err != nil {
			return err
		}
	default:
		if err := e.UpdateAssignedNumber(p.AssignedNumber); err != nil {
			return err
		}
	}
	if err := e.UpdatePaths(p.Paths); err != nil {
		return err
	}
	return nil
}

func (e *clusterElement) IsAssigned() bool {
	return e.assignedNumber != CLUSTER_NOT_ASSIGN
}
