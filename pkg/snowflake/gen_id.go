package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(machineID uint16) (err error) {
	sonyMachineID = machineID
	t, _ := time.Parse("2006-01-02", "2020-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sonyFlake not initialized")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
