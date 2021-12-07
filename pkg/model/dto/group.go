package dto

import "github.com/dhis2-sre/im-users/pkg/model"

type Group struct {
	Id                   uint                 `json:"id"`
	Name                 string               `json:"name"`
	Hostname             string               `json:"hostname"`
	Users                []User               `json:"users"`
	ClusterConfiguration ClusterConfiguration `json:"clusterConfiguration,omitempty"`
}

type ClusterConfiguration struct {
	Id                      uint   `json:"id"`
	GroupID                 uint   `json:"groupId"`
	KubernetesConfiguration []byte `json:"kubernetesConfiguration"`
}

func ToGroup(group *model.Group) Group {
	clusterConfiguration := ClusterConfiguration{
		Id:                      group.ClusterConfiguration.ID,
		GroupID:                 group.ID,
		KubernetesConfiguration: group.ClusterConfiguration.KubernetesConfiguration,
	}
	return Group{
		Id:                   group.ID,
		Name:                 group.Name,
		Hostname:             group.Hostname,
		Users:                toUsers(group.Users),
		ClusterConfiguration: clusterConfiguration,
	}
}

func toGroups(groups []model.Group) []Group {
	dtoGroups := make([]Group, len(groups))
	for i, group := range groups {
		dtoGroups[i] = ToGroup(&group)
	}
	return dtoGroups
}
