namespace go base

struct User {
  1: i64 UserId,
  2: i64 LikeNum,
  3: optional i64 Birthday,
  4: i8 Gender,
  5: i8 Type,
  6: i8 Enable,
  7: i32 CommentNum,
  8: i32 ArticleNum,
  9: string Status,
  10: string Company,
  11: string WxOpenid,
  12: string RealName,
  13: string NickName,
  14: string UserName,
  15: string Password,
  16: string Mobile,
  17: string Email,
  18: string Blog,
  19: string Avatar,
  20: string Description,
  21: string Location,
  22: string School
}


struct UserInfo {
  1: i64 UserId,
  2: i64 LikeNum,
  3: optional i64 Birthday,
  4: i8 Gender,
  5: i8 Type,
  6: i8 Enable,
  7: i32 CommentNum,
  8: i32 ArticleNum,
  9: string Status,
  10: string Company,
  11: string RealName,
  12: string NickName,
  13: string UserName,
  14: string Mobile,
  15: string Email,
  16: string Blog,
  17: string Avatar,
  18: string Description,
  19: string Location,
  20: string School
}