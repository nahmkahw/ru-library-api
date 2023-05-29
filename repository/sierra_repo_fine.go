package repository

import (
	"fmt"
	"ru-library-api/entity"
)

func (r *sierraRepoDB) FineId(fines *[]entity.FineRepo, personId string) error {

	err := r.postgres_db.Select(fines, `SELECT p.id,p.barcode,p.owed_amt,
				f.invoice_num,f.item_charge_amt,f.processing_fee_amt,
				f.billing_fee_amt,f.charge_code,f.paid_amt,i.id item_id,i.location_code,b.title
				FROM sierra_view.patron_view p 
				inner join sierra_view.fine f on p.id =  f.patron_record_id
				inner join sierra_view.item_record i on i.id =  f.item_record_metadata_id
				inner join sierra_view.bib_record_item_record_link l on l.item_record_id = i.id
				inner join sierra_view.bib_view b on b.id = l.bib_record_id
				where p.id = $1`, personId)
	if err != nil {
		fmt.Println("Error" + err.Error())
		return err
	}

	return nil
}
