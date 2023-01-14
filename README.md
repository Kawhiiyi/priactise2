# priactise2



  r.GET("/ping/u001", func(c *gin.Context) {
    c.JSON(http.*StatusOK*, gin.H{
     "receive_amount": "10",
    })
  })

  // /gin_demo/package_infos/u001
  r.GET("/gin_demo/package_infos/:user_id", func(c *gin.Context) {
    userId := c.Param("user_id")
    log.Printf("get userId from request %v", userId)

    var record RpReceiveRecord
    s := fmt.Sprintf("select * from rp_receive_record where user_id = '%s'", userId)
    err := db.QueryRow(s).Scan(&record.Id, &record.UserId, &record.GroupChatId,
     &record.RpId, &record.ReceiveAmount, &record.CreateTime, &record.NotifyTime)
    if err != nil {
     log.Printf("Can't find userId amount. userId : %v err : %v", userId, err)
     c.JSON(http.*StatusOK*, gin.H{
       "code":  "-1",
       "message": "receive not found",
     })
    } else {
     c.JSON(http.*StatusOK*, gin.H{
       "receive_amount": record.ReceiveAmount,
       "code":      "0",
       "message":    "success",
     })
    }
  })

  r.POST("/gin_demo/package_infos", func(c *gin.Context) {
    var packageInfo PackageInfo
    err := c.BindJSON(&packageInfo)
    if err != nil {
     log.Printf("Bind package info error %v", err)
     c.JSON(http.*StatusOK*, gin.H{
       "code":  "-2",
       "message": "bind error",
     })
     return
    }

    s := "insert into rp_receive_record (user_id, group_id, rp_id, receive_amount, create_time) values (?,?,?,?,?)"
    r, err := db.Exec(s, packageInfo.UserId, "insert_group001", "insert_rp001", packageInfo.ReceiveAmount, time.Now())
    if err != nil {
     log.Printf("insert data err: %v\n", err)
     c.JSON(http.*StatusOK*, gin.H{
       "code":  "-3",
       "message": "insert data error",
     })
     return
    } else {
     i, _ := r.LastInsertId()
     log.Printf("i: %v\n", i)
     c.JSON(http.*StatusOK*, gin.H{
       "code":    "0",
       "message":  "success",
       "primary_id": i,
     })
    }
  })


  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

 