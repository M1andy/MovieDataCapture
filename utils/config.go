package utils

import (
	"fmt"
	"os"
	"path"
	"time"

	"gopkg.in/ini.v1"
)

var loadConfigOptions = ini.LoadOptions{
	IgnoreInlineComment: true,
}

// config里位于common分区里的设置
type commonConfig struct {
	MainMode                 int    `ini:"main_mode"`
	SourceFolder             string `ini:"source_folder"`
	FailedFolder             string `ini:"failed_output_folder"`
	SuccessFolder            string `ini:"success_output_folder"`
	LinkMode                 bool   `ini:"link_mode"`
	ScanHardLink             bool   `ini:"scan_hardlink"`
	Auto_exit                bool   `ini:"auto_exit"`
	TranslateToSC            bool   `ini:"translate_to_sc"`
	MultiThreading           bool   `ini:"multi_threading"`
	ActorGender              string `ini:"actor_gender"`
	DelEmptyFolder           bool   `ini:"del_empty_folder"`
	NfoSkipDays              int    `ini:"nfo_skip_days"`
	IgnoreFailedList         bool   `ini:"ignore_failed_list"`
	DownloadOnlyMissingFiles bool   `ini:"download_only_missing_images"`
	MappingTableValidity     int    `ini:"mapping_table_validity"`
	Sleep                    int    `ini:"sleep"`
}

// config里位于advanced_sleep分区里的设置
type advancedSleepConfig struct {
	StopCounter int       `ini:"stop_counter"`
	RerunDelay  time.Time `ini:"rerun_delay"`
}

// config里位于Proxy分区里的设置
type proxyConfig struct {
	ProxySwitch bool   `ini:"switch"`
	ProxyType   string `ini:"type"`
	Host        string `ini:"proxy"`
	Timeout     int    `ini:"timeout"`
	RetryCount  int    `ini:"retry"`
	CacertFile  string `ini:"cacert_file"`
}

type Config struct {
	Common        commonConfig
	AdvancedSleep advancedSleepConfig
	Proxy         proxyConfig
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GenerateConfigPath() ([]string, error) {
	var SearchPath []string
	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("get user home directory failed: %v", err)
	}

	currentPath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get current home directory failed: %v", err)
	}

	// 添加搜索config的路径
	// SearchPath = append(SearchPath, path.Join(currentPath, "config.ini"))
	SearchPath = append(SearchPath, path.Join(currentPath, "config.ini"))
	SearchPath = append(SearchPath, path.Join(homePath, "mdc.ini"))
	SearchPath = append(SearchPath, path.Join(homePath, ".mdc.ini"))
	SearchPath = append(SearchPath, path.Join(homePath, ".mdc/config.ini"))
	SearchPath = append(SearchPath, path.Join(homePath, ".config/mdc/config.ini"))

	return SearchPath, nil
}

func mapConfig(config *Config, rawData *ini.File) error {
	if err := rawData.Section("common").MapTo(&config.Common); err != nil {
		return fmt.Errorf("mapping config error: %v", err)
	}
	if err := rawData.Section("advanced_sleep").MapTo(&config.AdvancedSleep); err != nil {
		return fmt.Errorf("mapping config error: %v", err)
	}
	if err := rawData.Section("proxy").MapTo(&config.Proxy); err != nil {
		return fmt.Errorf("mapping config error: %v", err)
	}
	return nil
}

func LoadConfig() (*Config, error) {
	paths, err := GenerateConfigPath()
	if err != nil {
		return nil, err
	}

	var config = new(Config)
	var path string

	for _, path = range paths {
		if ok, err := PathExists(path); ok {
			break
		} else {
			return nil, err
		}
	}
	rawData, err := ini.LoadSources(loadConfigOptions, path)
	if err != nil {
		return nil, fmt.Errorf("loading config error: %v", err)
	}

	if err := mapConfig(config, rawData); err != nil {
		return nil, err
	}

	return config, nil

}

func init() {
}
