package service

import (
	"encoding/json"
	"fmt"
	"log"
	"ru-library-api/entity"
	"ru-library-api/errs"
	"time"
)

func (g *sierraService) PatronById(personId string) (*[]entity.Patron, error) {

	key := "fine::" + personId
	patronCache, err := g.redis_cache.Get(ctx, key).Result()
	patronrepos := []entity.PatronRepo{}
	patrons := []entity.Patron{}
	if err == nil {
		log.Println(err)
		err = json.Unmarshal([]byte(patronCache), &patrons)
		if err != nil {
			return nil, err
		}
		fmt.Println("cache-patrons")
		return &patrons, nil
	}

	fmt.Println("database-patrons")

	err = g.sierraRepo.PatronId(&patronrepos, personId)
	if err != nil {
		return &patrons, err
	}

	for _, c := range patronrepos {
		patrons = append(patrons, entity.Patron{
			Id:          c.Id,
			Barcode:     c.Barcode,
			Fullname:    c.Fullname,
			PtypeCode:   c.PtypeCode,
			Description: c.Description,
		})
	}

	if len(patrons) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลสมาชิก.")
	}

	patronsJSON, err := json.Marshal(&patrons)
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	redisCachegrade := time.Unix(timeNow.Add(time.Second*10).Unix(), 0)

	err = g.redis_cache.Set(ctx, key, patronsJSON, redisCachegrade.Sub(timeNow)).Err()
	if err != nil {
		return nil, err
	}

	return &patrons, nil
}
