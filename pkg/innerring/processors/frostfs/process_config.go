package frostfs

import (
	nmClient "github.com/TrueCloudLab/frostfs-node/pkg/morph/client/netmap"
	frostfsEvent "github.com/TrueCloudLab/frostfs-node/pkg/morph/event/frostfs"
	"go.uber.org/zap"
)

// Process config event by setting configuration value from the mainchain in
// the sidechain.
func (np *Processor) processConfig(config *frostfsEvent.Config) {
	if !np.alphabetState.IsAlphabet() {
		np.log.Info("non alphabet mode, ignore config")
		return
	}

	prm := nmClient.SetConfigPrm{}

	prm.SetID(config.ID())
	prm.SetKey(config.Key())
	prm.SetValue(config.Value())
	prm.SetHash(config.TxHash())

	err := np.netmapClient.SetConfig(prm)
	if err != nil {
		np.log.Error("can't relay set config event", zap.Error(err))
	}
}
