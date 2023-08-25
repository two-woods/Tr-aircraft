package command

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func ExecCommand(name string, arg ...string) (result string) {
	log.Infof("start exec Command %s,%v", name, arg)

	cmd := exec.Command(name, arg...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("combined out:\n%s\n", string(out))
		log.Errorf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

// ExecCommandFromPath 在指定目录执行命令
// command.ExecCommandFromPath(projectPath+"/"+dir.Name(), `bash`, `-c`, `git tag -l |grep `+tagRegular+`|xargs git tag -d`)
func ExecCommandFromPath(path string, name string, arg ...string) (result string) {
	log.Infof("在%s目录下进行命令执行", path)
	err := os.Chdir(path)
	if err != nil {
		log.Errorf("%+v", err)
	}
	dir, err := os.Getwd()
	log.Infof("切换目录完成,现在工作目录%s", dir)
	return ExecCommand(name, arg...)
}
