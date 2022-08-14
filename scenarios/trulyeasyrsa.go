package scenarios

import (
	"bytes"
	"flag"
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
	caPwd               string = "CA password"
	serverPwd           string = "Server's password"
)

func getEasyRsaPath() (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		return "", err
	}
	return path + "/easyrsa", nil
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

func (s *EasyRsaScenarioImpl) Execute() {

	cP := flag.String("cp", "", caPwd)
	sP := flag.String("sp", "", serverPwd)

	if s.CertName == "" {
		fmt.Println("")
		return
	}

	if !checkEnv() {
		return
	}

	easyRsaPath, err := getEasyRsaPath()
	if err != nil {
		fmt.Printf("Error on call os.Getwd() %s\n", err)
		return
	}

	if s.CertType == common.CreateServerKey {
		buildCaArgs := []string{"build-ca"}
		if *cP == "" {
			buildCaArgs = append(buildCaArgs, "nopass")
		}

		buildServerFullArgs := []string{"build-server-full", s.CertName}
		if *sP == "" {
			buildServerFullArgs = append(buildServerFullArgs, "nopass")
		}

		workDir, _ := os.Getwd()

		stages := []CommandParams{
			{&easyRsaPath, []string{"init-pki"}, cP},
			{&cpCmd, []string{"-r", workDir + "/easy-rsa/x509-types", workDir + "/pki"}, cP},
			{&cpCmd, []string{workDir + "/easy-rsa/openssl-easyrsa.cnf", workDir + "/pki"}, cP},
			{&easyRsaPath, buildCaArgs, cP},
			{&easyRsaPath, []string{"gen-dh"}, cP},
			{&openvpnCmd, []string{"--genkey", "--secret", workDir + "/pki/ta.key"}, cP},
			{&easyRsaPath, []string{"gen-crl"}, cP},
			{&easyRsaPath, buildServerFullArgs, cP},
		}
		for _, stage := range stages {
			if !executeStage(stage) {
				fmt.Println("The programm will be stopped")
				return
			}
		}
	}

	if s.CertType == common.CreateClientKey {
		buildClientFullArgs := []string{"build-client-full", s.CertName}
		if *sP == "" {
			buildClientFullArgs = append(buildClientFullArgs, "nopass")
		}
		params := CommandParams{&easyRsaPath, buildClientFullArgs, cP}
		if !executeStage(params) {
			fmt.Println("")
		}
	}

}
