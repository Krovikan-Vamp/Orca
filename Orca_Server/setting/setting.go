package setting

import (
	"Orca_Server/sqlmgmt"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-ini/ini"
)

type commonConf struct {
	HttpPort  string
	CryptoKey string
}

var CommonSetting = &commonConf{}

type global struct {
	LocalHost      string // Local intranet IP
	ServerList     map[string]string
	ServerListLock sync.RWMutex
}

var GlobalSetting = &global{}

var cfg *ini.File

func Setup() {
	configFile := flag.String("c", "conf/app.ini", "-c conf/app.ini")
	addUserFlag := flag.Bool("au", false, "add user")
	delUserFlag := flag.Bool("du", false, "delete user")
	modUserFlag := flag.Bool("mu", false, "Change the password")
	flag.Parse()
	if *addUserFlag {
		addUser()
		os.Exit(0)
	}

	if *delUserFlag {
		delUser()
		os.Exit(0)
	}

	if *modUserFlag {
		modUserPwd()
		os.Exit(0)
	}

	var err error
	cfg, err = ini.Load(*configFile)
	if err != nil {
		log.Fatalf("setting.Setup, failed to parse 'conf/app.ini': %v", err)
	}

	mapTo("common", CommonSetting)
	if len(CommonSetting.CryptoKey) > 16 {
		CommonSetting.CryptoKey = CommonSetting.CryptoKey[:16]
	} else if len(CommonSetting.CryptoKey) < 16 {
		CommonSetting.CryptoKey = fmt.Sprintf("%016s", CommonSetting.CryptoKey)
	}

	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

func Default() {
	CommonSetting = &commonConf{
		HttpPort:  "6000",
		CryptoKey: "Adba723b7fe06819",
	}

	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

// mapTo maps the section to the given structure
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s error: %v", section, err)
	}
}

// getIntranetIp retrieves the local intranet IP
func getIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		// Check the IP address to determine if it's a loopback address
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func addUser() {
	username := ""
	password := ""
	repassword := ""
	userprompt := &survey.Input{
		Message: "Please enter the username: ",
	}
	survey.AskOne(userprompt, &username)
	pwdprompt := &survey.Password{
		Message: "Please enter the password: ",
	}
	survey.AskOne(pwdprompt, &password)
	pwdprompt = &survey.Password{
		Message: "Please enter the password again: ",
	}
	survey.AskOne(pwdprompt, &repassword)
	if password != repassword {
		log.Fatalf("The passwords do not match!")
		return
	}
	sqlmgmt.AddUser(username, password)
	log.Println("User added successfully!")
}

func delUser() {
	usernames := sqlmgmt.GetUsernames()
	selectUsername := ""
	prompt := &survey.Select{
		Message: "Please select the user to delete:",
		Options: usernames,
	}
	survey.AskOne(prompt, &selectUsername)
	sqlmgmt.DelUser(selectUsername)
	log.Println("User deleted successfully!")
}

func modUserPwd() {
	usernames := sqlmgmt.GetUsernames()
	selectUsername := ""
	password := ""
	repassword := ""
	prompt := &survey.Select{
		Message: "Please select the user:",
		Options: usernames,
	}
	survey.AskOne(prompt, &selectUsername)
	pwdprompt := &survey.Password{
		Message: "Please enter the password: ",
	}
	survey.AskOne(pwdprompt, &password)
	pwdprompt = &survey.Password{
		Message: "Please enter the password again: ",
	}
	survey.AskOne(pwdprompt, &repassword)
	if password != repassword {
		log.Fatalf("The passwords do not match!")
		return
	}
	sqlmgmt.ModUserPwd(selectUsername, password)
	log.Println("Password changed successfully!")
}
