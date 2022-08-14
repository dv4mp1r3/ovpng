package scenarios

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/dv4mp1r3/ovpngen/common"
)

type EasyRsaScenario interface {
	Scenario
}

type EasyRsaScenarioImpl struct {
	args     []string
	CertName string
	CertType string
	CaPwd    *string
	CertPwd  *string
}

type CommandParams struct {
	command       *string
	args          []string
	inputPassword *string
}

var (
	openvpnCmd string = "openvpn"
	cpCmd      string = "cp"
)

const (
	ScenarioEasyRsaName string = "easyrsa"
	exitErrorMsg        string = "The programm will be stopped"
)

func getEasyRsaPath() (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		return "", err
	}
	return path + "/easyrsa", nil
}

func (s *EasyRsaScenarioImpl) genCertArgsByType(pass string) []string {
	buildType := "build-server-full"
	if s.CertType == common.CreateClientKey {
		buildType = "build-client-full"
	}
	args := []string{buildType, s.CertName}
	if pass == "" {
		args = append(args, "nopass")
	}
	return args
}

func executeStage(params CommandParams) bool {
	cmd := exec.Command(*params.command, params.args...)
	in, err := cmd.StdinPipe()
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "EASYRSA_BATCH="+os.Getenv("EASYRSA_BATCH"))
	cmd.Env = append(cmd.Env, "EASYRSA_REQ_CN="+os.Getenv("EASYRSA_REQ_CN"))
	if err != nil {
		return false
	}
	go func() {
		defer in.Close()
		if *params.inputPassword == "" {
			return
		}
		fmt.Fprintln(in, *params.inputPassword)
		time.Sleep(1 * time.Second)
		fmt.Fprintln(in, *params.inputPassword)
	}()

	var b bytes.Buffer
	cmd.Stdout, cmd.Stderr = &b, &b

	if err := cmd.Start(); err != nil {
		fmt.Printf("cmd.Start err %s\n", err)
		return false
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("cmd.Wait err %s\n", err)
		fmt.Printf("%s\n", b.String())
		return false
	}
	fmt.Printf("%s\n", b.String())
	return true
}

func checkEnv() bool {
	if os.Getenv("EASYRSA_BATCH") != "yes" {
		fmt.Println("env EASYRSA_BATCH isnt't set or not 'yes'")
		return false
	}
	if os.Getenv("EASYRSA_REQ_CN") == "" {
		fmt.Println("env EASYRSA_REQ_CN isnt't set")
		return false
	}
	return true
}

func (s *EasyRsaScenarioImpl) genOnlyClientCertificate(ccPwd *string, easyRsaPath *string) {
	buildClientFullArgs := s.genCertArgsByType(*ccPwd)
	params := CommandParams{easyRsaPath, buildClientFullArgs, ccPwd}
	if !executeStage(params) {
		fmt.Println(exitErrorMsg)
	}
}

func (s *EasyRsaScenarioImpl) setupPki(caPwd *string, srvPwd *string, easyRsaPath *string) {
	buildCaArgs := []string{"build-ca"}
	if *caPwd == "" {
		buildCaArgs = append(buildCaArgs, "nopass")
	}

	buildServerFullArgs := s.genCertArgsByType(*srvPwd)

	workDir, _ := os.Getwd()

	stages := []CommandParams{
		{easyRsaPath, []string{"init-pki"}, caPwd},
		{&cpCmd, []string{"-r", workDir + "/easy-rsa/x509-types", workDir + "/pki"}, caPwd},
		{&cpCmd, []string{workDir + "/easy-rsa/openssl-easyrsa.cnf", workDir + "/pki"}, caPwd},
		{easyRsaPath, buildCaArgs, caPwd},
		{easyRsaPath, []string{"gen-dh"}, caPwd},
		{&openvpnCmd, []string{"--genkey", "--secret", workDir + "/pki/ta.key"}, caPwd},
		{easyRsaPath, []string{"gen-crl"}, caPwd},
		{easyRsaPath, buildServerFullArgs, caPwd},
	}
	for _, stage := range stages {
		if !executeStage(stage) {
			fmt.Println(exitErrorMsg)
			return
		}
	}
}

func (s *EasyRsaScenarioImpl) Execute() {

	if !checkEnv() {
		return
	}

	easyRsaPath, err := getEasyRsaPath()
	if err != nil {
		fmt.Printf("Error on call os.Getwd() %s\n", err)
		return
	}

	if s.CertType == common.CreateServerKey {
		s.setupPki(s.CaPwd, s.CertPwd, &easyRsaPath)
	} else if s.CertType == common.CreateClientKey {
		s.genOnlyClientCertificate(s.CertPwd, &easyRsaPath)
	}

}
