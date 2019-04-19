package plugin

import (
	_ "github.com/racalwang/tschain/plugin/consensus/init" //consensus init
	_ "github.com/racalwang/tschain/plugin/crypto/init"    //crypto init
	_ "github.com/racalwang/tschain/plugin/dapp/init"      //dapp init
	_ "github.com/racalwang/tschain/plugin/mempool/init"   //mempool init
	_ "github.com/racalwang/tschain/plugin/store/init"     //store init
)
