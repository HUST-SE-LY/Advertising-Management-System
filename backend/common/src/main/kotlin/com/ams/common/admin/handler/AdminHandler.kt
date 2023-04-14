package com.ams.backend.admin.handler

import com.ams.backend.admin.dto.AdminDTO
import com.ams.backend.admin.dto.AdminRegisterDTO
import com.ams.backend.admin.repository.AdminRepository
import com.ams.backend.admin.dto.AdminLoginDTO
import com.ams.backend.admin.model.Admin
import com.ams.backend.security.JwtService
import com.auth0.jwt.exceptions.JWTVerificationException
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.*

@Component
class AdminHandler(
    private val adminRepository: AdminRepository,
    private val bCryptPasswordEncoder: BCryptPasswordEncoder,
    private val jwtService: JwtService
) {


    suspend fun register(request: ServerRequest): ServerResponse {
        val body = request.awaitBody<String>()
        val adminRegisterDTO = Json.decodeFromString<AdminRegisterDTO>(string = body)
        if (adminRepository.findAdminByName(name = adminRegisterDTO.name) != null) {
            return ServerResponse.status(HttpStatus.CONFLICT).bodyValueAndAwait("Same name account already exists")
        }
        val encryptedPassword = bCryptPasswordEncoder.encode(adminRegisterDTO.password)
        val admin = Admin(name = adminRegisterDTO.name, password = encryptedPassword)

        // save admin
        adminRepository.save(admin)

        val (accessToken, refreshToken) = jwtService.generateTokenFromUser(admin.toUser())

        return ServerResponse.ok().headers { header ->
            header[HttpHeaders.AUTHORIZATION] = accessToken
        }.bodyValueAndAwait(Json.encodeToString(AdminDTO(admin.name)))
    }

    suspend fun login(request: ServerRequest): ServerResponse {
        val body = request.awaitBody<String>()
        val userLoginDTO = Json.decodeFromString<AdminLoginDTO>(string = body)

        val adminInDatabase: Admin =
            adminRepository.findAdminByName(userLoginDTO.name) ?: return ServerResponse.badRequest().buildAndAwait()
        if (!bCryptPasswordEncoder.matches(userLoginDTO.password, adminInDatabase.password)) {
            return ServerResponse.badRequest().buildAndAwait()
        }


        val (accessToken, refreshToken) = jwtService.generateTokenFromUser(adminInDatabase.toUser())

        val adminDTO = adminInDatabase.toAdminDTO()
        return ServerResponse.ok().headers { header ->
            header[HttpHeaders.AUTHORIZATION] = accessToken
        }.contentType(MediaType.APPLICATION_JSON).bodyValueAndAwait(Json.encodeToString(adminDTO))
    }

    suspend fun tokenTest(request: ServerRequest): ServerResponse {
        val token = request.headers().asHttpHeaders()[HttpHeaders.AUTHORIZATION]?.get(0) ?: ""
        val decodedJwt = try {
            jwtService.decodeAccessToken(token)
        } catch (e: JWTVerificationException) {
            return ServerResponse.status(HttpStatus.UNAUTHORIZED).bodyValueAndAwait(e.message ?: e.localizedMessage)
        }
        val adminName = decodedJwt.subject
        val admin = adminRepository.findAdminByName(adminName) ?: return ServerResponse.badRequest().bodyValueAndAwait("The token is expired")
        val adminDTO = admin.toAdminDTO()
        return ServerResponse.ok().bodyValueAndAwait(Json.encodeToString(adminDTO))
    }
}