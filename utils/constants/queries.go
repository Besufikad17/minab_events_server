package constants

var CreateEvent string = "mutation CreateEvent($title: String!, $description: String!, $user_id: Int!, $category_id: Int!, $location_id: Int!, $thumbnail: String!, $start_date: date!, $end_date: date!) { insert_events_one(object: {title: $title, description: $description, user_id: $user_id, category_id: $category_id, location_id: $location_id, thumbnail: $thumbnail, start_date: $start_date, end_date: $end_date}) { id title description thumbnail start_date end_date location { city venue } category { name } tags { name } } }"

var AddImages string = "mutation AddImages(   $images: [images_insert_input!]! ) {   insert_images(     objects: $images   ) {     affected_rows   } }"

var AddTicket string = "mutation AddTicket( 	$event_id: Int!,   $ticket_type: String!,   $description: String!,   $price: float8! ) {   insert_tickets_one(     object: {       event_id: $event_id,       ticket_type: $ticket_type,       description: $description,       price: $price     }   ) {     id   } }"

var CreateImage string = "mutation CreateImage(  	$event_id: Int!,   $url: String! ) {   insert_images_one(     object: {     	event_id: $event_id,       url: $url     }   ) {     id     url   } }"

var CreateLocation string = "mutation CreateLocation($city: String!, $venue: String!,  $lat: float8!, $lng: float8!) {   insert_locations_one(object: { city: $city, venue: $venue, latitude: $lat, longtiude: $lng }) {     id     city     venue   } }"

var CreatePayment string = "mutation CreatePayment(   $user_id: Int!,   $ticket_id: Int!,   $amount: float8!,   $status: String!,   $reference: String! ) {   insert_payments_one(     object: {       user_id: $user_id,       ticket_id: $ticket_id,       amount: $amount,       status: $status,       reference: $reference     }   ) {     id   } }"

var CreateTag string = "mutation CreateTag(  	$event_id: Int!,   $name: String! ) {   insert_tags_one(object: {event_id: $event_id, name: $name}) {       id     	name   } }"

var GetTicketById string = "query GetTicketById(   $id: Int! ) {   tickets(     where: {       id: {         _eq: $id       }     }   ) {     price 	event_id   ticket_type } }"

var GetEvent string = "query GetEvent($id: Int!) { events(where: {id: {_eq: $id}}) { title description start_date end_date location { city venue latitude longtiude }}}"

var GetEventById string = "query GetEventById ($id: Int!) {   events(where:{ id: { _eq: $id } }) {     id     user_id     title     description     category {         id         name     }     thumbnail     images {         url     }     start_date     end_date     location {         id         city         venue         latitude         longtiude         full_location     }     tags {         name     }     bookmarks {         user_id     }     tickets {         id         ticket_type         description         price     }     reservations {       user_id       ticket_id       status     }   } }"

var GetUserById string = "query GetUserById(   $id: Int!  ) {   users(     where:{       id: {         _eq: $id       }     }   ) {     first_name     last_name     email    phone_number   } }"

var Register string = "mutation Register(   $first_name: String!,    $last_name: String!,    $email: String!,    $phone_number: String!,    $password: String! ) {   insert_users_one(object: {     first_name: $first_name,      last_name: $last_name,      email: $email,      phone_number: $phone_number,      password: $password }   ) {   	id     first_name     last_name     email     phone_number   } }"

var ReserveEvent string = "mutation ReserveEvent(     $user_id: Int!,     $event_id: Int!,     $ticket_id: Int!,     $status: String!, ) {     insert_reservations_one(         object: {             user_id: $user_id,             event_id: $event_id,             ticket_id: $ticket_id,             status: $status         }     ) {        id     } }"

var SearchUser string = "query SearchUser($login_text: String!) {   users(where: {     _or: [       {email: {_eq: $login_text}},       {phone_number: {_eq: $login_text}}     ]   }) {     id     first_name     last_name     email     phone_number     password   } }"
