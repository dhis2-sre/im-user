package dto

import "github.com/dhis2-sre/im-users/pgk/model"

type Group struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Users    []User `json:"users"`
}

type ClusterConfiguration struct {
	Id                      uint   `json:"id"`
	GroupID                 uint   `json:"groupId"`
	KubernetesConfiguration []byte `json:"kubernetesConfiguration"`
}

func ToGroup(group *model.Group) Group {
	return Group{
		Id:       group.ID,
		Name:     group.Name,
		Hostname: group.Hostname,
		Users:    toUsers(group.Users),
	}
}

func toGroups(groups []model.Group) []Group {
	dtoGroups := make([]Group, len(groups))
	for i, group := range groups {
		dtoGroups[i] = ToGroup(&group)
	}
	return dtoGroups
}
