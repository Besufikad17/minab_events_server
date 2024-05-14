package constants

var Register string = "mutation Register(   $first_name: String!,    $last_name: String!,    $email: String!,    $phone_number: String!,    $password: String! ) {   insert_users_one(object: {     first_name: $first_name,      last_name: $last_name,      email: $email,      phone_number: $phone_number,      password: $password }   ) {   	id     first_name     last_name     email     phone_number   } }"

var SearchUser string = "query SearchUser($login_text: String!) {   users(where: {     _or: [       {email: {_eq: $login_text}},       {phone_number: {_eq: $login_text}}     ]   }) {     id     first_name     last_name     email     phone_number     password   } }"
