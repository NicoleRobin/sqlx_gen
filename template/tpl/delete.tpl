// Delete remove a row
func Delete(ctx context.Context, id int64) (int64, error) {
	e, err := models.GetDB()
	if err != nil {
		log.Error(ctx, "models.GetDB() failed", zap.Error(err))
		return 0, err
	}

	res, err := e.Bind(tableName).Delete(row).Where(db.Eq("id", id)).Exec(ctx)
	if err != nil {
		log.Error(ctx, "update db failed", zap.Error(err))
		return 0, err
	}

	return res.RowsAffected()
}
