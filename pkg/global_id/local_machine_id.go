package global_id

import "github.com/dgdts/CloudNoteServer/pkg/utils"

var _ machineIDGetter = (*localMachineIDGetter)(nil)

type localMachineIDGetter struct {
}

func (l *localMachineIDGetter) GetMachineID() (int, error) {
	ip, err := utils.GetLocalIP()
	if err != nil {
		return 0, err
	}

	return lastToInt(ip)
}
