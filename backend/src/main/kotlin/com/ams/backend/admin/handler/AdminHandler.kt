package com.ams.backend.admin.handler

import com.ams.backend.admin.dto.AdminDTO
import com.ams.backend.admin.dto.AdminRegisterDTO
import com.ams.backend.admin.repository.AdminRepository
import com.ams.backend.admin.dto.AdminLoginDTO
import com.ams.backend.admin.model.Admin
import com.ams.backend.security.JwtService
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.springframework.http.HttpHeaders
import org.springframework.http.MediaType
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.stereotype.Component
import org.springframework.web.reactive.function.server.*
import kotlin.time.Duration.Companion.days
import kotlin.time.Duration.Companion.hours

@Component
class AdminHandler(
    private val adminRepository: AdminRepository,
    private val bCryptPasswordEncoder: BCryptPasswordEncoder,
    private val jwtService: JwtService
) {

    companion object {
        private val TWO_HOURS = 2.hours
        private val SEVEN_DAYS = 7.days

    }

    private fun generateTokenFromAdmin(admin: Admin): Pair<String, String> {
        val adminUser = admin.toUser()
        val (accessToken, refreshToken) = run {
            val roles = adminUser.authorities.map { it.authority }.toTypedArray()
            jwtService.accessToken(adminUser.username, TWO_HOURS, roles) to jwtService.refreshToken(
                adminUser.username,
                SEVEN_DAYS, roles
            )
        }
        return accessToken to refreshToken
    }

    suspend fun register(request: ServerRequest): ServerResponse {
        val body = request.awaitBody<String>()
        val adminRegisterDTO = Json.decodeFromString<AdminRegisterDTO>(string = body)
        if (adminRepository.findAdminByName(name = adminRegisterDTO.name) != null) {
            return ServerResponse.badRequest().bodyValueAndAwait("Same name account already exists")
        }
        val encryptedPassword = bCryptPasswordEncoder.encode(adminRegisterDTO.password)
        val admin = Admin(name = adminRegisterDTO.name, password = encryptedPassword)

        // save admin
        adminRepository.save(admin)

        val (accessToken, refreshToken) = generateTokenFromAdmin(admin)

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


        val (accessToken, refreshToken) = generateTokenFromAdmin(adminInDatabase)

        val adminDTO = AdminDTO(name = adminInDatabase.name)
        return ServerResponse.ok().headers { header ->
            header[HttpHeaders.AUTHORIZATION] = accessToken
        }.contentType(MediaType.APPLICATION_JSON).bodyValueAndAwait(Json.encodeToString(adminDTO))
    }
}