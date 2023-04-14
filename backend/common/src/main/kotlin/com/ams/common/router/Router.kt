package com.ams.backend.router

import com.ams.backend.admin.handler.AdminHandler
import org.springframework.http.MediaType
import org.springframework.web.reactive.function.server.coRouter

fun router(adminHandler: AdminHandler) = coRouter {
    (accept(MediaType.APPLICATION_JSON)).nest {
        "/admin".nest {
            POST("/register", adminHandler::register)
            POST("/login", adminHandler::login)
            GET("/token", adminHandler::tokenTest)
        }
    }
}
