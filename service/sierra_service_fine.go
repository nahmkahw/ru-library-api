package service

import (
	"encoding/json"
	"fmt"
	"log"
	"ru-library-api/entity"
	"ru-library-api/errs"
	"time"
)

func (g *sierraService) FineById(personId string) (*[]entity.Fine, error) {

	key := "fine::" + personId
	fineCache, err := g.redis_cache.Get(ctx, key).Result()
	finerepos := []entity.FineRepo{}
	fines := []entity.Fine{}
	if err == nil {

		err = json.Unmarshal([]byte(fineCache), &fines)

		if err != nil {
			return nil, err
		}

		fmt.Println("cache-fines")

		return &fines, nil
	}

	fmt.Println("database-fines")

	err = g.sierraRepo.FineId(&finerepos, personId)

	if err != nil {
		log.Println(err.Error())
		return &fines, err
	}

	for _, c := range finerepos {
		fines = append(fines, entity.Fine{
			Id:            c.Id,
			Barcode:       c.Barcode,
			Amount:        c.Amount,
			InvoiceNum:    c.InvoiceNum,
			ItemCharge:    c.ItemCharge,
			ProcessingFee: c.ProcessingFee,
			BillingFee:    c.BillingFee,
			ChargeCode:    c.ChargeCode,
			PaidAmount:    c.PaidAmount,
			ItemId:        c.ItemId,
			LocationCode:  c.LocationCode,
			Title:         c.Title,
		})
	}

	if len(fines) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลค่าปรับ.")
	}

	finesJSON, err := json.Marshal(&fines)
	if err != nil {
		return nil, err
	}
	timeNow := time.Now()
	redisCachegrade := time.Unix(timeNow.Add(time.Second*10).Unix(), 0)
	err = g.redis_cache.Set(ctx, key, finesJSON, redisCachegrade.Sub(timeNow)).Err()

	if err != nil {
		return nil, err
	}

	return &fines, nil
}
