package link2xray

import (
	"github.com/xtls/libxray/share"
	"github.com/xtls/xray-core/infra/conf"
	"gitlab.com/tozd/go/errors"
)

func SharedLinkToXrayConfig(link string) (name *string, config *conf.Config, err error) {
	config, err = share.ConvertShareLinksToXrayJson(link)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	name = config.OutboundConfigs[0].SendThrough
	config.OutboundConfigs[0].SendThrough = nil

	return name, config, err
}
