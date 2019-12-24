# gohello

## Summary

* use this main package to call sub package generate screen shots - to detect UI outage

---

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
```
