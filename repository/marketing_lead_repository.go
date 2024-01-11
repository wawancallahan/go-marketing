package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadRepository interface {
	FindAll(tx *gorm.DB) (*[]model.MarketingLead, error)
	Create(tx *gorm.DB, marketingLead *model.MarketingLead) error
	Find(tx *gorm.DB, id string) (*model.MarketingLead, error)
	Update(tx *gorm.DB, marketingLead *model.MarketingLead) error
	Delete(tx *gorm.DB, id string) error
	UpdatePartiallyRegistered(tx *gorm.DB) error
	UpdateRegistered(tx *gorm.DB) error
	UpdateDuplicateEntry(tx *gorm.DB) error
	UpdateNonDuplicateEntry(tx *gorm.DB) error
}

type MarketingLeadRepositoryImpl struct{}

func NewMarketingLeadRepository() *MarketingLeadRepositoryImpl {
	return &MarketingLeadRepositoryImpl{}
}

func (r *MarketingLeadRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.MarketingLead, error) {
	var marketingLead []model.MarketingLead

	err := tx.Find(&marketingLead).Error

	if err != nil {
		return nil, err
	}

	return &marketingLead, nil
}

func (r *MarketingLeadRepositoryImpl) Create(tx *gorm.DB, marketingLead *model.MarketingLead) error {
	err := tx.Create(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) Find(tx *gorm.DB, id string) (*model.MarketingLead, error) {
	var marketingLead model.MarketingLead

	err := tx.Take(&marketingLead, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &marketingLead, nil
}

func (r *MarketingLeadRepositoryImpl) Update(tx *gorm.DB, marketingLead *model.MarketingLead) error {
	err := tx.Updates(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var marketingLead model.MarketingLead

	err := tx.Where("id = ?", id).Delete(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) UpdatePartiallyRegistered(tx *gorm.DB) error {
	err := tx.Exec(`
		UPDATE pm_marketing.marketing_leads as ml SET activation_status = 'PARTIALLY_REGISTERED'
		WHERE (exists (
		SELECT 1 FROM public.pj_user as pu WHERE pu.email = ml.email
		) OR exists (
		SELECT 1 FROM public.pj_loan_details as pld WHERE pld.mobile_number = ml.phone_number
		)) AND ml.activation_status = 'NOT_REGISTERED' AND ml.status = 'FOLLOWUP'`,
	).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) UpdateRegistered(tx *gorm.DB) error {
	err := tx.Exec(`
		UPDATE pm_marketing.marketing_leads as ml SET activation_status = 'REGISTERED'
		WHERE (exists (
		SELECT 1 FROM public.pj_user as pu WHERE pu.email = ml.email
		) AND exists (
		SELECT 1 FROM public.pj_loan_details as pld WHERE pld.mobile_number = ml.phone_number
		)) AND (ml.activation_status = 'NOT_REGISTERED' OR ml.activation_status = 'PARTIALLY_REGISTERED') AND ml.status = 'FOLLOWUP'`,
	).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) UpdateDuplicateEntry(tx *gorm.DB) error {
	err := tx.Exec(`
		UPDATE pm_marketing.marketing_leads SET is_duplicate = true 
		WHERE is_duplicate = false AND
			id IN (
				SELECT 
					id 
				FROM (
					SELECT 
						id,
						ROW_NUMBER() OVER (
							PARTITION BY email
							ORDER BY created_at ASC) AS row_num
					FROM 
						pm_marketing.marketing_leads
					UNION
					SELECT 
						id,
						ROW_NUMBER() OVER (
							PARTITION BY phone_number
							ORDER BY created_at ASC) AS row_num
					FROM 
						pm_marketing.marketing_leads
			) t WHERE row_num > 1
		);`,
	).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) UpdateNonDuplicateEntry(tx *gorm.DB) error {
	err := tx.Exec(`
		UPDATE pm_marketing.marketing_leads ml SET is_duplicate = false
		WHERE is_duplicate = true AND
		exists ( 
			SELECT 
				phone_number, email, COUNT(email), COUNT(phone_number)
			FROM
				pm_marketing.marketing_leads
			GROUP BY 
				email, phone_number
			HAVING 
				COUNT(email) = 1 AND COUNT(phone_number) = 1
		) AND id IN (
			SELECT 
				id 
			FROM (
				SELECT 
					id,
					ROW_NUMBER() OVER (
						PARTITION BY email
						ORDER BY created_at ASC) AS row_num
				FROM 
					pm_marketing.marketing_leads
				UNION
				SELECT 
					id,
					ROW_NUMBER() OVER (
						PARTITION BY phone_number
						ORDER BY created_at ASC) AS row_num
				FROM 
					pm_marketing.marketing_leads
			) t WHERE row_num = 1
		)`,
	).Error

	if err != nil {
		return err
	}

	return nil
}
