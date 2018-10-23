










SELECT
  A.`id`,
  A.`course_name`,

  COUNT(B.`score`) "总人数",

  COUNT((B.`score` >= 0 AND B.`score` <= 60) OR NULL)                               "[0,60] 人数",
  COUNT((B.`score` > 60 AND B.`score` <= 70) OR NULL)                               "(60,70] 人数",
  COUNT((B.`score` > 70 AND B.`score` <= 85) OR NULL)                               "(70,85] 人数",
  COUNT((B.`score` > 85 AND B.`score` <= 100) OR NULL)                              "(85,100] 人数",


  ((COUNT((B.`score` >= 0 AND B.`score` <= 60) OR NULL) * 100) / COUNT(B.`score`))  "[0,60] 人数所占比例",
  ((COUNT((B.`score` > 60 AND B.`score` <= 70) OR NULL) * 100) / COUNT(B.`score`))  "(60,70] 人数所占比例",
  ((COUNT((B.`score` > 70 AND B.`score` <= 85) OR NULL) * 100) / COUNT(B.`score`))  "(70,85] 人数所占比例",
  ((COUNT((B.`score` > 85 AND B.`score` <= 100) OR NULL) * 100) / COUNT(B.`score`)) "(85,100] 人数所占比例"

FROM `sql`.`course` A LEFT JOIN `sql`.`score` B ON A.`id` = B.`course_id`
GROUP BY A.`id`;





