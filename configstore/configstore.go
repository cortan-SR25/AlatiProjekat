package configstore

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

type ConfigStore struct {
	cli *api.Client
}

func New() (*ConfigStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConfigStore{
		cli: client,
	}, nil
}

func (cs *ConfigStore) GetAllConfigs() ([]*Config, error) { //done
	kv := cs.cli.KV()
	data, _, err := kv.List(allConfigs, nil)
	if err != nil {
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func (cs *ConfigStore) PostConfig(config *Config) (*Config, error) { //done
	kv := cs.cli.KV()

	config.Id = uuid.NewString()
	sid, _ := generateConfigKey(config.Id, config.Version)

	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (cs *ConfigStore) GetConfigById(id string) ([]*Config, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructConfigKey(id, ""), nil)
	if err != nil {
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func (cs *ConfigStore) GetConfigByIdAndVersion(id string, version string) ([]*Config, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructConfigKey(id, version), nil)
	if err != nil {
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func (cs *ConfigStore) DeleteConfigById(id string) (map[string]string, error) { //done
	kv := cs.cli.KV()
	_, err := kv.DeleteTree(constructConfigKey(id, ""), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (cs *ConfigStore) DeleteConfigByIdAndVersion(id string, version string) (map[string]string, error) { //done
	kv := cs.cli.KV()
	_, err := kv.DeleteTree(constructConfigKey(id, version), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (cs *ConfigStore) GetAllGroups() ([]*CfGroup, error) { //done
	kv := cs.cli.KV()
	data, _, err := kv.List(allGroups, nil)
	if err != nil {
		return nil, err
	}

	cfgroups := []*CfGroup{}
	for _, pair := range data {

		cfgroup := &CfGroup{}
		err = json.Unmarshal(pair.Value, cfgroup)
		if err != nil {
			return nil, err
		}
		cfgroups = append(cfgroups, cfgroup)
	}

	return cfgroups, nil
}

func (cs *ConfigStore) PostCfGroup(cfgroup *CfGroup) (*CfGroup, error) { //done
	kv := cs.cli.KV()

	groupId := uuid.New().String()
	cfgroup.Id = groupId

	for i := 0; i < len(cfgroup.Configurations); i++ {
		configId := uuid.New().String()
		cfgroup.Configurations[i].Id = configId
		sid := generateGroupKey(groupId, cfgroup.Version, cfgroup.Configurations[i].Id, cfgroup.Configurations[i].Labels)
		data, err := json.Marshal(cfgroup.Configurations[i])
		if err != nil {
			return nil, err
		}

		p := &api.KVPair{Key: sid, Value: data}
		_, err = kv.Put(p, nil)
		if err != nil {
			return nil, err
		}

	}
	bid := generateGroupKey(cfgroup.Id, cfgroup.Version, "", "")

	data, err := json.Marshal(cfgroup)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: bid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return cfgroup, nil
}

func (cs *ConfigStore) GetCfGroupById(id string) ([]*CfGroup, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructGroupKey(id, "", "", ""), nil)
	if err != nil {
		return nil, err
	}

	cfgroups := []*CfGroup{}
	for _, pair := range data {
		cfgroup := &CfGroup{}
		fmt.Print(pair.Value)
		err = json.Unmarshal(pair.Value, cfgroup)
		if err != nil {
			return nil, err
		}
		cfgroups = append(cfgroups, cfgroup)
	}
	return cfgroups, nil
}

func (cs *ConfigStore) GetCfGroupByIdAndVersion(id string, version string) ([]*CfGroup, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructGroupKey(id, version, "", ""), nil)
	if err != nil {
		return nil, err
	}

	cfgroups := []*CfGroup{}
	for _, pair := range data {
		cfgroup := &CfGroup{}
		err = json.Unmarshal(pair.Value, cfgroup)
		if err != nil {
			return nil, err
		}
		cfgroups = append(cfgroups, cfgroup)
	}
	return cfgroups, nil
}

func (cs *ConfigStore) DeleteCfGroupById(id string) (map[string]string, error) { //done
	kv := cs.cli.KV()
	_, err := kv.DeleteTree(constructGroupKey(id, "", "", ""), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (cs *ConfigStore) DeleteCfGroupByIdAndVersion(id string, version string) (map[string]string, error) { //done
	kv := cs.cli.KV()
	_, err := kv.DeleteTree(constructGroupKey(id, "", "", ""), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (cs *ConfigStore) GetGroupConfigByIdAndLabel(groupId string, version string, configId string, labels string) ([]*Config, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructGroupKey(groupId, version, configId, labels), nil)
	if err != nil {
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func (cs *ConfigStore) GetGroupConfigByLabel(groupId string, version string, labels string) ([]*Config, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructGroupKey(groupId, version, "", labels), nil)
	if err != nil {
		return nil, err
	}

	configs := []*Config{}
	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

//metoda za brisanje jedne konfiguracije iz grupe
func (cs *ConfigStore) DeleteGroupConfigByLabelAndId(groupId string, version string, labels string, configId string) (map[string]string, error) { //done
	kv := cs.cli.KV()

	_, errDeleteConfig := kv.DeleteTree(constructGroupKey(groupId, version, configId, labels), nil)
	if errDeleteConfig != nil {
		return nil, errDeleteConfig
	}

	var oldCfGroup = &CfGroup{}
	var newCfGroup = &CfGroup{}

	data, _, err := kv.List(constructGroupKey(groupId, version, "", ""), nil)
	if err != nil {
		return nil, err
	}

	for _, pair := range data {
		err = json.Unmarshal(pair.Value, oldCfGroup)
		if err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(oldCfGroup.Configurations); i++ {
		if oldCfGroup.Configurations[i].Id != configId {
			newCfGroup.Configurations = append(newCfGroup.Configurations, oldCfGroup.Configurations[i])
		}
	}

	newCfGroup.Id = oldCfGroup.Id
	newCfGroup.Version = oldCfGroup.Version

	_, errDeleteGroup := kv.DeleteTree(constructGroupKey(groupId, version, "", ""), nil)
	if errDeleteGroup != nil {
		return nil, errDeleteGroup
	}

	sid := generateGroupKey(newCfGroup.Id, newCfGroup.Version, "", "")

	dataCfGroup, err := json.Marshal(newCfGroup)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: dataCfGroup}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": configId}, nil
}

func (cs *ConfigStore) PutGroupConfigByGroupId(config *Config, groupId string) (*CfGroup, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructGroupKey(groupId, "", "", ""), nil)
	if err != nil {
		return nil, err
	}

	var newCfGroup *CfGroup
	for _, pair := range data {
		cfgroup := &CfGroup{}
		err = json.Unmarshal(pair.Value, cfgroup)
		if err != nil {
			return nil, err
		}
		_, err := kv.DeleteTree(constructGroupKey(groupId, "", "", ""), nil)
		if err != nil {
			return nil, err
		}
		config.Id = uuid.NewString()

		newCfGroup = cfgroup
		newCfGroup.Configurations = cfgroup.Configurations
		newCfGroup.Configurations = append(newCfGroup.Configurations, config)

		bid := generateGroupKey(newCfGroup.Id, newCfGroup.Version, "", "")

		data, err := json.Marshal(newCfGroup)
		if err != nil {
			return nil, err
		}

		p := &api.KVPair{Key: bid, Value: data}
		_, err = kv.Put(p, nil)
		if err != nil {
			return nil, err
		}

		sid := generateGroupKey(groupId, newCfGroup.Version, config.Id, config.Labels)
		configData, configErr := json.Marshal(config)
		if configErr != nil {
			return nil, err
		}

		b := &api.KVPair{Key: sid, Value: configData}
		_, err = kv.Put(b, nil)
		if err != nil {
			return nil, err
		}

	}
	return newCfGroup, nil
}

func (cs *ConfigStore) PutConfig(newConfig *Config, configId string) (*Config, error) { //done
	kv := cs.cli.KV()

	data, _, err := kv.List(constructConfigKey(configId, ""), nil)
	if err != nil {
		return nil, err
	}

	for _, pair := range data {
		config := &Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		newConfig.Id = config.Id

		_, err := kv.DeleteTree(constructConfigKey(config.Id, ""), nil)
		if err != nil {
			return nil, err
		}

		sid, _ := generateConfigKey(newConfig.Id, newConfig.Version)

		data, err := json.Marshal(newConfig)
		if err != nil {
			return nil, err
		}

		p := &api.KVPair{Key: sid, Value: data}
		_, err = kv.Put(p, nil)
		if err != nil {
			return nil, err
		}
	}

	return newConfig, nil
}
