package repository

import (
	"ru-library-api/entity"
)

func (r *sierraRepoDB) PatronId(fines *[]entity.PatronRepo, id string) error {
	err := r.postgres_db.Select(fines, `SELECT p.id,p.barcode,pf.last_name,p.ptype_code,pp.description
				FROM sierra_view.patron_view p 
				inner join sierra_view.patron_record_fullname pf on p.id = pf.patron_record_id
				inner join sierra_view.ptype_property_name pp on p.ptype_code = pp.ptype_id and pp.iii_language_id = 33
				where p.barcode = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
