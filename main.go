package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "os/exec"
    "time"

    yaml "gopkg.in/yaml.v2"
)

var icons = []string{
    "face-angel",
    "face-sad",
    "face-angry",
    "face-sick",
    "face-cool",
    "face-smile-big",
    "face-crying",
    "face-smile",
    "face-devilish",
    "face-smirk",
    "face-embarrassed",
    "face-surprise",
    "face-glasses",
    "face-tired",
    "face-kiss",
    "face-uncertain",
    "face-laugh",
    "face-wink",
    "face-monkey",
    "face-worried",
    "face-plain",
    "faceaspberry",
}

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
    tb := t.Add(time.Duration(-15) * time.Minute)
    if strTime == t.Format("15:04") || strTime == tb.Format("15:04") {
        return true
    }

    return false
}

func notify(t Task) {
    rand.Seed(time.Now().Unix())
    n := rand.Int() % len(icons)
    cmd := exec.Command(
        "/usr/bin/notify-send",
        "-i",
        icons[n],
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
