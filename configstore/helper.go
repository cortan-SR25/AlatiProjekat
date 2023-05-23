package configstore

import (
	"fmt"
)

/*consul baza funkcionise tako sto se value (objekat, u nasem slucaju config i group) smesta u putanju (key) pomocu koje mozemo
filtrirati pretragu entiteta i vrsiti upite
*/
const ( //konsante za generisanje tih key-eva (putanja) po string sablonu u metodama generateConfigKey i generateGroupKey
	allConfigs    = "config"
	allGroups     = "cfgroups"
	config        = "config/%s"
	configVersion = "config/%s/%s"
	group         = "cfgroups/%s"
	groupVersion  = "cfgroups/%s/%s"

	/* Posto pretraga u bazi funkcionise po prefiks prinicipu (npr: vrati mi objekte pod kljucem "cfgroups/{id}/",
	baza vraca grupu konfiguracija sa tim ID-jem ali i pojedinacno sve konfiguracije koje se nalaze u njoj, po njenoj daljoj putanji
	"cfgroups/{id}/......{version}/config/{labels}/{configId}")
	To ne zelimo da se desi jer se konfiguracije vec vide u grupama a i JSON svkako ne moze da Unmarshal-uje dva objekta
	razlicitog tipa iz baze u jedan tip.
	Sa druge strane ne bi bilo pametno ni da konfiguracije iz grupa smestamo u putanje sa preiksom "config" jer su tu
	nezavisne konfiguracije koje nisu u grupama pa to ne bi imalo smisla.
	Najpametnije bi bilo staviti ih u treci prefiks putanja npr ovako umesto "cfgroups" i "config" da bude "groupConfig"*/
	groupVersionConfigLabelId = "groupConfig/%s/%s/%s/%s/%s"
	groupVersionConfigLabel   = "groupConfig/%s/%s/%s/%s"
)

//metoda za generisanje key putanje za config koji smestamo u bazu / vadimo iz baze
func generateConfigKey(id string, version string) (string, string) {
	if version != "" {
		return fmt.Sprintf(configVersion, id, version), id
	} else {
		return fmt.Sprintf(config, id), id
	}

}

func constructConfigKey(id string, version string) string {
	if version != "" {
		return fmt.Sprintf(configVersion, id, version)
	} else {
		return fmt.Sprintf(config, id)
	}

}

func generateGroupKey(groupId string, version string, configId string, configLabels string) string {
	if configLabels != "" && configId != "" {
		return fmt.Sprintf(groupVersionConfigLabelId, groupId, version, "config", configLabels, configId)
	} else if configLabels != "" && configId == "" {
		return fmt.Sprintf(groupVersionConfigLabel, groupId, version, "config", configLabels)
	} else if configLabels == "" && configId == "" && version != "" {
		return fmt.Sprintf(groupVersion, groupId, version)
	} else {
		return fmt.Sprintf(group, groupId)
	}
}

func constructGroupKey(id string, version string, configId string, configLabels string) string {
	if configLabels != "" && configId != "" {
		return fmt.Sprintf(groupVersionConfigLabelId, id, version, "config", configLabels, configId)
	} else if configLabels != "" && configId == "" {
		return fmt.Sprintf(groupVersionConfigLabel, id, version, "config", configLabels)
	} else if configLabels == "" && configId == "" && version != "" {
		return fmt.Sprintf(groupVersion, id, version)
	} else {
		return fmt.Sprintf(group, id)
	}
}
