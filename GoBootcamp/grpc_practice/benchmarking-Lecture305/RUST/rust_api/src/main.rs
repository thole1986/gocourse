// // USING ACTIX TO CREATE THE API
use actix_web::{post, web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
struct HelloRequest {
    name: String,
}

#[derive(Serialize)]
struct HelloResponse {
    message: String,
}

#[post("/v1/greet")]
async fn greet(req: web::Json<HelloRequest>) -> impl Responder {
    if req.name.len() < 5 || req.name.len() > 50 || !req.name.chars().all(char::is_alphabetic) {
        let response = serde_json::json!({
            "error": "Invalid request: name must be 5-50 characters long and contain only letters"
        });
        return HttpResponse::BadRequest().json(response);
    }

    let resp = HelloResponse {
        message: format!("Hello, {}", req.name),
    };

    HttpResponse::Ok().json(resp)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(greet)
    })
    .bind("127.0.0.1:3000")?
    .run()
    .await
}


// // // // USING ROCKET TO CREATE THE API, defaults to port 8000
// #[macro_use]
// extern crate rocket;

// use rocket::serde::{json::Json, Deserialize, Serialize};

// #[derive(Deserialize)]
// struct HelloRequest {
//     name: String,
// }

// #[derive(Serialize)]
// struct HelloResponse {
//     message: String,
// }

// #[post("/v1/greet", format = "json", data = "<req>")]
// fn greet(req: Json<HelloRequest>) -> Json<HelloResponse> {
//     if req.name.len() < 5 || req.name.len() > 50 || !req.name.chars().all(char::is_alphabetic) {
//         return Json(HelloResponse {
//             message: "Invalid request: name must be 5-50 characters long and contain only letters".to_string(),
//         });
//     }

//     Json(HelloResponse {
//         message: format!("Hello, {}", req.name),
//     })
// }

// #[launch]
// fn rocket() -> _ {
//     rocket::build().mount("/", routes![greet])
// }
