你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Dgraph

Copyright 2020 Ardan Labs  
bill@ardanlabs.com

## Licensing

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

## About The Project

Please read the project wiki (Comming Soon)

https://github.com/ardanlabs/dgraph/wiki

## Learn More

**To learn about Corporate training events please contact Ardan Labs:**

William Kennedy  
ArdanLabs (www.ardanlabs.com)  
bill@ardanlabs.com  


Getting Started
https://developer.twitter.com/en/docs/basics/getting-started

Get an account
https://developer.twitter.com/en/docs/basics/developer-portal/overview

Apply for a developer account
https://developer.twitter.com/en/application/use-case

Create an App

curl -u ${DGRAPH_TWITTER_API_KEY}:${DGRAPH_TWITTER_SECRET_KEY} \
  --data 'grant_type=client_credentials' \
  'https://api.twitter.com/oauth2/token'

curl --request GET \
  --url https://api.twitter.com/1.1/friends/ids.json?screen_name=awefulBrown \
  --header "authorization: bearer ${TWITTER_TOKEN}" \
  --header "content-type: application/json"

curl --request GET \
  --url https://api.twitter.com/1.1/users/show.json?user_id=699263 \
  --header "authorization: bearer ${TWITTER_TOKEN}" \
  --header "content-type: application/json"

  curl --request GET \
  --url https://api.twitter.com/1.1/users/show.json?screen_name=awefulBrown \
  --header "authorization: bearer ${TWITTER_TOKEN}" \
  --header "content-type: application/json"
