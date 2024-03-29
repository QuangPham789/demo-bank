package db

// func TestStore_TransferTx(t *testing.T) {
// 	store := NewStore(testDB)
//
// 	acc1 := createRandomAccount(t)
// 	acc2 := createRandomAccount(t)
//
// 	n := 5
// 	amount := int64(10)
//
// 	errs := make(chan error)
// 	results := make(chan TransferTxResult)
//
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			result, err := store.TransferTx(context.Background(),
// 				TransferTxParams{
// 					FromAccountID: acc1.ID,
// 					ToAccountID:   acc2.ID,
// 					Amount:        amount,
// 				})
// 			errs <- err
// 			results <- result
// 		}()
// 	}
// 	existed := make(map[int]bool)
// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)
// 		result := <-results
// 		require.NotEmpty(t, result)
//
// 		transfer := result.Transfer
// 		require.NotEmpty(t, transfer)
// 		require.Equal(t, acc1.ID, transfer.FromAccountID)
// 		require.Equal(t, acc2.ID, transfer.ToAccountID)
// 		require.Equal(t, amount, transfer.Amount)
// 		require.NotZero(t, transfer.ID)
// 		require.NotZero(t, transfer.CreatedAt)
//
// 		// check entries
// 		fromEntry := result.FromEntry
// 		require.NotEmpty(t, fromEntry)
// 		require.Equal(t, acc1.ID, fromEntry.AccountID)
// 		require.Equal(t, -amount, fromEntry.Amount)
// 		require.NotZero(t, fromEntry.ID)
// 		require.NotZero(t, fromEntry.CreatedAt)
//
// 		_, err = store.GetEntry(context.Background(), fromEntry.ID)
// 		require.NoError(t, err)
//
// 		toEntry := result.ToEntry
// 		require.NotEmpty(t, toEntry)
// 		require.Equal(t, acc2.ID, toEntry.AccountID)
// 		require.Equal(t, amount, toEntry.Amount)
// 		require.NotZero(t, toEntry.ID)
// 		require.NotZero(t, toEntry.CreatedAt)
//
// 		_, err = store.GetEntry(context.Background(), toEntry.ID)
// 		require.NoError(t, err)
//
// 		// check accounts
// 		fromAccount := result.FromAccount
// 		require.NotEmpty(t, fromAccount)
// 		require.Equal(t, acc1.ID, fromAccount.ID)
//
// 		toAccount := result.ToAccount
// 		require.NotEmpty(t, toAccount)
// 		require.Equal(t, acc2.ID, toAccount.ID)
// 		// check balances
// 		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
//
// 		diff1 := acc1.Balance - fromAccount.Balance
// 		diff2 := toAccount.Balance - acc2.Balance
// 		require.Equal(t, diff1, diff2)
// 		require.True(t, diff1 > 0)
// 		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount
//
// 		k := int(diff1 / amount)
// 		require.True(t, k >= 1 && k <= n)
// 		require.NotContains(t, existed, k)
// 		existed[k] = true
// 	}
// }
