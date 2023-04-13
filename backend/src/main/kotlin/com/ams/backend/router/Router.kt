package com.ams.backend.router

import com.ams.backend.admin.handler.AdminHandler
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.MediaType
import org.springframework.web.reactive.function.server.coRouter

@Configuration
class Router(private val adminHandler: AdminHandler) {

    @Bean
    fun route() = coRouter {
        (accept(MediaType.APPLICATION_JSON)).nest {
            "/admin".nest {
                POST("/register", adminHandler::register)
                POST("/login", adminHandler::login)
            }
        }
    }
}