package com.ams.backend.utils

import com.ams.backend.model.Role
import com.ams.backend.security.SecurityConstants
import io.jsonwebtoken.Claims
import io.jsonwebtoken.Jwts
import io.jsonwebtoken.SignatureAlgorithm
import io.jsonwebtoken.io.Serializer
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken
import org.springframework.security.core.Authentication
import org.springframework.security.core.authority.SimpleGrantedAuthority
import java.security.Key
import java.util.Date
import javax.crypto.spec.SecretKeySpec
import kotlin.io.encoding.Base64

object JwtTokenUtil {

    // The length of JWT_SECRET should be long
    // from https://www.allkeysgenerator.com/
    private const val JWT_SECRET =
        "fcb3af795ad147698cfc6987a4e93ba931f5c91d50ad49478787b1e97b4ab6f5ed282cd9c6148948f86ad06b792cfd26e005db8ab79466a8b17556e643f93e8ca7308ed42ec4974907269ee01e592f8"


    private val signatureAlgorithm = SignatureAlgorithm.HS256

    fun generateToken(username: String,/* roles: List<String>,*/ isRemember: Boolean): String {
        val claims = Jwts.claims().setSubject(username).apply {
//            set(SecurityConstants.TOKEN_ROLE_CLAIM, roles)
        }
        return Jwts.builder().apply {
            setHeaderParam("typ", SecurityConstants.TOKEN_TYPE)
            setSubject(username)
            setClaims(claims)
            setAudience(SecurityConstants.TOKEN_AUDIENCE)
            setIssuer(SecurityConstants.TOKEN_ISSUER)
            setIssuedAt(Date())
            setExpiration(generateExpirationDate(isRemember))
        }.signWith(getSignKey()).compact()
    }

    fun extractUsername(token: String): String {
        return extractClaim(token, Claims::getSubject)
    }

    fun validateToken(token: String): Boolean {
        return try {
            getTokenBody(token)
            true
        } catch (e: Throwable) {
            false
        }
    }

    fun isTokenExpired(token: String): Boolean {
        val expiredDate = getExpiredDateFromToken(token)
        return expiredDate.before(Date())
    }

    private fun extractClaim(token: String, claimsResolver: Claims.() -> String): String {
        val claims = getTokenBody(token)
        return claims.claimsResolver()
    }

    private fun getExpiredDateFromToken(token: String): Date {
        val claims = getTokenBody(token)
        return claims.expiration
    }

    @Suppress("UNCHECKED_CAST")
    fun getAuthentication(token: String): Authentication {
        val claims = getTokenBody(token)
        val authorities = try {
            val roles = claims[SecurityConstants.TOKEN_ROLE_CLAIM] as List<String>
            roles.map(::SimpleGrantedAuthority)
        } catch (e: Exception) {
            listOf(SimpleGrantedAuthority(Role.User.description))
        }
        val username = claims.subject
        return UsernamePasswordAuthenticationToken(username, token, authorities)
    }

    private fun getTokenBody(token: String): Claims {
        return Jwts.parserBuilder().setSigningKey(getSignKey()).build().parseClaimsJws(token).body
    }

    private fun generateExpirationDate(isRemember: Boolean): Date {
        return Date(System.currentTimeMillis() + if (isRemember) SecurityConstants.EXPIRATION_REMEMBER_TIME else SecurityConstants.EXPIRATION_TIME)
    }


    private fun getSignKey(): Key {
        // Sign JWT with the key
        val apiKeySecretBytes = Base64.encode(JWT_SECRET.encodeToByteArray()).encodeToByteArray()
        return SecretKeySpec(apiKeySecretBytes, signatureAlgorithm.jcaName)
    }
}