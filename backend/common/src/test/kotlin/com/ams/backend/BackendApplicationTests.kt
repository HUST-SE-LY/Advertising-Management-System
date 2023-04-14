package com.ams.backend

import com.ams.backend.security.JwtService
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest

@SpringBootTest
class BackendApplicationTests(@Autowired private val jwtService: JwtService) {

    @Test
    fun contextLoads() {

    }

}
