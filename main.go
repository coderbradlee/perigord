// Example main file for a native dapp, replace with application code
package main

import (
	//	"context"

	//	"github.com/polyswarm/perigord/migration"

	"context"
	"fmt"
	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"
	"log"
	"perigord/bindings"
	// "perigord/migrations"
)

func main() {
	// Run our migrations
	//migration.RunMigrations(context.Background())
	network.InitNetworks()
	nw, err := network.Dial("dev")
	if err != nil {
		log.Fatalln("could not connect to dev network: ", err)
	}
	// Run our migrations
	if err := migration.RunMigrations(context.Background(), nw, true); err != nil {
		log.Fatalln("error running migrations: ", err)
	}
	session, ok := contract.Session("MonsterCore").(*bindings.MonsterCoreSession)
	if !ok {
		log.Fatalln("error retrieving session")
	}
	// name, _ := session.Name()
	totalSupply, _ := session.TotalSupply()
	// symbol, _ := session.Symbol()
	// fmt.Printf("Let's spend some %s\n", name)
	// fmt.Printf("There are %d %s in total\n", totalSupply, symbol)
	fmt.Println("supply:", totalSupply)
}
