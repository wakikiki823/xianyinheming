'use strict';
exports.main = async (event, context) => {
  // 获取用户ID
  const { uid } = await uniID.checkToken(event.uniIdToken);
  
  // 查询帖子作者
  const post = await db.collection('posts').doc(event.postId).get();
  
  // 验证权限
  if (post.data.userId !== uid) {
    return { code: 403, message: '无删除权限' };
  }
  
  // 执行删除
  await db.collection('posts').doc(event.postId).remove();
  
  return { code: 200, message: '删除成功' };
}; 