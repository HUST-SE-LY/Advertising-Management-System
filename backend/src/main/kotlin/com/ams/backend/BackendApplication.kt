package com.ams.backend

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.ApplicationContextInitializer
import org.springframework.context.support.GenericApplicationContext
import org.springframework.context.support.beans
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder

@SpringBootApplication
class BackendApplication


fun main(args: Array<String>) {
    runApplication<BackendApplication>(*args)
}
