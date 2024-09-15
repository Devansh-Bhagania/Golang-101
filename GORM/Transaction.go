err := db.Transaction(func(tx *gorm.DB) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}

	if err := tx.Create(&order).Error; err != nil {
		return err
	}

	return nil
})

if err != nil {
	log.Printf("Transaction failed: %v", err)
}
