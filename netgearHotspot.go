package main

import (
	"encoding/json"
	"fmt"
	"github.com/andelf/go-curl"
	"github.com/johnmccabe/go-bitbar"
	"log"
	"net/http"
)

// built with https://mholt.github.io/json-to-go/
// full model.json return below from Netgear hotspot
type NetgearReturn struct {
	Custom struct {
		LastWifiChan      int  `json:"lastWifiChan"`
		HiddenMenuEnabled bool `json:"hiddenMenuEnabled"`
		HideAdminPassword bool `json:"hideAdminPassword"`
		End               int  `json:"end"`
	} `json:"custom"`
	Webd struct {
		AdminPassword     string `json:"adminPassword"`
		HideAdminPassword bool   `json:"hideAdminPassword"`
		HintNumber        int    `json:"hintNumber"`
		End               string `json:"end"`
	} `json:"webd"`
	Lcd struct {
		BacklightActive bool   `json:"backlightActive"`
		End             string `json:"end"`
	} `json:"lcd"`
	Sim struct {
		Pin struct {
			Mode  string `json:"mode"`
			Retry int    `json:"retry"`
			End   string `json:"end"`
		} `json:"pin"`
		Puk struct {
			Retry int `json:"retry"`
		} `json:"puk"`
		Mep struct {
			End string `json:"end"`
		} `json:"mep"`
		PhoneNumber   string `json:"phoneNumber"`
		Iccid         string `json:"iccid"`
		Imsi          string `json:"imsi"`
		Status        string `json:"status"`
		SprintSimLock int    `json:"sprintSimLock"`
		End           string `json:"end"`
	} `json:"sim"`
	Sms struct {
		Ready bool   `json:"ready"`
		End   string `json:"end"`
	} `json:"sms"`
	Session struct {
		UserRole            string `json:"userRole"`
		Lang                string `json:"lang"`
		HintDisplayPassword string `json:"hintDisplayPassword"`
		SecToken            string `json:"secToken"`
	} `json:"session"`
	General struct {
		DefaultLanguage    string `json:"defaultLanguage"`
		PRIid              string `json:"PRIid"`
		Activated          bool   `json:"activated"`
		TCAaccepted        bool   `json:"TCAaccepted"`
		DevTemperature     int    `json:"devTemperature"`
		VerMajor           int    `json:"verMajor"`
		VerMinor           int    `json:"verMinor"`
		Environment        string `json:"environment"`
		CurrTime           int    `json:"currTime"`
		TimeZoneOffset     int    `json:"timeZoneOffset"`
		DeviceName         string `json:"deviceName"`
		UseMetricSystem    bool   `json:"useMetricSystem"`
		FactoryResetStatus string `json:"factoryResetStatus"`
		SetupCompleted     bool   `json:"setupCompleted"`
		WarrantyDateCode   string `json:"warrantyDateCode"`
		LanguageSelected   bool   `json:"languageSelected"`
		UpTime             int    `json:"upTime"`
		SystemAlertList    struct {
			List []struct {
			} `json:"list"`
			Count int `json:"count"`
		} `json:"systemAlertList"`
		APIVersion        string `json:"apiVersion"`
		CompanyName       string `json:"companyName"`
		ConfigURL         string `json:"configURL"`
		ProfileURL        string `json:"profileURL"`
		PinChangeURL      string `json:"pinChangeURL"`
		PortCfgURL        string `json:"portCfgURL"`
		PortFilterURL     string `json:"portFilterURL"`
		WifiACLURL        string `json:"wifiACLURL"`
		SupportedLangList []struct {
			ID        string `json:"id,omitempty"`
			IsCurrent string `json:"isCurrent,omitempty"`
			IsDefault string `json:"isDefault,omitempty"`
			Label     string `json:"label,omitempty"`
			Token1    string `json:"token1,omitempty"`
			Token2    string `json:"token2,omitempty"`
		} `json:"supportedLangList"`
	} `json:"general"`
	Power struct {
		PMState string `json:"PMState"`
		SmState string `json:"SmState"`
		AutoOff struct {
			OnUSBdisconnect struct {
				End string `json:"end"`
			} `json:"onUSBdisconnect"`
			OnIdle struct {
				Timer struct {
					End string `json:"end"`
				} `json:"timer"`
			} `json:"onIdle"`
		} `json:"autoOff"`
		Standby struct {
			OnIdle struct {
				Timer struct {
					End string `json:"end"`
				} `json:"timer"`
			} `json:"onIdle"`
		} `json:"standby"`
		AutoOn struct {
			End string `json:"end"`
		} `json:"autoOn"`
		BatteryTemperature  int    `json:"batteryTemperature"`
		BatteryVoltage      int    `json:"batteryVoltage"`
		BattChargeLevel     int    `json:"battChargeLevel"`
		BattChargeSource    string `json:"battChargeSource"`
		BatteryState        string `json:"batteryState"`
		BattChargeAlgorithm string `json:"battChargeAlgorithm"`
		Charging            bool   `json:"charging"`
		DeviceTempCritical  bool   `json:"deviceTempCritical"`
		ResetRequired       string `json:"resetRequired"`
		WifiOff             struct {
			End string `json:"end"`
		} `json:"wifiOff"`
		Boost struct {
			CableConnected bool   `json:"cableConnected"`
			End            string `json:"end"`
		} `json:"boost"`
		Lpm bool   `json:"lpm"`
		End string `json:"end"`
	} `json:"power"`
	Wwan struct {
		PrlVersion               int    `json:"prlVersion"`
		LTEBandPriority          string `json:"LTEBandPriority"`
		NetScanStatus            string `json:"netScanStatus"`
		LTEeHRPDConfig           string `json:"LTEeHRPDConfig"`
		RoamingEnhancedIndicator int    `json:"roamingEnhancedIndicator"`
		RoamingMode              string `json:"roamingMode"`
		RoamingGuardDom          string `json:"roamingGuardDom"`
		RoamingGuardIntl         string `json:"roamingGuardIntl"`
		RoamingType              string `json:"roamingType"`
		RoamMenuDisplay          bool   `json:"roamMenuDisplay"`
		AutoBandRegionChanged    bool   `json:"autoBandRegionChanged"`
		CurrentNWserviceType     string `json:"currentNWserviceType"`
		RegisterRejectCode       int    `json:"registerRejectCode"`
		NetRegMode               string `json:"netRegMode"`
		IPv6                     string `json:"IPv6"`
		Roaming                  bool   `json:"roaming"`
		IP                       string `json:"IP"`
		RegisterNetworkDisplay   string `json:"registerNetworkDisplay"`
		RAT                      string `json:"RAT"`
		BandRegion               []struct {
			Index   int    `json:"index,omitempty"`
			Name    string `json:"name,omitempty"`
			Current bool   `json:"current,omitempty"`
		} `json:"bandRegion"`
		Autoconnect string `json:"autoconnect"`
		Profile     struct {
			PromptForApnSelection bool   `json:"promptForApnSelection"`
			End                   string `json:"end"`
		} `json:"profile"`
		DataUsage struct {
			Total struct {
				LteBillingTx  int    `json:"lteBillingTx"`
				LteBillingRx  int    `json:"lteBillingRx"`
				CdmaBillingTx int    `json:"cdmaBillingTx"`
				CdmaBillingRx int    `json:"cdmaBillingRx"`
				GwBillingTx   int    `json:"gwBillingTx"`
				GwBillingRx   int    `json:"gwBillingRx"`
				LteLifeTx     int    `json:"lteLifeTx"`
				LteLifeRx     int    `json:"lteLifeRx"`
				CdmaLifeTx    int    `json:"cdmaLifeTx"`
				CdmaLifeRx    int    `json:"cdmaLifeRx"`
				GwLifeTx      int    `json:"gwLifeTx"`
				GwLifeRx      int    `json:"gwLifeRx"`
				End           string `json:"end"`
			} `json:"total"`
			Server struct {
				AccountType    string `json:"accountType"`
				SubAccountType string `json:"subAccountType"`
				End            string `json:"end"`
			} `json:"server"`
			ServerDataRemaining       int    `json:"serverDataRemaining"`
			ServerDataTransferred     int64  `json:"serverDataTransferred"`
			ServerDataTransferredIntl int    `json:"serverDataTransferredIntl"`
			ServerDataValidState      string `json:"serverDataValidState"`
			ServerDaysLeft            int    `json:"serverDaysLeft"`
			ServerErrorCode           string `json:"serverErrorCode"`
			ServerLowBalance          bool   `json:"serverLowBalance"`
			ServerMsisdn              string `json:"serverMsisdn"`
			ServerRechargeURL         string `json:"serverRechargeUrl"`
			DataWarnEnable            bool   `json:"dataWarnEnable"`
			PlanSize                  int    `json:"planSize"`
			PlanDescription           string `json:"planDescription"`
			PrepaidStackedPlans       int    `json:"prepaidStackedPlans"`
			PrepaidStackedPlansIntl   int    `json:"prepaidStackedPlansIntl"`
			PrepaidAccountState       string `json:"prepaidAccountState"`
			AccountType               string `json:"accountType"`
			Share                     struct {
				Enabled               bool   `json:"enabled"`
				DataTransferredOthers int    `json:"dataTransferredOthers"`
				LastSync              string `json:"lastSync"`
				End                   string `json:"end"`
			} `json:"share"`
			Generic struct {
				DataLimitValid         bool   `json:"dataLimitValid"`
				UsageHighWarning       int    `json:"usageHighWarning"`
				LastSucceeded          string `json:"lastSucceeded"`
				BillingDay             int    `json:"billingDay"`
				NextBillingDate        string `json:"nextBillingDate"`
				LastSync               string `json:"lastSync"`
				BillingCycleRemainder  int    `json:"billingCycleRemainder"`
				BillingCycleLimit      int    `json:"billingCycleLimit"`
				DataTransferred        int64  `json:"dataTransferred"`
				DataTransferredRoaming int    `json:"dataTransferredRoaming"`
				LastReset              string `json:"lastReset"`
				UserDisplayFormat      string `json:"userDisplayFormat"`
				End                    string `json:"end"`
			} `json:"generic"`
		} `json:"dataUsage"`
		NetManualNoCvg       bool   `json:"netManualNoCvg"`
		Connection           string `json:"connection"`
		ConnectionType       string `json:"connectionType"`
		CurrentPSserviceType string `json:"currentPSserviceType"`
		Ca                   struct {
			SCCcount int    `json:"SCCcount"`
			End      string `json:"end"`
		} `json:"ca"`
		ConnectionText  string `json:"connectionText"`
		SessDuration    int    `json:"sessDuration"`
		SessStartTime   int    `json:"sessStartTime"`
		DataTransferred struct {
			Totalb string `json:"totalb"`
			Rxb    string `json:"rxb"`
			Txb    string `json:"txb"`
		} `json:"dataTransferred"`
		SignalStrength struct {
			Rssi int    `json:"rssi"`
			Rscp int    `json:"rscp"`
			Ecio int    `json:"ecio"`
			Rsrp int    `json:"rsrp"`
			Rsrq int    `json:"rsrq"`
			Bars int    `json:"bars"`
			Sinr int    `json:"sinr"`
			End  string `json:"end"`
		} `json:"signalStrength"`
	} `json:"wwan"`
	Wwanadv struct {
		CurBand           string `json:"curBand"`
		RadioQuality      int    `json:"radioQuality"`
		Country           string `json:"country"`
		RAC               int    `json:"RAC"`
		LAC               int    `json:"LAC"`
		MCC               string `json:"MCC"`
		MNC               string `json:"MNC"`
		MNCFmt            int    `json:"MNCFmt"`
		CellID            int    `json:"cellId"`
		ChanID            int    `json:"chanId"`
		PrimScode         int    `json:"primScode"`
		PlmnSrvErrBitMask int    `json:"plmnSrvErrBitMask"`
		ChanIDUl          int    `json:"chanIdUl"`
		TxLevel           int    `json:"txLevel"`
		RxLevel           int    `json:"rxLevel"`
		End               string `json:"end"`
	} `json:"wwanadv"`
	Wifi struct {
		ClientCount int    `json:"clientCount"`
		Country     string `json:"country"`
		Wps         struct {
			End string `json:"end"`
		} `json:"wps"`
		Guest struct {
			MaxClientCnt   int    `json:"maxClientCnt"`
			Enabled        bool   `json:"enabled"`
			SSID           string `json:"SSID"`
			AccessProfile  string `json:"accessProfile"`
			TimerEnable    bool   `json:"timerEnable"`
			TimerTimestamp int    `json:"timerTimestamp"`
			TimerValue     int    `json:"timerValue"`
			Chan           int    `json:"chan"`
			Mode           string `json:"mode"`
			DHCP           struct {
				Range struct {
					End string `json:"end"`
				} `json:"range"`
			} `json:"DHCP"`
		} `json:"guest"`
		End string `json:"end"`
	} `json:"wifi"`
	Router struct {
		HostName               string `json:"hostName"`
		DomainName             string `json:"domainName"`
		IPPassThroughEnabled   bool   `json:"ipPassThroughEnabled"`
		IPPassThroughSupported bool   `json:"ipPassThroughSupported"`
		ClientList             struct {
			List []struct {
				IP     string `json:"IP,omitempty"`
				MAC    string `json:"MAC,omitempty"`
				Name   string `json:"name,omitempty"`
				Media  string `json:"media,omitempty"`
				Source string `json:"source,omitempty"`
			} `json:"list"`
			Count int `json:"count"`
		} `json:"clientList"`
		End string `json:"end"`
	} `json:"router"`
	Cradle struct {
		Mode      bool   `json:"mode"`
		SmartMode bool   `json:"smartMode"`
		URL       string `json:"url"`
		End       string `json:"end"`
	} `json:"cradle"`
	Accesscontrol struct {
		Nlpc struct {
			End int `json:"end"`
		} `json:"nlpc"`
		Blocksites struct {
			End int `json:"end"`
		} `json:"blocksites"`
		SchedulerBlockingActive bool `json:"schedulerBlockingActive"`
		End                     int  `json:"end"`
	} `json:"accesscontrol"`
	Fota struct {
		Fwupdater struct {
			End string `json:"end"`
		} `json:"fwupdater"`
	} `json:"fota"`
	UI struct {
		PromptActivation bool `json:"promptActivation"`
		End              int  `json:"end"`
	} `json:"ui"`
}

func getCookie() string {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	u := "http://192.168.1.1"
	user := "Admin"
	pass := "getoffmylawn"
	easy.Setopt(curl.OPT_URL, u)
	easy.Setopt(curl.OPT_NOBODY, 1)
	easy.Setopt(curl.OPT_USERNAME, user)
	easy.Setopt(curl.OPT_PASSWORD, pass)
	easy.Setopt(curl.OPT_FOLLOWLOCATION, 1)
	easy.Setopt(curl.OPT_COOKIEFILE, "./hotcookie.jar")
	easy.Setopt(curl.OPT_COOKIEJAR, "./hotcookie.jar")

	easy.Perform()

	return ""
}

func getJson() string {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	u := "http://192.168.1.1/api/model.json"
	user := "Admin"
	pass := "getoffmylawn"
	easy.Setopt(curl.OPT_URL, u)
	easy.Setopt(curl.OPT_NOBODY, 1)
	easy.Setopt(curl.OPT_USERNAME, user)
	easy.Setopt(curl.OPT_PASSWORD, pass)
	easy.Setopt(curl.OPT_FOLLOWLOCATION, 1)
	easy.Setopt(curl.OPT_COOKIEFILE, "./hotcookie.jar")
	easy.Setopt(curl.OPT_COOKIEJAR, "./hotcookie.jar")

	easy.Perform()

	return ""
}

func main() {
	url := fmt.Sprintf("http://192.168.1.1/api/model.json")

	// get the cookie
	getCookie()

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record NetgearReturn

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	const logoBase64  = "R0lGODlhDgAOAIABAAAAAAAAACH/C1hNUCBEYXRhWE1QPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4gPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNS4zLWMwMTEgNjYuMTQ1NjYxLCAyMDEyLzAyLzA2LTE0OjU2OjI3ICAgICAgICAiPiA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPiA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIiB4bWxuczpzdFJlZj0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3NUeXBlL1Jlc291cmNlUmVmIyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ1M2IChNYWNpbnRvc2gpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOkNDNUMyMDNGMDQyODExRTJCMDkwQkU4MjIwMDc5RTY2IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOkNDNUMyMDQwMDQyODExRTJCMDkwQkU4MjIwMDc5RTY2Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6Q0M1QzIwM0QwNDI4MTFFMkIwOTBCRTgyMjAwNzlFNjYiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6Q0M1QzIwM0UwNDI4MTFFMkIwOTBCRTgyMjAwNzlFNjYiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4B//79/Pv6+fj39vX08/Lx8O/u7ezr6uno5+bl5OPi4eDf3t3c29rZ2NfW1dTT0tHQz87NzMvKycjHxsXEw8LBwL++vby7urm4t7a1tLOysbCvrq2sq6qpqKempaSjoqGgn56dnJuamZiXlpWUk5KRkI+OjYyLiomIh4aFhIOCgYB/fn18e3p5eHd2dXRzcnFwb25tbGtqaWhnZmVkY2JhYF9eXVxbWllYV1ZVVFNSUVBPTk1MS0pJSEdGRURDQkFAPz49PDs6OTg3NjU0MzIxMC8uLSwrKikoJyYlJCMiISAfHh0cGxoZGBcWFRQTEhEQDw4NDAsKCQgHBgUEAwIBAAAh+QQBAAABACwAAAAADgAOAAACHoyPqQmw7F4LU7qKpo75+rNRnxFu5decoqia6Nt6BQA7"

	app := bitbar.New()
	//app.StatusLine(fmt.Sprintf("%+v", record.Wwan.SignalStrength.Bars)).TemplateImage(logoBase64)
	app.StatusLine("").TemplateImage(logoBase64)

	submenu := app.NewSubMenu()
	submenu.Line(fmt.Sprint())
	submenu.Line(fmt.Sprintf("Band: %s",record.Wwanadv.CurBand))
	submenu.Line(fmt.Sprintf("Bars: %d",record.Wwan.SignalStrength.Bars))
	submenu.Line(fmt.Sprintf("RSRP: %d",record.Wwan.SignalStrength.Rsrp))
	submenu.Line(fmt.Sprintf("SINR: %d",record.Wwan.SignalStrength.Sinr))
	submenu.Line(fmt.Sprintf("Batt: %d%%",record.Power.BattChargeLevel))
	submenu.Line(fmt.Sprintf("Role: %s",record.Session.UserRole))
	submenu.Line(fmt.Sprintf("Clients: %d",record.Wifi.ClientCount))

	app.Render()
}