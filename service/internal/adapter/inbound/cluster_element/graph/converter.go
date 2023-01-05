package cluster_element_graph

import (
	"strconv"

	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/cluster_element"
)

func BatchMarshal(ces []cluster_element.ClusterElement) []*model.ClusterElement {
	ceMap := make(map[cluster_element.Identifier]*model.ClusterElement)
	result := make([]*model.ClusterElement, len(ces))
	for i, ce := range ces {
		m := &model.ClusterElement{
			ID:              strconv.Itoa(int(ce.Identifier())),
			DbscanProfileID: ce.DbscanProfileIdentifier(),
			SpotsProfileID:  ce.SpotProfileIdentifier(),
			SpotID:          ce.SpotIdentifier(),
			AssignedNumber:  ce.AssignedNumber(),
			Paths:           make([]*model.ClusterElement, 0, len(ce.Paths())),
		}
		result[i] = m
		ceMap[ce.Identifier()] = m
	}
	for i, ce := range ces {
		for _, cei := range ce.Paths() {
			result[i].Paths = append(result[i].Paths, ceMap[cei])
		}
	}
	return result
}
