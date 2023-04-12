package com.ams.backend.service

import com.ams.backend.dto.UserDTO
import com.ams.backend.dto.UserLoginDTO
import com.ams.backend.model.jwt.JwtUser
import com.ams.backend.utils.JwtTokenUtil
import org.springframework.security.authentication.BadCredentialsException
import org.springframework.security.core.context.SecurityContextHolder
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.stereotype.Service

@Service
class AuthService(val adminService: AdminService, val bCryptPasswordEncoder: BCryptPasswordEncoder) {

    suspend fun authLogin(userLoginDTO: UserLoginDTO): JwtUser {
        val username = userLoginDTO.name

        val admin = adminService.getAdminByName(username)
        if (bCryptPasswordEncoder.matches(userLoginDTO.password,  admin.password)) {
            val token = JwtTokenUtil.generateToken(username = username, isRemember = userLoginDTO.rememberMe)

            val authentication = JwtTokenUtil.getAuthentication(token)
            SecurityContextHolder.getContext().authentication = authentication

            val userDTO = UserDTO(name = username)

            return JwtUser(token, userDTO)
        }
        throw BadCredentialsException("The user name or password error.");
    }

    suspend fun logout() = SecurityContextHolder.clearContext()

}