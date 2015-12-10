package entity_effectors

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"net/url"
	"errors"
	"strings"
	"bytes"
	"strconv"
)

func EffectorList(network *net.Network, application, entity string) []models.EffectorSummary {
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors", application, entity)
	body, err := network.SendGetRequest(path)
	if err != nil {
		fmt.Println(err)
	}

	var effectorList []models.EffectorSummary
	err = json.Unmarshal(body, &effectorList)
	if err != nil {
		fmt.Println(err)
	}
	return effectorList
}

func TriggerEffector(network *net.Network, application, entity, effector string, params []string, args []string) (string, error) {
	if len(params) != len(args) {
		return "", errors.New(strings.Join([]string{"Parameters not supplied:", strings.Join(params, ", ")}, " "))
	}
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors/%s", application, entity, effector)
	data := url.Values{}
	for i := range params {
		data.Set(params[i], args[i])
	}
	req := network.NewPostRequest(path, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	body, err := network.SendRequest(req)
	if err != nil {
		return "", err
	}
	return string(body), nil
}