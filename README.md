# **API Doc**

## **Go DB Server** 
- https://your_dns



## **Gmail Sender**
### <u>Send ( POST : /gmail/send )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body
```
  {
    recipient: { recipient: string },
    msg: { msg: string }
  }
```


## **Memo**
### <u>Create ( POST : /memo/create?user_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body
```
  {
    content: { content : string }
  }
```
### <u>Update ( PUT : /memo/update?memo_id={ memo_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body
```
  {
    content: { content : string }
  }
```
### <u>Get ( GET : /memo/get?memo_id={ memo_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response
```
  {
    "memo_id": { user_id : uint64 },
    "content": { content : string },
    "user_id": { user_id : uint64 },
    "has_read": { has_read : uint8 }
  }
```
### <u>Get Memos By User ID ( GET : /memo/get_memos_by_user_id?user_id_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response
```
  [
    {
      "memo_id": { user_id : uint64 },
      "content": { content : string },
      "user_id": { user_id : uint64 },
      "has_read": { has_read : uint8 }
    }, ...
  ]
```
### <u>Delete ( DELETE : /memo/delete?memo_id={ memo_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```

## **Schedule**
### <u>Create ( POST : /schedule/create?user_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body
```
  {
    title: { title : string },
    note: { note : string },
    start_time: { start_time : string // yyyy-mm-dd hh:mm },
    end_time: { end_time : string // yyyy-mm-dd hh:mm }
  }
```

### <u>Update ( PUT : /schedule/update?schedule_id={ schedule_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body
```
  {
    title: { title : string },
    note: { note : string },
    start_time: { start_time : string // yyyy-mm-dd hh:mm },
    end_time: { end_time : string // yyyy-mm-dd hh:mm }
  }
```

### <u>Get An Event ( GET : /schedule/get?schedule_id={ schedule_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```

- Response
```
  {
    title: { title : string },
    note: { note : string },
    start_time: { start_time : string // yyyy-mm-dd hh:mm },
    end_time: { end_time : string // yyyy-mm-dd hh:mm },
    user_id: { user_id : uint64 }
  }
```

### <u>Get Events In A Day ( GET : /schedule/get_day?user_id={ user_id : uint64 }&day={ day: string // yyyy-mm-dd } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```

- Response
```
[
  {
    title: { title : string },
    note: { note : string },
    start_time: { start_time : string // yyyy-mm-dd hh:mm },
    end_time: { end_time : string // yyyy-mm-dd hh:mm },
    user_id: { user_id : uint64 },
    created_time: { created_time : string },
    updated_time: { updated_time : string }
  }, ...
]
```

### <u>Get Events In A Month ( GET : /schedule/get_month?user_id={ user_id : uint64 }&month={ month: string // yyyymm } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```

- Response
```
[
  {
    title: { title : string },
    note: { note : string },
    start_time: { start_time : string // yyyy-mm-dd hh:mm },
    end_time: { end_time : string // yyyy-mm-dd hh:mm },
    user_id: { user_id : uint64 },
    created_time: { created_time : string },
    updated_time: { updated_time : string }
  }, ...
]
```

### <u>Delete ( DELETE : /schedule/delete?schedule_id={ schedule_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```

## **User**
### <u>Create ( POST : /user/create )</u>
- Body
```
  {
    "user_account": { user_account : string },
    "user_password": { user_password : string },
    "first_name": { first_name : string(optional) },
    "last_name": { last_name : string(optional) },
    "gender": { gender : uint8(optional) }, // Default: 0, Male: 1, Female: 2, Diver: 3
    "birthday": { birthday : string(optional) // yyyy-mm-dd }, Default: 1000-01-01
    "email": { email : string },
    "phone": { phone : string(optional) },
    "user_profile": { user_profile : string(optional) },
    "role": { role : uint8 } 
  }
 ```

### <u>Update ( PUT : /user/update?user_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body 
```
  {
    "first_name": { first_name : string(optional) },
    "last_name": { last_name : string(optional) },
    "gender": { gender : uint8(optional) },
    "birthday": { birthday : string(optional) // yyyy-mm-dd },
    "email": { email : string(optional) },
    "phone": { phone : string(optional) },
    "user_profile": { user_profile : string(optional) }
  }
```

### <u>Upload Profile ( POST : /user/upload_user_profile?user_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Body 
```
  {
    "file": { file : multipart.File }
  }
```

### <u>Get ( GET : /user/get?user_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response
```
  {
    "user_id": { user_id : uint64 },
    "user_account": { user_account : string },
    "first_name": { first_name : string },
    "last_name": { last_name : string },
    "gender": { gender : uint8 },
    "birthday": { birthday : string // yyyy-mm-dd },
    "email": { email : string },
    "phone": { phone : string },
    "user_profile": { user_profile : string }
  }
```

### <u>Get By Account ( GET : /user/get_by_account?user_account={ user_account : string } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response
```
  {
    "user_id": { user_id : uint64 },
    "user_account": { user_account : string },
    "first_name": { first_name : string },
    "last_name": { last_name : string },
    "gender": { gender : uint8 },
    "user_profile": { user_profile : string }
  }
```

### <u>Get All ( GET : /user/get_all )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response
```
[
  {
    "user_id": { user_id : uint64 },
    "user_account": { user_account : string },
    "first_name": { first_name : string },
    "last_name": { last_name : string },
    "gender": { gender : uint8 },
    "birthday": { birthday : string // yyyy-mm-dd },
    "email": { email : string },
    "phone": { phone : string },
    "user_profile": { user_profile : string }
  }, ...
]
```

### <u>Get Specific Roles ( GET : /user/get_specific_roles?role=${ role: uint8 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response
```
[
  {
    "user_id": { user_id : uint64 },
    "user_account": { user_account : string },
    "first_name": { first_name : string },
    "last_name": { last_name : string },
    "gender": { gender : uint8 },
    "birthday": { birthday : string // yyyy-mm-dd },
    "email": { email : string },
    "phone": { phone : string },
    "user_profile": { user_profile : string }
  }, ...
]
```

### <u>Delete ( DELETE : /user/delete?user_id={ user_id : uint64 } )</u>
- Authorization 
```
Basic Base64Encode(login_token: string)
````
### <u>Login ( POST : /user/login )</u>
- Authorization 
```
Basic Base64Encode{{ user_account: string }:{ Base64Encode(user_password: string) }}
````
- Body
```
{
  "role": {role: uint8} 
}
```
- Response 
```
{
    login_token: {login_token: string},
    user_id: {user_id: uint64}
}
```

### <u>Get Login Token ( GET : /user/get_login_token?user_id={ user_id : uint64 } )</u>
- Authorization
```
Basic Base64Encode(login_token: string)
```
- Response 
```
token: string
```