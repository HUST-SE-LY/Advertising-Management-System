package com.ams.backend.model.jwt

import com.ams.backend.dto.UserDTO

data class JwtUser(val token: String, val user: UserDTO)