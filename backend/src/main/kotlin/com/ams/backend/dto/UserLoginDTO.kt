package com.ams.backend.dto


data class UserLoginDTO(val name: String, val password: String, val rememberMe: Boolean = false)
