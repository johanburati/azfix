package util

import (
    "bytes"
	"io/ioutil"
    "net"
    "log"
    "github.com/go-ini/ini"
)

func ReadOSRelease(configfile string) map[string]string {
    cfg, err := ini.Load(configfile)
    if err != nil {
        log.Fatal("Fail to read file: ", err)
    }

    ConfigParams := make(map[string]string)
    ConfigParams["PRETTY_NAME"] = cfg.Section("").Key("PRETTY_NAME").String()
    ConfigParams["NAME"] = cfg.Section("").Key("NAME").String()
    ConfigParams["VERSION_ID"] = cfg.Section("").Key("VERSION_ID").String()
    ConfigParams["VERSION"] = cfg.Section("").Key("VERSION").String()
    ConfigParams["VERSION_CODENAME"] = cfg.Section("").Key("VERSION_CODENAME").String()
    ConfigParams["ID"] = cfg.Section("").Key("ID").String()
    ConfigParams["ID_LIKE"] = cfg.Section("").Key("ID_LIKE").String()
    ConfigParams["HOME_URL"] = cfg.Section("").Key("HOME_URL").String()
    ConfigParams["SUPPORT_URL"] = cfg.Section("").Key("SUPPORT_URL").String()
    ConfigParams["BUG_REPORT_URL"] = cfg.Section("").Key("BUG_REPORT_URL").String() 
    ConfigParams["PRIVACY_POLICY_URL"] = cfg.Section("").Key("PRIVACY_POLICY_URL").String()
    ConfigParams["UBUNTU_CODENAME"] = cfg.Section("").Key("UBUNTU_CODENAME").String()
    return ConfigParams
}

func IsAzure() bool {
    if d, err := ioutil.ReadFile("/sys/class/dmi/id/chassis_asset_tag"); err == nil {
        return bytes.Compare(d, []byte("7783-7084-3265-9085-8269-3286-77\n")) == 0
    }
    return false
}

/*
func CheckTcpPort(host string, ports []string) {
    for _, port := range ports {
        timeout := time.Second
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
        if err != nil {
            fmt.Println("Connecting error:", err)
        }
        if conn != nil {
            defer conn.Close()
            fmt.Println("Opened", net.JoinHostPort(host, port))
        }
    }
}
*/

// CheckRepoUrl checks if the repo url is available
// curl -k https://rhui-1.microsoft.com/pulp/repos/microsoft-azure-rhel8/repodata/repomd.xml
//
func CheckRepoUrl(RepoUrl string) ([]net.IP, error) {
    ips, err := net.LookupIP(RepoUrl)
    if err != nil {
        return nil, err
    }
	return ips, nil
}
