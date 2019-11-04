// This file is auto-generated, don't edit it. Thanks.

package baseclient

import (
	"github.com/alibabacloud-go/tea/tea"
)

type BaseClient struct {
	Domain    string `json:"domain xml:"domain"`
	AuthToken string `json:"authToken xml:"authToken"`
}

func (client *BaseClient) InitClient(config map[string]string) error {
	panic("the constructor is un-implemented!")
}

func (client *BaseClient) ToJSON(a map[string]string) string {
	panic("the method is un-implemented!")
}

func (client *BaseClient) ReadJSON(a *tea.Response) map[string]string {
	panic("the method is un-implemented!")
}

func (client *BaseClient) ToQuery(a map[string]string) map[string]interface{} {
	panic("the method is un-implemented!")
}
