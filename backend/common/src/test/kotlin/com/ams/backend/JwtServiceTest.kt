package com.ams.backend

import com.ams.backend.security.JwtService
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import kotlin.time.Duration.Companion.days

@SpringBootTest
class JwtServiceTest(@Autowired private val jwtService: JwtService) {

    @Test
    fun test() {
        val username = "fnfunfunc"
        val expiredDuration = 1.days
        val token = jwtService.accessToken(username = username, expirationDuration = expiredDuration, roles = arrayOf())
        val decodedJwt = jwtService.decodeAccessToken(token)
        assert(username == decodedJwt.subject)
    }
}