package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os/exec"
    "time"

    yaml "gopkg.in/yaml.v2"
)

type Task struct {
    Job     string `yaml:"job"`
    Time    string `yaml:"time"`
    Message string `yaml:"message"`
}

type YamlFile struct {
    Tasks []Task
}

func (y *YamlFile) readFile() {
    file, err := ioutil.ReadFile("task.yml")
    if err != nil {
        log.Fatal("Error loading yml file")
    }
    errY := yaml.Unmarshal(file, &y)
    if errY != nil {
        log.Fatalf("error: %v", errY)
    }
}

func checkTime(strTime string) bool {
    loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
    t := time.Now().In(loc)
    if strTime == t.Format("15:04") {
        return true
    }

    return false
}

func notify(t Task) {
    cmd := exec.Command(
        "notify-send",
        fmt.Sprintf("%s at %s", t.Job, t.Time),
        t.Message,
    )

    _, err := cmd.Output()
    if err != nil {
        log.Fatalf("error: %v", err)
        return
    }

}

func main() {
    y := YamlFile{}
    y.readFile()
    for _, task := range y.Tasks {
        if checkTime(task.Time) {
            notify(task)
        }
    }
}
