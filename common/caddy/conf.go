package caddy

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/pydio/cells/common/caddy/maintenance"
	"github.com/pydio/cells/common/config"
	"github.com/pydio/cells/discovery/install/assets"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/pydio/cells/common/crypto/providers"
	"github.com/pydio/cells/common/proto/install"
)

var (
	DefaultCaUrl        = "https://acme-v02.api.letsencrypt.org/directory"
	DefaultCaStagingUrl = "https://acme-staging-v02.api.letsencrypt.org/directory"
)

type SiteConf struct {
	*install.ProxyConfig
	// Parsed values from proto oneOf
	TLS     string
	TLSCert string
	TLSKey  string
	// Parsed External host if any
	ExternalHost string
	// Custom Root for this site
	WebRoot string
}

func (s SiteConf) Redirects() map[string]string {
	rr := make(map[string]string)
	for _, bind := range s.GetBinds() {
		parts := strings.Split(bind, ":")
		var host, port string
		if len(parts) == 2 {
			host = parts[0]
			port = parts[1]
			if host == "" {
				continue
			}
		} else {
			host = bind
		}
		if port == "" {
			rr["http://"+host] = "https://" + host
		} else if port == "80" {
			continue
		} else {
			rr["http://"+host] = "https://" + host + ":" + port
		}
	}

	return rr
}

func SiteConfFromProxyConfig(pc *install.ProxyConfig) (SiteConf, error) {
	bc := SiteConf{
		ProxyConfig: proto.Clone(pc).(*install.ProxyConfig),
	}
	if pc.ReverseProxyURL != "" {
		if u, e := url.Parse(pc.ReverseProxyURL); e == nil {
			bc.ExternalHost = u.Host
		}
	}
	if bc.TLSConfig == nil {
		for i, b := range bc.Binds {
			bc.Binds[i] = "http://" + b
		}
	} else {
		switch v := bc.TLSConfig.(type) {
		case *install.ProxyConfig_Certificate, *install.ProxyConfig_SelfSigned:
			certFile, keyFile, err := providers.LoadCertificates(pc)
			if err != nil {
				return bc, err
			}
			bc.TLSCert = certFile
			bc.TLSKey = keyFile
		case *install.ProxyConfig_LetsEncrypt:
			bc.TLS = v.LetsEncrypt.Email
		}
	}
	if bc.Maintenance {
		mDir, e := GetMaintenanceRoot()
		if e != nil {
			return bc, e
		}
		bc.WebRoot = mDir
	}
	return bc, nil
}

func SitesToCaddyConfigs(sites []*install.ProxyConfig) (caddySites []SiteConf, er error) {
	for _, proxyConfig := range sites {
		if bc, er := SiteConfFromProxyConfig(proxyConfig); er == nil {
			caddySites = append(caddySites, bc)
			if proxyConfig.HasTLS() && proxyConfig.GetLetsEncrypt() != nil {
				le := proxyConfig.GetLetsEncrypt()
				if le.AcceptEULA {
					caddytls.Agreed = true
				}
				if le.StagingCA {
					caddytls.DefaultCAUrl = DefaultCaStagingUrl
				} else {
					caddytls.DefaultCAUrl = DefaultCaUrl
				}
			}
		} else {
			return caddySites, er
		}
	}
	return caddySites, nil
}

var maintenanceDir string

func GetMaintenanceRoot() (string, error) {
	if maintenanceDir != "" {
		return maintenanceDir, nil
	}
	dir, err := assets.GetAssets("./maintenance/src")
	if err != nil {
		dir = filepath.Join(config.ApplicationWorkingDir(), "static", "maintenance")
		if err, _, _ := assets.RestoreAssets(dir, maintenance.PydioMaintenanceBox, nil); err != nil {
			return "", errors.Wrap(err, "could not restore maintenance package")
		}
	}
	maintenanceDir = dir
	return dir, nil
}
