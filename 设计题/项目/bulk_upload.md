
bulk upload:
目标：
1. 保证服务器稳定性，减少压力
2. 数据合格校验，记录进程（执行行数和结果），记录最终结果，如果失败生成失败文件对应一行的原因
3. 用户可以上传多个task

实现思路：
1. csv代替xlsx
2. 根据上传类型决定文件处理顺序，VGI/VML
3. 每一个字段通过对应表的map对应权限

具体实现:
1. 用户上传 -> CSV / zip? -> 大小和重复检验 -> 上传S3 得到url -> create task -> 发SQS -> 改task状态为queue
2. SQS: zip Handler -> file Handler -> create / update merchant -> create / update item -> generate error file handler
3. 
  3.1 zip handler: unzip -> uplaod all files to s3 -> update timeline in auditlog and progress in redis -> produce datafile msg
  3.2 data file: download/parse file -> generate merchant/item msg -> update timeline in auditlog and progress in redis
  3.3 merchant handler: validate/create/update merchant
  3.4 item handler: validate/create/update item
  3.5 error report generator: generate error file -> upload s3 -> update url to DB
  3.6 report collector(时间段执行一次): SETNX task validation collector key -> check task progress by validation taskIDs -> update status to DB -> delete task validation collector key

存储:
1. auditlog
2. DB
  2.1 task
3. redis: 
  3.0 SETNX task validation collector key
  3.1 validating task IDs => taskProgress:taskID - object(status, totalSheetCount, sheetName2TotalCount, sheetName2SuccessCount, merchantDocID2ID, sheetName2ErrorRowIndex2Msg)
  3.2 taskProgress:taskID:sheetName:name

难点:
1. 当上传达大文件时timeout问题？ 用grpc的客户端流RPC来解决
2. sqs重复消费？ merchant/item 由FIFO来保证只消费一次
3. 


怎么上线
