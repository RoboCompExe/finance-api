func Transfer(c *gin.Context) {
	var req struct {
		From   string
		To     string
		Amount int64
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer tx.Rollback()

	var balance int64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id=$1 FOR UPDATE", req.From).Scan(&balance)

	if err != nil || balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id=$2", req.Amount, req.From)
	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id=$2", req.Amount, req.To)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
