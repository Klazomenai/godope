package main

import (
  "flag"
  "log"
  "os"

  "github.com/digitalocean/godo"
  "golang.org/x/oauth2"
)

var appVersion string
var err error

func main() {
  version := flag.Bool("version", false, "Print the version and exit.")
  flag.Parse()
  if *version {
    log.Printf(appVersion)
    os.Exit(0)
  }

  accessToken := os.Getenv("DO_KEY")
  if accessToken == "" {
    log.Fatal("Usage: DO_KEY environment variable must be set.")
  }

  peerTag := os.Getenv("DO_TAG")

  // setup dependencies
  oauthClient := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
  apiClient := godo.NewClient(oauthClient)

  // collect list of all droplets
  var drops []godo.Droplet
  if peerTag != "" {
    drops, err = DropletListTags(apiClient.Droplets, peerTag)
  } else {
    drops, err = DropletList(apiClient.Droplets)
  }
  failIfErr(err)

  // sort out relevant host entries
  privatePeers := SortDroplets(drops)

  // append host to /etc/hosts
  AddEntry(privatePeers)
}

func failIfErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
