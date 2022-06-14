package app

import (
	"CryptoTest/internal/repositories"
	servicesb "CryptoTest/internal/services"
	"github.com/jasonlvhit/gocron"
	"log"
)

var kucoinService servicesb.KucoinService
var cbrService servicesb.CbrService

func SetupSchedulerJobs() {
	go func() {
		kucoinService = servicesb.NewKucoinService(&repositories.Repos)
		cbrService = servicesb.NewCbrService(&repositories.Repos)

		if err := gocron.Every(1).Day().From(gocron.NextTick()).Do(everyDay); err != nil {
			log.Println("Cant run UpdateSymbols cron")
		}

		if err := gocron.Every(10).Second().Do(everyTenSeconds, "btc-usdt"); err != nil {
			log.Println("Cant run UpdateSymbols cron")
		}

		<-gocron.Start()
	}()
}

func everyDay() {
	_, err := kucoinService.UpdateSymbols()
	if err != nil {
		log.Printf("Cannot update 'KuCoin' symbols: %s\n", err)
	}

	_, err = cbrService.UpdateSymbols()
	if err != nil {
		log.Printf("Cannot update 'CBR' symbols: %s\n", err)
	}

	_, err = cbrService.UpdateCourses()
	if err != nil {
		log.Printf("Cannot update 'CBR' courses: %s\n", err)
	}
}

func everyTenSeconds(pair string) {
	_, err := kucoinService.UpdateStats(pair)
	if err != nil {
		log.Printf("Cand update exchange rate for %s, err: %s\n", pair, err)
	}
	log.Printf("Success update exchange rate for %s\n", pair)
}
