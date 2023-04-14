package com.ams.backend.configuation

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder

@Configuration
class GlobalConfiguration {

    @Bean
    fun bcryptPasswordEncoder() = BCryptPasswordEncoder()


}