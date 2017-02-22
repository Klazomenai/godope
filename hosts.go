package main

import "github.com/lextoumbourou/goodhosts"

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func AddEntry(entries [][]string) {
  hosts, err := goodhosts.NewHosts()
  check(err)

  for i := range entries {
    // Note that nothing will be added to the hosts file until ``hosts.Flush`` is called.
    hosts.Add(entries[i][1], entries[i][0])
  }
  if err := hosts.Flush(); err != nil {
    panic(err)
  }
}
