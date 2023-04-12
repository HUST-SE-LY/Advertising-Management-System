package com.ams.backend.controller

import com.ams.backend.dto.UserLoginDTO
import com.ams.backend.dto.UserRegisterDTO
import com.ams.backend.dto.AdminDTO
import com.ams.backend.security.SecurityConstants
import com.ams.backend.service.AdminService
import com.ams.backend.service.AuthService
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.springframework.dao.EmptyResultDataAccessException
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/admin")
class AdminController(private val adminService: AdminService, private val authService: AuthService) {

    @PostMapping("/register")
    suspend fun register(@RequestBody userRegisterDTO: UserRegisterDTO): ResponseEntity<String> {
        val registeredAdmin = adminService.registerAdmin(userRegisterDTO.name, userRegisterDTO.password)
        val responseBody = AdminDTO(admin = registeredAdmin.name)
        return ResponseEntity<String>(
            Json.encodeToString(responseBody),
            HttpStatus.OK
        )
    }

    @PostMapping("/login")
    suspend fun login(@RequestBody userLoginDTO: UserLoginDTO): ResponseEntity<String> {
        val jwtUser = authService.authLogin(userLoginDTO)
        val httpHeaders = HttpHeaders().also {
            it[SecurityConstants.TOKEN_HEADER] = SecurityConstants.TOKEN_PREFIX + jwtUser.token
        }
        return ResponseEntity<String>(
            Json.encodeToString(jwtUser.user),
            httpHeaders,
            HttpStatus.OK
        )
    }

    // Only for debugging
    @GetMapping("/list")
    suspend fun listAllAdmins(): ResponseEntity<String> {
        val admins = adminService.listAllAdmins()
        val response = admins.map { AdminDTO(it.name) }
        return ResponseEntity<String>(
            Json.encodeToString(response),
            HttpStatus.OK
        )
    }

    @ExceptionHandler(IllegalArgumentException::class)
    fun handleIllegalArgumentException(exception: IllegalArgumentException): ResponseEntity<String> {
        val responseBody = AdminDTO(errorMessage = exception.message)
        return ResponseEntity<String>(Json.encodeToString(responseBody), HttpStatus.BAD_REQUEST)
    }

    @ExceptionHandler(EmptyResultDataAccessException::class)
    fun handleEmptyResultDataAccessException(exception: EmptyResultDataAccessException): ResponseEntity<String> {
        val responseBody = AdminDTO(errorMessage = exception.message)
        return ResponseEntity<String>(Json.encodeToString(responseBody), HttpStatus.NOT_FOUND)
    }
}