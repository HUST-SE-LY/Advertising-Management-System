package com.ams.backend.security

import com.auth0.jwt.JWT
import com.auth0.jwt.algorithms.Algorithm
import com.auth0.jwt.interfaces.DecodedJWT
import org.springframework.beans.factory.annotation.Value
import org.springframework.security.core.GrantedAuthority
import org.springframework.security.core.authority.SimpleGrantedAuthority
import org.springframework.security.core.userdetails.User
import org.springframework.stereotype.Service
import java.util.*
import kotlin.time.Duration
import kotlin.time.Duration.Companion.days
import kotlin.time.Duration.Companion.hours
import kotlin.time.DurationUnit

@Service
class JwtService(
    @Value("\${app.secret}") val secret: String,
    @Value("\${app.refresh}") val refresh: String
) {

    fun accessToken(username: String, expirationDuration: Duration, roles: Array<String>): String {
        return generate(username, expirationDuration, roles, secret)
    }

    fun decodeAccessToken(accessToken: String): DecodedJWT {
        return decode(secret, accessToken)
    }

    fun refreshToken(username: String, expirationDuration: Duration, roles: Array<String>): String {
        return generate(username, expirationDuration, roles, refresh)
    }

    fun decodeRefreshToken(refreshToken: String): DecodedJWT {
        return decode(refresh, refreshToken)
    }

    fun getRoles(decodedJWT: DecodedJWT) = decodedJWT.getClaim("role").asList(String::class.java)
        .map { SimpleGrantedAuthority(it) }

    private fun generate(
        username: String,
        expirationInMillis: Duration,
        roles: Array<String>,
        signature: String
    ): String {
        return SecurityConstants.TOKEN_PREFIX + JWT.create()
            .withSubject(username)
            .withExpiresAt(Date(System.currentTimeMillis() + expirationInMillis.toLong(DurationUnit.MILLISECONDS)))

            .withArrayClaim("role", roles)
            .sign(Algorithm.HMAC512(signature.toByteArray()))
    }

    private fun decode(signature: String, token: String): DecodedJWT {
        return JWT.require(Algorithm.HMAC512(signature.toByteArray()))
            .build()
            .verify(token.replace(SecurityConstants.TOKEN_PREFIX, ""))
    }


    fun generateTokenFromUser(user: User): Pair<String, String> {
        val (accessToken, refreshToken) = run {
            val roles = user.authorities.map { it.authority }.toTypedArray()
            accessToken(user.username, ACCESS_TOKEN_EXPIRATION, roles) to refreshToken(
                user.username,
                REFRESH_TOKEN_EXPIRATION, roles
            )
        }
        return accessToken to refreshToken
    }

    fun decodeAccessTokenToUser(accessToken: String): User {
        val decodedJWT = decodeAccessToken(accessToken)
        val username = decodedJWT.subject
        val roles = decodedJWT.getClaim("role").asList(String::class.java).map {
            GrantedAuthority { it }
        }
        return User(username, "", roles)
    }

    companion object {
        private val ACCESS_TOKEN_EXPIRATION = 2.hours
        private val REFRESH_TOKEN_EXPIRATION = 7.days
    }
}