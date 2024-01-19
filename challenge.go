package raft

import (
	"encoding/json"
	"errors"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

const (
	ContractName   = "status_manage"
	ContractMethod = "challenge"
)

const (
	AppNameKey     = "app_name"
	AppVersionKey  = "app_version"
	MethodKey      = "method"
	ParamKey       = "param"
	PreTxStatusKey = "pre_tx_status"
	TxStatusKey    = "tx_status"
)

type ChallengeResp struct {
	AppName            string `json:"app_name"`
	IsChallengeSuccess bool   `json:"is_challenge_success"`
}

/*
appName := string(params["app_name"])
	appVersion := string(params["app_version"])
	method := string(params["method"])
	req := params["param"]
	preTxStatus := params["pre_tx_status"]
	txStatus := params["tx_status"]
	if appName == "" || appVersion == "" || method == "" || preTxStatus == nil {
		return sdk.Error("invalid params")
	}
*/

func Challenge(args []*common.KeyValuePair, configPath string) (isSuccess bool, err error) {
	// todo 证书等逻辑
	client, err := sdk.NewChainClient(sdk.WithConfPath(configPath))
	if err != nil {
		return
	}
	defer client.Stop()
	resp, err := client.InvokeContract(ContractName, ContractMethod, "", args, -1, true)
	if err != nil {
		return
	}
	if resp.Code != 0 {
		return false, errors.New(resp.Message)
	}
	result := &ChallengeResp{}
	err = json.Unmarshal(resp.ContractResult.Result, result)
	if err != nil {
		return
	}
	return result.IsChallengeSuccess, nil
}
