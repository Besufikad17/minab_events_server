package constants

var CreateEvent string = "mutation CreateEvent(   $title: String!,   $description: String!,  $user_id: Int!,   $category_id: Int!,   $location_id: Int!,   $image: String!,   $enterance_fee: float8!,   $start_date: date!,   $end_date: date!, ) { 	insert_events_one(     object: {       title: $title,        description: $description,        user_id: $user_id,  category_id: $category_id,       location_id: $location_id,       image: $image,       enterance_fee: $enterance_fee,       start_date: $start_date,       end_date: $end_date,     }   ) {     id     title     description     image     enterance_fee     start_date     end_date     location {       city       venue     }     category {       name     }     tags {       name     }   }   }"

var CreateLocation string = "mutation CreateLocation($city: String!, $venue: String!) {   insert_locations_one(object: { city: $city, venue: $venue }) {     id     city     venue   } }"

var CreateTag string = "mutation CreateTag(  	$event_id: Int!,   $name: String! ) {   insert_tags_one(object: {event_id: $event_id, name: $name}) {       id     	name   } }"

var GetUserById string = "query GetUserById(   $id: Int!  ) {   users(     where:{       id: {         _eq: $id       }     }   ) {     first_name     last_name     email   } }"

var Register string = "mutation Register(   $first_name: String!,    $last_name: String!,    $email: String!,    $phone_number: String!,    $password: String! ) {   insert_users_one(object: {     first_name: $first_name,      last_name: $last_name,      email: $email,      phone_number: $phone_number,      password: $password }   ) {   	id     first_name     last_name     email     phone_number   } }"

var SearchUser string = "query SearchUser($login_text: String!) {   users(where: {     _or: [       {email: {_eq: $login_text}},       {phone_number: {_eq: $login_text}}     ]   }) {     id     first_name     last_name     email     phone_number     password   } }"
