namespace go demo

struct User {
    1:required i32 id,
    2:required string name,
    3:required string avatar,
    4:required string address,
    5:required string mobile,
}


struct UserList {
    1:required list<User> userList,
    2:required i32 page,
    3:required i32 limit,
}

// 重新定义类型名称，同c语言
typedef map<string, string> Data

// 定义响应体结构
struct Response {
    1:required i32 errcode,
    2:required string errmsg,
    3:required Data data,
}

// 定义服务接口，相当于go的interface
service Greeter {
    Response SayHello(
        1:required User user
    )

    Response GetUser(
        1:required i32 uid
    )
}