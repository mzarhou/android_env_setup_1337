package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Item struct {
	name        string
	target_name string
	url         string
	command     string
}

func main() {
	log.Printf("Start")
	start := time.Now()

	// get username
	cmd := exec.Command("whoami")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	username := outb.String()
	username = username[:len(username)-1]
	download_path := "/Users/" + username + "/goinfre/temp_download_32234234223/"
	devtools_path := "/Users/" + username + "/goinfre/devtools/"
	apps_path := "/Users/" + username + "/goinfre/apps/"
	sdk_path := devtools_path + "Android/sdk"

	items := []Item{
		{
			name:        "flutter.zip",
			target_name: "flutter",
			command:     "unzip " + download_path + "flutter.zip",
			url:         "https://storage.googleapis.com/flutter_infra_release/releases/stable/macos/flutter_macos_2.5.3-stable.zip",
		},
		{
			name:        "jdk.tar.gz",
			target_name: "jdk",
			command:     "tar -xvf " + download_path + "jdk.tar.gz",
			url:         "https://download.oracle.com/java/17/latest/jdk-17_macos-x64_bin.tar.gz",
		},
		{
			name:        "gradle.zip",
			target_name: "gradle",
			command:     "unzip " + download_path + "gradle.zip",
			url:         "https://downloads.gradle-dn.com/distributions/gradle-7.3.1-all.zip",
		},
		{
			name:        "android-studio.dmg",
			target_name: "android-studio",
			command:     "hdiutil attach " + download_path + "android-studio.dmg",
			url:         "https://redirector.gvt1.com/edgedl/android/studio/install/2020.3.1.25/android-studio-2020.3.1.25-mac.dmg",
		},
	}

	if err := makeDir(download_path); err != nil {
		panic(err)
	}

	if err := makeDir(devtools_path); err != nil {
		panic(err)
	}

	if err := makeDir(apps_path); err != nil {
		panic(err)
	}

	if err := os.Chdir(devtools_path); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, item := range items {
		go func(_item Item) {
			defer wg.Done()
			work(_item, download_path, apps_path)
		}(item)
	}
	wg.Wait()

	// remove download folder
	cmd = exec.Command("bash", "-c", "rm -rf "+download_path)
	cmd.Run()

	// create sdk folder
	cmd = exec.Command("bash", "-c", "mkdir -p "+sdk_path)
	cmd.Run()

	elapsed := time.Since(start)
	log.Printf("task took %s", elapsed)
}

func work(item Item, download_path string, apps_path string) {
	// download file
	fmt.Println("start downloading..." + item.url)
	err := DownloadFile(download_path+item.name, item.url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + item.url)
	// extract
	fmt.Println("extracting " + item.name + "...")
	cmd := exec.Command("bash", "-c", item.command)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	// rename
	var command string
	if item.name != "android-studio.dmg" {
		command = "mv " + item.target_name + "* " + item.target_name
	} else {
		command = "cp -R /Volumes/Android\\ Studio\\ -\\ Arctic\\ Fox\\ \\|\\ 2020.3.1\\ Patch\\ 3/Android\\ Studio.app " + apps_path
	}
	cmd = exec.Command("bash", "-c", command)
	cmd.Run()
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func makeDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		return err
	}
	return nil
}
