export default {
  // 阿里云OSS配置示例
  oss: {
    region: 'oss-cn-hangzhou',
    accessKeyId: 'YOUR_ACCESS_KEY',
    accessKeySecret: 'YOUR_SECRET_KEY',
    bucket: 'your-bucket-name'
  },
  
  // 腾讯云COS配置示例
  cos: {
    Region: 'ap-shanghai',
    SecretId: 'YOUR_SECRET_ID',
    SecretKey: 'YOUR_SECRET_KEY',
    Bucket: 'your-bucket-name-1250000000'
  },
  
  // uniCloud配置示例
  uniCloud: {
    provider: 'aliyun', // 阿里云版
    spaceId: 'your-space-id',  // 服务空间ID
    clientSecret: {
      "accessKeyId": "your-access-key",
      "accessKeySecret": "your-secret-key"
    }
  },
  
  gcs: {
    bucketName: 'your-bucket',
    apiKey: 'AIzaSyABCD...',
    authDomain: 'your-project.firebaseapp.com',
    storageBucket: 'gs://your-bucket'
  }
} 