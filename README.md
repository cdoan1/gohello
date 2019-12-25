# gohello

## Summary

---

* use this main package to call sub package runner to generate screen shots - to detect UI outage
* read a yaml file
* unmarshal yaml data
* marshal yaml data
* works with simple data structure
* simple main package go program

## Data Structure

implements a cluster data structure, which is a dictionary, or map
then a clusters data structure, which is an array of cluster

## Project Structure

---

reference: https://github.com/golang-standards/project-layout

/cmd - main application for this project

/pkg - reusable code

/internal - non-reusable code

/vendor - application dependency

/githooks - github webhooks

## Usage

```bash
go build cmd/gohello.go

./gohello

2019/12/24 17:34:45 hello world!
2019/12/24 17:34:45 Reading clusters from /Users/cd1/go/src/github.com/cdoan1/gohello/data.yaml
2019/12/24 17:34:45 name        scale1
2019/12/24 17:34:45 calling chrome runner ...
2019/12/24 17:34:47 login screen - arguments
2019/12/24 17:34:49 login screen - welcome
2019/12/24 17:34:59 overview view -
2019/12/24 17:35:04 cluster view -
2019/12/24 17:35:08 policies view -
2019/12/24 17:35:11 name        offline
2019/12/24 17:35:11 calling chrome runner ...
2019/12/24 17:35:13 login screen - arguments
2019/12/24 17:35:14 login screen - welcome
2019/12/24 17:35:23 overview view -
2019/12/24 17:35:33 cluster view -
2019/12/24 17:35:36 policies view -
```

This generates the screen shot in the directory `./output/<cluster>/*.png`

Then you can use this project to compare if there are any changes in the screenshots.

```bash
-rw-r--r--  1 cd1  staff   66493 Dec 24 17:35 clusters.png
-rw-r--r--  1 cd1  staff   39495 Dec 24 17:34 fullScreenshot.png
-rw-r--r--  1 cd1  staff   41420 Dec 24 17:34 login-filled.png
-rw-r--r--  1 cd1  staff   98327 Dec 24 17:35 overview.png
-rw-r--r--  1 cd1  staff   46870 Dec 24 17:35 policies.png
-rw-r--r--  1 cd1  staff  322084 Dec 24 17:34 welcome.png
```

Get and build the diff-image project.

```bash
go get -u github.com/murooka/go-diff-image
```

Then, you can use the diff-image tool to compare different snapshot iterations.

```bash
diff-image scale1/policies.png offline/policies.png
```
