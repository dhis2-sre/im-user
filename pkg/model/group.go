package model

import (
	"go.mozilla.org/sops/v3/cmd/sops/formats"
	"go.mozilla.org/sops/v3/decrypt"
	"gorm.io/gorm"
	"log"
)

const AdministratorGroupName = "administrators"

type Group struct {
	gorm.Model
	Name                 string               `gorm:"unique;"`
	Hostname             string               `gorm:"unique;"`
	Users                []User               `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClusterConfiguration ClusterConfiguration `json:"clusterConfiguration"`
}

type ClusterConfiguration struct {
	gorm.Model
	GroupID                 uint
	KubernetesConfiguration []byte
}

func (c ClusterConfiguration) GetKubernetesConfigurationInCleartext() ([]byte, error) {
	return c.decrypt(c.KubernetesConfiguration, "yaml")
}

func (c ClusterConfiguration) decrypt(data []byte, format string) ([]byte, error) {
	kubernetesConfigurationCleartext, err := decrypt.DataWithFormat(data, formats.FormatFromString(format))
	if err != nil {
		log.Printf("Error decrypting: %s\n", err)
		return nil, err
	}
	return kubernetesConfigurationCleartext, nil
}
