const port = process.env.PORT || 3000
const ONLINE_API_URL = ``
const TIMESTAMP = 1618481602784
/**
 * @description  1. 参数名ASCII码从小到大排序（字典序）
 * @param {object} obj
 * @returns {object} 返回排序好的对象
 */
function sortByASCII(obj) {
  var arr = new Array()
  var num = 0
  for (var i in obj) {
    arr[num] = i
    num++
  }
  var sortArr = arr.sort()
  var sortObj = {}
  for (var i in sortArr) {
    sortObj[sortArr[i]] = obj[sortArr[i]]
  }
  return sortObj
}

/**
 * @description 2. 使用 URL 键值对的格式（即key1=value1&key2=value2）拼接成字符串
 * @param {string|number|boolean} param
 * @param {*} key
 * @param {*} encode
 * @returns {string} 返回 URL 键值对字符串
 */
function urlEncode(param, key, encode) {
  if (param == null) return ''
  var paramStr = ''
  var t = typeof param
  if (t == 'string' || t == 'number' || t == 'boolean') {
    // paramStr += '&' + key + '=' + ((encode==null||encode) ? encodeURIComponent(param) : param)
    paramStr += '&' + key + '=' + param
  } else {
    for (var i in param) {
      var k =
        key == null
          ? i
          : key + (param instanceof Array ? '[' + i + ']' : '.' + i)
      paramStr += urlEncode(param[i], k, encode)
    }
  }
  return paramStr
}

// 统一下单必填的数据（除了sign），以下数据为测试数据
let dataObject = {
  appId: '30487427', // 游戏上架后分配的 appId
  openId: 'TN_VDhvcnhkellJNFlWZkFtckZsK0Z5OXZSWFBCWVY0VUY0WVdFclV1Z3BrRHNlYkJwR0FxZElicW9tSjR5ditUQQ', // token
  timestamp: TIMESTAMP, // 建议作为常量存储时间戳，因为后续签名会对时间戳做校验，如：const TIMESTAMP = new Date().getTime()
  productName: '600钻石',
  productDesc: '600钻石',
  count: 1,
  price: 60000,
  currency: 'CNY',
  cpOrderId: '142',
  appVersion: '116',
  engineVersion: '1080',
  callBackUrl: `https://test.lynlzqy.com/uc_p25/PayCallbackAction/3/3/5` // 服务器接收平台返回数据的接口回调地址
}
// 获得第一步结果，URL 键值对，如: appId=&appVersion=1.0.0&callBackUrl=http://127.0.0.1:3000/payResult&count=1&cpOrderId=1&currency=CNY&engineVersion=1045&openId=TN_dmhkN3BNWWQ3T0RCNHdIUnlzc2JVc0ljRnZxUit5TWNhODJhWlM3MXpyeUN4ekdaZmhoTXVlU3FyZ0l0c05ieA&price=1&productDesc=testpay&productName=测试&timestamp=1592557604326
let dataString = urlEncode(sortByASCII(dataObject))

console.log(dataString)

appId=30487427&appVersion=116&callBackUrl=https://test.lynlzqy.com/uc_p25/PayCallbackAction/3/3/5&count=1&cpOrderId=143&currency=CNY&engineVersion=1080&openId=TN_VDhvcnhkellJNFlWZkFtckZsK0Z5OXZSWFBCWVY0VUY0WVdFclV1Z3BrRHNlYkJwR0FxZElicW9tSjR5ditUQQ&price=60000&productDesc=600钻石&productName=600钻石&timestamp=1618481602784
appId=30487427&appVersion=116&callBackUrl=https://test.lynlzqy.com/uc_p25/PayCallbackAction/3/3/5&count=1&cpOrderId=143&currency=CNY&engineVersion=1080&openId=TN_VDhvcnhkellJNFlWZkFtckZsK0Z5OXZSWFBCWVY0VUY0WVdFclV1Z3BrRHNlYkJwR0FxZElicW9tSjR5ditUQQ&price=60000&productDesc=600钻石&productName=600钻石&timestamp=1618481602784
